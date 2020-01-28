package terraform

import (
	"encoding/json"
	"os/exec"
	"path"
	"testing"
)

// PlanAndShow returns terraform plan
func PlanAndShow(t *testing.T, terraformOptions *Options) *TerraformPlan {
	t.Helper()

	// Terraform plan
	tfPlanOutput := "terraform.tfplan"
	_, planErr := RunTerraformCommandE(t, terraformOptions, FormatArgs(terraformOptions, "plan", "-input=false", "-lock=false", "-out="+tfPlanOutput)...)
	if planErr != nil {
		t.Fatal(planErr)
	}

	// Read and parse the plan output
	planOutputPath := path.Join(terraformOptions.TerraformDir, tfPlanOutput)
	// if use RunTerraformCommandAndGetStdoutE() then json output in test log
	cmd := exec.Command("terraform", "show", "-json", planOutputPath)
	cmd.Dir = terraformOptions.TerraformDir
	jsonBytes, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	var plan TerraformPlan
	jsonErr := json.Unmarshal(jsonBytes, &plan)
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	return &plan
}

// InitAndPlanAndShow returns terraform plan
func InitAndPlanAndShow(t *testing.T, terraformOptions *Options) *TerraformPlan {
	t.Helper()

	Init(t, terraformOptions)
	Get(t, terraformOptions)

	return PlanAndShow(t, terraformOptions)
}
