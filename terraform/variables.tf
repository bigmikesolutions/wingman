variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "eu-west-1"
}

variable "aws_access_key" {
  description = "AWS Access Key"
  type        = string
}

variable "aws_secret_key" {
  description = "AWS Secret Key"
  type        = string
}

variable "cognito_user_pool_name" {
  description = "Name of Cognito User Pool"
  type        = string
  default     = "wingman"
}

variable "cognito_app_client_name" {
  description = "Name of Cognito App Client"
  type        = string
}

variable "cognito_idp_endpoint" {
  description = "Name of Cognito App Client"
  type        = string
}