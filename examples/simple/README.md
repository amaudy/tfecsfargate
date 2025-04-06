# Simple ECS Fargate Example

This example demonstrates how to use the ECS Fargate module to deploy a simple service running an Nginx container.

## Prerequisites

- AWS CLI configured with appropriate credentials
- Terraform installed (>= 0.13.0)

## Usage

1. Copy the example variables file and modify as needed:

```bash
cp terraform.tfvars.example terraform.tfvars
```

2. Edit `terraform.tfvars` to configure your AWS region, environment name, and other parameters:

```hcl
aws_region  = "ap-southeast-1"  # Change to your region
environment = "dev"
name_prefix = "myproject"       # Change to your project name
```

3. Adjust AWS authentication:

**Using AWS profiles**
```bash
export AWS_PROFILE=devops-toryordonline
```

**Using direct credentials**
```bash
export AWS_ACCESS_KEY_ID=your_access_key
export AWS_SECRET_ACCESS_KEY=your_secret_key
```

4. Initialize Terraform:

```bash
terraform init
```

5. Plan the deployment:

```bash
terraform plan
```

6. Apply the configuration:

```bash
terraform apply
```

7. When finished, destroy the resources:

```bash
terraform destroy
```

## What it deploys

- ECS Cluster
- ECS Task Definition for a simple Nginx container
- ECS Service running on Fargate
- VPC, Subnets, and Security Group
- IAM Roles for task execution and task running