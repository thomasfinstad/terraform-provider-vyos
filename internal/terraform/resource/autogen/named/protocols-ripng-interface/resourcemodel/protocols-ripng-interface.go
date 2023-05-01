// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPngInterface describes the resource data model.
type ProtocolsRIPngInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsRIPngInterfaceSplitHorizon *ProtocolsRIPngInterfaceSplitHorizon `tfsdk:"split_horizon" json:"split-horizon,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPngInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"ripng",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		// LeafNodes

		// TagNodes

		// Nodes

		"split_horizon": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPngInterfaceSplitHorizon{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Split horizon parameters

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPngInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsRIPngInterfaceSplitHorizon).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsRIPngInterfaceSplitHorizon)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["split-horizon"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRIPngInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["split-horizon"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsRIPngInterfaceSplitHorizon = &ProtocolsRIPngInterfaceSplitHorizon{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsRIPngInterfaceSplitHorizon)
		if err != nil {
			return err
		}
	}

	return nil
}
