variable "name" {
  type    = string
  default = "terratest-dev"
}

variable "source_bucket" {
  type    = string
  default = "my-terratest-bucket-dv"
}

variable "project" {
  type    = string
  default = "terratest-sample"
}

variable "squad" {
  type    = string
  default = "MySquad"
}

variable "policy_json" {
  type = string
}

variable "environment_variables" {
  default = {}
}

variable "timeout" {
  default = "10"
}

variable "memory_size" {
  default = "1024"
}

variable "handler" {
  default = "src/myHandler"
}

variable "layers" {
  default = []
}

variable "runtime" {
  default = "python3.8"
}

variable "architectures" {
  default = ["x86_64"]
}
