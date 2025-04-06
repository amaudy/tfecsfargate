# AWS ECS Fargate Cluster Terraform Module

Hi there! 👋 This is a beginner-friendly Terraform module that creates an Amazon ECS cluster ready for Fargate tasks. This module focuses exclusively on setting up the ECS cluster infrastructure.

## What Does This Module Do?

This module helps you:

- Create an ECS cluster optimized for Fargate
- Configure CloudWatch Container Insights for monitoring
- Apply consistent tagging to your cluster resources

## How to Use This Module

Here's a simple example to get you started. Don't worry if it looks complicated - we've included full working examples in the `examples` folder that you can copy and modify!

```hcl
module "ecs_cluster" {
  source = "github.com/amaudy/tfecsfargate"

  # Name for your ECS cluster
  cluster_name = "my-fargate-cluster"
  
  # Enable CloudWatch Container Insights for monitoring
  enable_container_insights = true
  
  # Add labels to your resources
  tags = {
    Environment = "dev"
    Project     = "my-project"
  }
}
```

## Ready-to-Use Example

We've created a basic example that you can run with minimal changes:

1. **Simple Example** (`examples/simple`): Sets up a basic ECS cluster ready for Fargate tasks

The example has a README file with step-by-step instructions!

## What You Need Before Starting

* Terraform version 0.13.0 or newer
* AWS provider version 3.0.0 or newer

## Settings You Can Change (Inputs)

Here are all the settings you can customize in this module:

| Setting Name | What It Does | Type | Default | Required? |
|------|-------------|------|---------|:--------:|
| cluster_name | Name for your ECS cluster | `string` | - | Yes |
| enable_container_insights | Whether to enable CloudWatch Container Insights for monitoring | `bool` | `true` | No |
| tags | Labels to add to all your AWS resources | `map(string)` | `{}` | No |

## What You Get Back (Outputs)

After you create your resources, you'll get back these values that you can use elsewhere:

| Name | Description |
|------|-------------|
| cluster_id | ID of your ECS cluster |
| cluster_arn | Full Amazon Resource Name of your cluster |
| cluster_name | Name of your ECS cluster |

## Security Considerations

This module has been scanned with Checkov to identify security best practices. The module enables Container Insights by default which follows AWS security best practices for ECS cluster monitoring.

For more security recommendations when using this cluster, see the [SECURITY.md](SECURITY.md) document.

```bash
# Run a security scan on the module (requires Python and Checkov)
python -m venv .venv
source .venv/bin/activate
pip install checkov
checkov -d . --framework terraform
```

## Testing and Validation

### Local Testing

This module includes automated tests written with [Terratest](https://terratest.gruntwork.io/):

```bash
# Install test dependencies
make test-init

# Run all tests (skips by default - see test/README.md)
make test
```

For more information on running tests, see the [test/README.md](test/README.md) file.

### Continuous Integration

This module includes GitHub Actions workflows for continuous validation and security scanning:

1. **Terraform Validation**: Runs `terraform validate` on the module and examples, plus TFLint and TFSec scans
   - Automatically runs on push to main and pull requests
   - Ensures the module follows Terraform best practices
   - Enforces style conventions via TFLint

2. **Checkov Security Scan**: Runs Checkov for Terraform security scanning
   - Automatically runs on push to main, pull requests, and a weekly schedule
   - Posts security findings as comments on pull requests
   - Generates SARIF reports for GitHub code scanning dashboard

The workflows use the latest versions of all tools to ensure up-to-date security checks.

## Need Help?

Check out our examples folder for complete working examples, or open an issue on GitHub if you run into any problems!