provider "aws" {
  region = var.aws_region
}

module "ecs_cluster" {
  source = "../../"

  cluster_name             = "${var.name_prefix}-cluster"
  enable_container_insights = var.enable_container_insights
  
  tags = {
    Environment = var.environment
    Project     = "${var.name_prefix}-ecs-cluster"
  }
}

output "cluster_arn" {
  value = module.ecs_cluster.cluster_arn
}

output "cluster_name" {
  value = module.ecs_cluster.cluster_name
}
