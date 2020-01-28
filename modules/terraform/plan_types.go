package terraform

import (
	"encoding/json"
	"testing"
)

// After is representations of the object value after the action.
type After struct {
	raw    []byte
	values map[string]interface{}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (after *After) UnmarshalJSON(data []byte) error {
	after.raw = data
	return nil
}

// Change describes the change that will be made to the indicated object.
type Change struct {
	Actions []string `json:"actions"`
	After   *After   `json:"after"`
}

// As converts After to v
func (after *After) As(t *testing.T, v interface{}) {
	t.Helper()

	if err := json.Unmarshal(after.raw, v); err != nil {
		t.Fatal(err)
	}
}

// GetValue returns value
func (after *After) GetValue(t *testing.T, name string) interface{} {
	t.Helper()

	if after.values == nil {
		after.As(t, &after.values)
	}

	v, ok := after.values[name]
	if !ok {
		t.Fatalf("%s not found", name)
	}

	return v
}

// GetValueAsMapArray returns map array
func (after *After) GetValueAsMapArray(t *testing.T, name string) []map[string]interface{} {
	t.Helper()

	values := after.GetValueAsArray(t, name)
	array := make([]map[string]interface{}, len(values))
	for i, v := range values {
		array[i] = v.(map[string]interface{})
	}
	return array
}

// GetValueAsMap returns map
func (after *After) GetValueAsMap(t *testing.T, name string) map[string]interface{} {
	t.Helper()

	v := after.GetValue(t, name)
	m, ok := v.(map[string]interface{})
	if !ok {
		t.Fatalf("%s failed to convert to map.", name)
	}

	return m
}

// GetValueAsArray returns array
func (after *After) GetValueAsArray(t *testing.T, name string) []interface{} {
	t.Helper()

	v := after.GetValue(t, name)
	array, ok := v.([]interface{})
	if !ok {
		t.Fatalf("%s failed to convert to array.", name)
	}

	return array
}

// StateOutput is a values representation object derived from the values in the state.
type StateOutput struct {
	Sensitive bool        `json:"sensitive"`
	Value     interface{} `json:"value"`
}

// StateValues is the representation of used in both state and plan
// output to describe current state and planned state.
type StateValues struct {
	Outputs map[string]*StateOutput `json:"outputs"`
}

// ResourceChange is a description of the individual change actions that Terraform
// plans to use to move from the prior state to a new state matching the
// configuration.
type ResourceChange struct {
	Address string  `json:"address"`
	Change  *Change `json:"change"`
}

// TerraformPlan represents the Terraform plan.
type TerraformPlan struct {
	FormatVersion    string             `json:"format_version"`
	TerraformVersion string             `json:"terraform_version"`
	PlannedValues    *StateValues       `json:"planned_values"`
	ResourceChanges  []*ResourceChange  `json:"resource_changes"`
	OutputChanges    map[string]*Change `json:"output_changes"`
}

// GetResource get a resource change
func (plan *TerraformPlan) GetResource(t *testing.T, address string) *ResourceChange {
	t.Helper()

	for _, resource := range plan.ResourceChanges {
		if resource.Address == address {
			return resource
		}
	}

	t.Fatalf("Resource not found: %s", address)
	return nil
}
