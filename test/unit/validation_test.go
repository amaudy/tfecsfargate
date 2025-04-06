package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// TestTerraformSimpleValidation validates the simple example without deploying
func TestTerraformSimpleValidation(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/simple",
		// Just validate the config, don't deploy anything
		NoColor: true,
	})

	// Run terraform init and terraform validate
	terraform.Init(t, terraformOptions)
	terraform.Validate(t, terraformOptions)
}

