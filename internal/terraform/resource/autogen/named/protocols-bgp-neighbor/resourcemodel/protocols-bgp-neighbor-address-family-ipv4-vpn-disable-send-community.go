// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended types.String `tfsdk:"extended" json:"extended,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard types.String `tfsdk:"standard" json:"standard,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"extended": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending extended community attributes to this peer

`,
		},

		"standard": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending standard community attributes to this peer

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended.IsUnknown() {
		jsonData["extended"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard.IsUnknown() {
		jsonData["standard"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["extended"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityExtended = basetypes.NewStringNull()
	}

	if value, ok := jsonData["standard"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunityStandard = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
