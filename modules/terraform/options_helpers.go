package terraform

// CreateOptions create a Options
func CreateOptions(testDir string, vars map[string]interface{}) *Options {
	terraformOptions := &Options{
		TerraformDir: testDir,
		Vars:         vars,
		NoColor:      true,
	}

	return terraformOptions
}
