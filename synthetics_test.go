package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDataDogSyntheticExample(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "./",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variable: synthetic-monitor-id
	output := terraform.Output(t, terraformOptions, "synthetic-monitor-id")

	// Convert synthetic-monitor-id float to Int64
	flt, _, err := big.ParseFloat(output, 10, 0, big.ToNearestEven)
	if err != nil {
		/* handle any parsing errors here */
	}
	MonitorID, _ := flt.Int64()

	// Spin up DD API client and call Monitors API for info on newly-created monitor by ID
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.GetMonitor(ctx, MonitorID, *datadogV1.NewGetMonitorOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Structure and parse JSON response
	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	m := make(map[string]interface{})
	jsonError := json.Unmarshal(responseContent, &m)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	// Test that type of alert == synthetics, ideally the test case would be a running monitor
	assert.Equal(t, "synthetics alert", m["type"])
}
