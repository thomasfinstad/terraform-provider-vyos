// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfAreaAreaTypeStub describes the resource data model.
type VrfNameProtocolsOspfAreaAreaTypeStub struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost types.String `tfsdk:"default_cost" json:"default-cost,omitempty"`
	LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary   types.String `tfsdk:"no_summary" json:"no-summary,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaAreaTypeStub) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_cost": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Summary-default cost

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Summary default cost  |

`,
		},

		"no_summary": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfAreaAreaTypeStub) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost.IsNull() && !o.LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost.IsUnknown() {
		jsonData["default-cost"] = o.LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost.ValueString()
	}

	if !o.LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary.IsNull() && !o.LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary.IsUnknown() {
		jsonData["no-summary"] = o.LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary.ValueString()
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
func (o *VrfNameProtocolsOspfAreaAreaTypeStub) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["default-cost"]; ok {
		o.LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaAreaTypeStubDefaultCost = basetypes.NewStringNull()
	}

	if value, ok := jsonData["no-summary"]; ok {
		o.LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaAreaTypeStubNoSummary = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
