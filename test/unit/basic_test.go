package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestEcsClusterBasicExample tests the basic example of the ECS Cluster module
func TestEcsClusterBasicExample(t *testing.T) {
	// Skip this test in CI/CD environments that don't have AWS credentials
	// Remove this line if running in a proper CI/CD environment with AWS credentials
	t.Skip("Skipping test in local environment. Run manually with AWS credentials.")

	// Construct the terraform options with default retryable errors
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/simple",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_region":               "us-west-2",
			"environment":              "test",
			"name_prefix":              "terratest",
			"enable_container_insights": true,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of outputs
	clusterArn := terraform.Output(t, terraformOptions, "cluster_arn")
	clusterName := terraform.Output(t, terraformOptions, "cluster_name")

	// Verify we're getting back the outputs we expect
	assert.Contains(t, clusterArn, "cluster")
	assert.Equal(t, "terratest-cluster", clusterName)
}

// TestEcsCleanup is a special test function that can be called to clean up resources
// if regular tests fail and don't clean up properly
func TestEcsCleanup(t *testing.T) {
	t.Skip("Skipping cleanup. Run manually if needed with: go test -run TestEcsCleanup")

	// Clean up simple example
	simpleOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/simple",
		Vars: map[string]interface{}{
			"aws_region":               "us-west-2",
			"environment":              "test",
			"name_prefix":              "terratest",
			"enable_container_insights": true,
		},
	})
	terraform.Destroy(t, simpleOptions)
}