// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ContainerNameEnvironment describes the resource data model.
type ContainerNameEnvironment struct {
	// LeafNodes
	LeafContainerNameEnvironmentValue types.String `tfsdk:"value" json:"value,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameEnvironment) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set environment option value

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Set environment option value  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNameEnvironment) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafContainerNameEnvironmentValue.IsNull() && !o.LeafContainerNameEnvironmentValue.IsUnknown() {
		jsonData["value"] = o.LeafContainerNameEnvironmentValue.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerNameEnvironment) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["value"]; ok {
		o.LeafContainerNameEnvironmentValue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNameEnvironmentValue = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
