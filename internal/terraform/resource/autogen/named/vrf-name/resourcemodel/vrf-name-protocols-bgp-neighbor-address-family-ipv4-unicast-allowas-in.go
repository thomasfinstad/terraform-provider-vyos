// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber types.String `tfsdk:"number" json:"number,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of times AS is allowed in path  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.IsUnknown() {
		jsonData["number"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["number"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
