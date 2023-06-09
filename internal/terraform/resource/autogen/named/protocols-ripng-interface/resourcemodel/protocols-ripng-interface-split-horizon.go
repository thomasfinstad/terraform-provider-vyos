// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsRIPngInterfaceSplitHorizon describes the resource data model.
type ProtocolsRIPngInterfaceSplitHorizon struct {
	// LeafNodes
	LeafProtocolsRIPngInterfaceSplitHorizonDisable       types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse types.String `tfsdk:"poison_reverse" json:"poison-reverse,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngInterfaceSplitHorizon) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
		},

		"poison_reverse": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPngInterfaceSplitHorizon) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.IsNull() && !o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.IsUnknown() {
		jsonData["disable"] = o.LeafProtocolsRIPngInterfaceSplitHorizonDisable.ValueString()
	}

	if !o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.IsNull() && !o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.IsUnknown() {
		jsonData["poison-reverse"] = o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse.ValueString()
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
func (o *ProtocolsRIPngInterfaceSplitHorizon) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafProtocolsRIPngInterfaceSplitHorizonDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngInterfaceSplitHorizonDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["poison-reverse"]; ok {
		o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
