package test

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestLambdaModule(t *testing.T) {
	t.Parallel()

	modulePath := "../terraform/modules/lambda"

	policyJSON := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect":   "Allow",
				"Action":   "s3:GetObject",
				"Resource": "arn:aws:s3:::my-terratest-bucket-dv/*",
			},
		},
	}

	policyJSONString, err := json.Marshal(policyJSON)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	terraformOptions := &terraform.Options{
		TerraformDir: modulePath,
		Vars: map[string]interface{}{
			"name":            "my-lambda-function",
			"project":         "test",
			"handler":         "src/myHandler",
			"runtime":         "python3.8",
			"source_bucket":   "my-terratest-bucket-dv",
			"timeout":         60,
			"memory_size":     128,
			"policy_json":     string(policyJSONString),
		},
	}

	createOrSelectWorkspace(t, terraformOptions, "dev")

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	lambdaFunctionName := terraform.Output(t, terraformOptions, "function_name")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
	assert.NoError(t, err)

	client := lambda.New(sess)

	describeInput := &lambda.GetFunctionInput{
		FunctionName: aws.String(lambdaFunctionName),
	}

	describeOutput, err := client.GetFunction(describeInput)
	assert.NoError(t, err)

	// Assert that the Lambda function exists and has the expected values
	assert.Equal(t, "my-lambda-function-dev", aws.StringValue(describeOutput.Configuration.FunctionName))
	assert.Equal(t, "python3.8", aws.StringValue(describeOutput.Configuration.Runtime))
	assert.Equal(t, int64(60), *describeOutput.Configuration.Timeout)
	assert.Equal(t, "src/myHandler", aws.StringValue(describeOutput.Configuration.Handler))
}

func createOrSelectWorkspace(t *testing.T, terraformOptions *terraform.Options, workspaceName string) {
	_, err := terraform.RunTerraformCommandE(t, terraformOptions, "workspace", "select", workspaceName)
	if err != nil {
		terraform.RunTerraformCommand(t, terraformOptions, "workspace", "new", workspaceName)
	}
}
