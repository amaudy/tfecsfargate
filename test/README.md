# Testing ECS Fargate Terraform Module

This directory contains tests for the ECS Fargate Terraform module using [Terratest](https://terratest.gruntwork.io/), a Go library that makes it easier to write automated tests for your infrastructure code.

## Prerequisites

- Go 1.16 or later
- AWS credentials configured
- Terraform installed

## Test Structure

- `unit/`: Unit tests for each component of the module

## Running Tests

To run all tests:

```bash
cd /path/to/tfecsfargate
make test
```

To run only unit tests:

```bash
make test-unit
```

Note: Tests are skipped by default since they require AWS credentials and will create real resources in your AWS account. To run them, you need to remove the `t.Skip()` lines in the test files.

## What's Being Tested

1. **Basic Functionality (TestEcsFargateBasicExample)**:
   - Verifies the simple example deploys correctly
   - Checks that outputs like cluster ARN and service name are set correctly

2. **Autoscaling Enabled (TestEcsFargateWithAutoscalingEnabled)**:
   - Tests that autoscaling resources are created when enabled
   - Verifies the output indicating autoscaling is enabled

3. **Autoscaling Disabled (TestEcsFargateWithAutoscalingDisabled)**:
   - Tests that autoscaling resources are not created when disabled
   - Verifies the output indicating autoscaling is disabled

## Cleanup

Tests automatically clean up resources by running `terraform destroy` when they finish. However, if a test fails, you might need to manually clean up resources.

To forcibly clean up resources from a failed test:

```bash
make test-cleanup
```