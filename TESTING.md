# Testing the ECS Fargate Module

This document explains how to test the AWS ECS Fargate Terraform module in your AWS environment.

## Prerequisites

- AWS CLI installed and configured
- Terraform CLI installed (>= 0.13.0)
- AWS account with permissions to create necessary resources

## Test in the devops-toryordonline AWS Account

1. Clone the repository:
```bash
git clone https://github.com/amaudy/tfecsfargate.git
cd tfecsfargate
```

2. Choose an example to test:
- `examples/simple` - Basic ECS Fargate deployment
- `examples/with_autoscaling` - ECS with Auto Scaling

```bash
cd examples/simple
```

3. Copy the example variables file:
```bash
cp terraform.tfvars.example terraform.tfvars
```

4. Edit the `terraform.tfvars` file to customize your deployment:
```
aws_region  = "ap-southeast-1"  # Update with your preferred region
environment = "dev"
name_prefix = "toryord"         # Update with your project name
```

5. Set up your AWS credentials:
```bash
export AWS_PROFILE=devops-toryordonline
```

6. Initialize Terraform:
```bash
terraform init
```

7. Plan your deployment:
```bash
terraform plan
```

8. Apply the configuration:
```bash
terraform apply
```

9. Verify the deployment in the AWS Console:
- Check ECS clusters
- Verify the service is running correctly

10. Clean up when finished:
```bash
terraform destroy
```

## Test Cases

When testing this module, verify:

1. **Basic Functionality**:
   - ECS cluster is created
   - Task definition is registered
   - Service is deployed and running

2. **Auto Scaling**:
   - Scaling policies are created
   - Service scales when thresholds are breached (you can test this by simulating load)

## Best Practices

- Keep test deployments short-lived to minimize costs
- Use unique prefixes to avoid name conflicts with existing resources
- Test in a non-production AWS account
- Validate all resources are properly tagged

