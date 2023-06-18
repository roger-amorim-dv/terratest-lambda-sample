resource "aws_lambda_function" "main" {
  function_name = "${var.name}-${terraform.workspace}"
  role          = aws_iam_role.lambda_role.arn
  handler       = var.handler
  runtime       = var.runtime
  s3_bucket     = var.source_bucket
  s3_key        = "release/latest.zip"
  timeout       = var.timeout
  memory_size   = var.memory_size
  architectures = var.architectures
  layers        = var.layers

  tags = {
    Name    = var.name
    Project = var.project
    Squad   = var.squad
  }

  environment {
    variables = {
      QUEUE_NAME = "my_queue_name"
      TOPIC_NAME = "my_topic_name"
    }
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = "terratest-lambda-role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
