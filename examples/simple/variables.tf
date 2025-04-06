variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-west-2"
}

variable "environment" {
  description = "Environment name for tagging and resource naming"
  type        = string
  default     = "dev"
}

variable "name_prefix" {
  description = "Prefix to use for resource naming"
  type        = string
  default     = "example"
}

variable "enable_container_insights" {
  description = "Whether to enable CloudWatch Container Insights for the cluster"
  type        = bool
  default     = true
}