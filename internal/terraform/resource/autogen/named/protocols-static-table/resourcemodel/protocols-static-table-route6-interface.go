// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsStaticTableRoutesixInterface describes the resource data model.
type ProtocolsStaticTableRoutesixInterface struct {
	// LeafNodes
	LeafProtocolsStaticTableRoutesixInterfaceDisable  types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafProtocolsStaticTableRoutesixInterfaceDistance types.String `tfsdk:"distance" json:"distance,omitempty"`
	LeafProtocolsStaticTableRoutesixInterfaceVrf      types.String `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTableRoutesixInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for this route  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of VRF to leak to  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticTableRoutesixInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsStaticTableRoutesixInterfaceDisable.IsNull() && !o.LeafProtocolsStaticTableRoutesixInterfaceDisable.IsUnknown() {
		jsonData["disable"] = o.LeafProtocolsStaticTableRoutesixInterfaceDisable.ValueString()
	}

	if !o.LeafProtocolsStaticTableRoutesixInterfaceDistance.IsNull() && !o.LeafProtocolsStaticTableRoutesixInterfaceDistance.IsUnknown() {
		jsonData["distance"] = o.LeafProtocolsStaticTableRoutesixInterfaceDistance.ValueString()
	}

	if !o.LeafProtocolsStaticTableRoutesixInterfaceVrf.IsNull() && !o.LeafProtocolsStaticTableRoutesixInterfaceVrf.IsUnknown() {
		jsonData["vrf"] = o.LeafProtocolsStaticTableRoutesixInterfaceVrf.ValueString()
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
func (o *ProtocolsStaticTableRoutesixInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["disable"]; ok {
		o.LeafProtocolsStaticTableRoutesixInterfaceDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticTableRoutesixInterfaceDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["distance"]; ok {
		o.LeafProtocolsStaticTableRoutesixInterfaceDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticTableRoutesixInterfaceDistance = basetypes.NewStringNull()
	}

	if value, ok := jsonData["vrf"]; ok {
		o.LeafProtocolsStaticTableRoutesixInterfaceVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticTableRoutesixInterfaceVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
