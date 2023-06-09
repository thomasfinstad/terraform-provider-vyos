// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ContainerRegistry describes the resource data model.
type ContainerRegistry struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafContainerRegistryDisable types.String `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
	NodeContainerRegistryAuthentication *ContainerRegistryAuthentication `tfsdk:"authentication" json:"authentication,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerRegistry) GetVyosPath() []string {
	return []string{
		"container",
		"registry",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerRegistry) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Registry Name

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ContainerRegistryAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerRegistry) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafContainerRegistryDisable.IsNull() && !o.LeafContainerRegistryDisable.IsUnknown() {
		jsonData["disable"] = o.LeafContainerRegistryDisable.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeContainerRegistryAuthentication).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeContainerRegistryAuthentication)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["authentication"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerRegistry) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafContainerRegistryDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerRegistryDisable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["authentication"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeContainerRegistryAuthentication = &ContainerRegistryAuthentication{}

		err = json.Unmarshal(subJSONStr, o.NodeContainerRegistryAuthentication)
		if err != nil {
			return err
		}
	}

	return nil
}
