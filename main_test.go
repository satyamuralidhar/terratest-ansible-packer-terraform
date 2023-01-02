package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// intializing variables
var (
	globalBackendConf = make(map[string]interface{})
	globalEnvVars     = make(map[string]string)
	rsg               = "packer-rsg"
	packerimage       = "packerterrform01"
)

const (
	apiVersion              = "2019-06-01"
	resourceProvisionStatus = "Succeeded"
)

//set environmnet varibales

func setEnvVars() (map[string]string, error) {
	/* fetch env varibales form bashprofile */
	TF_VAR_ARM_CLIENT_ID := os.Getenv("TF_VAR_AZURE_CLIENT_ID")
	TF_VAR_ARM_CLIENT_SECRET := os.Getenv("TF_VAR_AZURE_CLIENT_SECRET")
	TF_VAR_ARM_TENANT_ID := os.Getenv("TF_VAR_AZURE_TENANT_ID")
	TF_VAR_ARM_SUBSCRIPTION_ID := os.Getenv("TF_VAR_AZURE_SUBSCRIPTION_ID")

	/* create env vars from globalEnvVars to call Terraform to Terratest */

	if TF_VAR_ARM_CLIENT_ID != "" {
		globalEnvVars["TF_VAR_ARM_CLIENT_ID"] = TF_VAR_ARM_CLIENT_ID
		globalEnvVars["TF_VAR_ARM_CLIENT_SECRET"] = TF_VAR_ARM_CLIENT_SECRET
		globalEnvVars["TF_VAR_ARM_TENANT_ID"] = TF_VAR_ARM_TENANT_ID
		globalEnvVars["TF_VAR_ARM_SUBSCRIPTION_ID"] = TF_VAR_ARM_SUBSCRIPTION_ID
	}

	return globalEnvVars, nil

}

/* Lamp Server Creation terratest */

func TestTerraform_LampServerInstallation(t *testing.T) {
	t.Parallel()

	/* call Env funtion */
	setEnvVars()
	expectedLocation := "eastus"
	expectedAddressSpace := "192.168.0.0/16"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// Set the path to the Terraform code that will be tested.
		TerraformDir: ".",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{

			"rsg":         rsg,
			"location":    expectedLocation,
			"vnet_cidr":   expectedAddressSpace,
			"packerimage": packerimage,
			//tags = var.tags
		},

		// globalvariables for user account
		EnvVars: globalEnvVars,
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,

		// Reconfigure is required if module deployment and go test pipelines are running in one stage
		Reconfigure: true,
	})
	defer terraform.Destroy(t, terraformOptions)
	time.Sleep(8 * time.Second)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	init, err := terraform.InitE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(init)

	plan, err := terraform.PlanE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(plan)
	//terraform applying
	apply, err := terraform.ApplyE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(apply)

	//out, err := terraform.OutputJsonE(t, terraformOptions, "resource_name")
	expectedVnetAddressSpace := terraform.Output(t, terraformOptions, "vnet_cidr")
	packerImageName := terraform.Output(t, terraformOptions, "packerimage")
	resourceGroupName := terraform.Output(t, terraformOptions, "rsg")
	expectedVmLocation := terraform.Output(t, terraformOptions, "location")

	fmt.Printf("VnetAddressSpace :: %s\n", expectedVnetAddressSpace)
	fmt.Printf("resourceGroupName :: %s\n", resourceGroupName)
	fmt.Printf("packerImage :: %s\n", packerImageName)
	fmt.Printf("Location :: %s\n", expectedVmLocation)
}
