// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsRIPDistributeListInterfacePrefixList describes the resource data model.
type ProtocolsRIPDistributeListInterfacePrefixList struct {
	// LeafNodes
	LeafProtocolsRIPDistributeListInterfacePrefixListIn  types.String `tfsdk:"in" json:"in,omitempty"`
	LeafProtocolsRIPDistributeListInterfacePrefixListOut types.String `tfsdk:"out" json:"out,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPDistributeListInterfacePrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to input packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to input packets  |

`,
		},

		"out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to output packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to output packets  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRIPDistributeListInterfacePrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsRIPDistributeListInterfacePrefixListIn.IsNull() && !o.LeafProtocolsRIPDistributeListInterfacePrefixListIn.IsUnknown() {
		jsonData["in"] = o.LeafProtocolsRIPDistributeListInterfacePrefixListIn.ValueString()
	}

	if !o.LeafProtocolsRIPDistributeListInterfacePrefixListOut.IsNull() && !o.LeafProtocolsRIPDistributeListInterfacePrefixListOut.IsUnknown() {
		jsonData["out"] = o.LeafProtocolsRIPDistributeListInterfacePrefixListOut.ValueString()
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
func (o *ProtocolsRIPDistributeListInterfacePrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["in"]; ok {
		o.LeafProtocolsRIPDistributeListInterfacePrefixListIn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPDistributeListInterfacePrefixListIn = basetypes.NewStringNull()
	}

	if value, ok := jsonData["out"]; ok {
		o.LeafProtocolsRIPDistributeListInterfacePrefixListOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsRIPDistributeListInterfacePrefixListOut = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
