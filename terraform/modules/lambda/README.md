<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.lambda_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.lambda_policy_attachment](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_architectures"></a> [architectures](#input\_architectures) | n/a | `list` | <pre>[<br>  "x86_64"<br>]</pre> | no |
| <a name="input_environment_variables"></a> [environment\_variables](#input\_environment\_variables) | n/a | `map` | `{}` | no |
| <a name="input_handler"></a> [handler](#input\_handler) | n/a | `string` | `"src/myHandler"` | no |
| <a name="input_layers"></a> [layers](#input\_layers) | n/a | `list` | `[]` | no |
| <a name="input_memory_size"></a> [memory\_size](#input\_memory\_size) | n/a | `string` | `"1024"` | no |
| <a name="input_name"></a> [name](#input\_name) | n/a | `string` | `"terratest-dev"` | no |
| <a name="input_policy_json"></a> [policy\_json](#input\_policy\_json) | n/a | `string` | n/a | yes |
| <a name="input_project"></a> [project](#input\_project) | n/a | `string` | `"terratest-sample"` | no |
| <a name="input_runtime"></a> [runtime](#input\_runtime) | n/a | `string` | `"python3.8"` | no |
| <a name="input_source_bucket"></a> [source\_bucket](#input\_source\_bucket) | n/a | `string` | `"my-terratest-bucket-dv"` | no |
| <a name="input_squad"></a> [squad](#input\_squad) | n/a | `string` | `"MySquad"` | no |
| <a name="input_timeout"></a> [timeout](#input\_timeout) | n/a | `string` | `"5"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | n/a |
| <a name="output_function_name"></a> [function\_name](#output\_function\_name) | n/a |
| <a name="output_invoke_arn"></a> [invoke\_arn](#output\_invoke\_arn) | n/a |
<!-- END_TF_DOCS -->