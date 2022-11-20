package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Example s3 test
// Ideally tests need to run in a fully separate env under another AWS account
// Here I use asia region for that purpouse
func TestS3Bucket(t *testing.T) {
	// TODO: refactor so that test and snippets variables are set via one config
	awsRegion := "ap-south-1"

	// Authentication method for terraform >= v4 has changed
	// it will not check env variables therefore
	// using aws-vault for passing these variables from keychain process
	// wouldn't be possible
	// TODO: find a solution to upgrade to terraform v4
	terraformOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../snippets/",

		// Override values set in terraform.tfvars
		Vars: map[string]interface{}{
			"bucket_name": fmt.Sprintf("-%v", strings.ToLower(random.UniqueId())),
		},

		// set env vars
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Destroy after running a test
	defer terraform.Destroy(t, terraformOpts)

	// Deploy the infrastructure with the options defined above
	terraform.InitAndApply(t, terraformOpts)

	// Get the bucket ID
	bucketID := terraform.Output(t, terraformOpts, "bucket_id")

	// Testing s3 bucket configuration
	// The key point of this test
	// Get the versioning status
	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)

	// Test for expected configuration
	assert.Equal(t, "Enabled", actualStatus)
}
