package test

import (
	"fmt"
	"testing"

	awsSdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestKms(t *testing.T) {
	t.Parallel()

	awsRegion := "us-east-1"

	aliasName := fmt.Sprintf("test/%s", random.UniqueId())

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../kms",

		Vars: map[string]interface{}{
			"region":      awsRegion,
			"environment": "dev",
			"alias_name":  aliasName,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	kmsClient := aws.NewKmsClient(t, awsRegion)

	kmsArn := aws.GetCmkArn(t, awsRegion, "alias/"+aliasName)

	result, err := kmsClient.DescribeKey(&kms.DescribeKeyInput{
		KeyId: awsSdk.String(kmsArn),
	})

	fmt.Println(result.KeyMetadata)

	keyRotationStatus, err := kmsClient.GetKeyRotationStatus(&kms.GetKeyRotationStatusInput{
		KeyId: awsSdk.String(kmsArn),
	})

	if err != nil {
		fmt.Printf(err.Error())
	}

	expectedDescription := "description"
	actualDescription := terraform.Output(t, terraformOptions, "description")

	assert.Equal(t, expectedDescription, actualDescription)

	expectedKmsKeyUsage := "ENCRYPT_DECRYPT"
	actualKmsKeyUsage := *result.KeyMetadata.KeyUsage

	assert.Equal(t, expectedKmsKeyUsage, actualKmsKeyUsage)

	expectedKmsEnabled := true
	actualKmsKeyEnabled := *result.KeyMetadata.Enabled

	assert.Equal(t, expectedKmsEnabled, actualKmsKeyEnabled)

	expectedKeyRotationStatus := true
	actualKeyRotationStatus := *keyRotationStatus.KeyRotationEnabled

	assert.Equal(t, expectedKeyRotationStatus, actualKeyRotationStatus)

}
