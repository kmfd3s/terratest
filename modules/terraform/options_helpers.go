package terraform

import (
	"os"
	"path/filepath"
	"strings"
)

// CreateOptions create a Options
func CreateOptions(testDir string, vars map[string]interface{}) *Options {
	terraformOptions := &Options{
		TerraformDir: testDir,
		Vars:         vars,
		NoColor:      true,
	}

	return terraformOptions
}

// TgCreateOptions create a Options for terragrunt
func TgCreateOptions(terraformDir string, relativeModulePath string, testID string, vars map[string]interface{}) *Options {
	env := map[string]string{
		"TERRAGRUNT_DOWNLOAD":  filepath.Join(os.TempDir(), testID),
		"TG_TEST_ID":           testID,
		"TG_USE_LOCAL_BACKEND": "true",
		"TF_CLI_ARGS":          "-no-color",
	}

	if source := os.Getenv("TERRAGRUNT_SOURCE"); 0 < len(source) {
		if index := strings.Index(source, "//"); -1 < index {
			source = source[:index]
		}
		env["TERRAGRUNT_SOURCE"] = source + "//" + relativeModulePath
	}

	terraformOptions := &Options{
		TerraformDir:    terraformDir,
		TerraformBinary: "terragrunt",
		Lock:            false,
		Vars:            vars,
		EnvVars:         env,
	}

	return terraformOptions
}
