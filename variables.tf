variable "cluster_name" {
  description = "The name of the ECS cluster"
  type        = string
}

variable "enable_container_insights" {
  description = "Whether to enable CloudWatch Container Insights for the cluster"
  type        = bool
  default     = true
}

variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}
