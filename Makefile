.PHONY: test test-unit test-init test-cleanup

# Install dependencies for testing
test-init:
	cd test && go mod tidy

# Run all tests
test: test-unit

# Run unit tests
test-unit: test-init
	cd test && go test -v ./unit -timeout 30m

# Cleanup test resources if needed
test-cleanup:
	cd test/unit && go test -timeout 10m -run TestEcsCleanup

# Format terraform files
fmt:
	terraform fmt -recursive

# Validate terraform files
validate:
	terraform validate

# Check terraform code with tflint
lint:
	tflint --recursive