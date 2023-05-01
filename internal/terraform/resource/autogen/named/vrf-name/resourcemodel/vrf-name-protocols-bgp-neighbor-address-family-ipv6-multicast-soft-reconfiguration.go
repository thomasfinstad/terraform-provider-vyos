// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound types.String `tfsdk:"inbound" json:"inbound,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound.IsUnknown() {
		jsonData["inbound"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfiguration) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["inbound"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastSoftReconfigurationInbound = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
