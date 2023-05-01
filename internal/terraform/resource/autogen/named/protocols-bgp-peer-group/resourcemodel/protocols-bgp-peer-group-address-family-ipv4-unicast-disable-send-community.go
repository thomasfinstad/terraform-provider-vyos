// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended types.String `tfsdk:"extended" json:"extended,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard types.String `tfsdk:"standard" json:"standard,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended.IsUnknown() {
		jsonData["extended"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard.IsUnknown() {
		jsonData["standard"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard.ValueString()
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunity) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["extended"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityExtended = basetypes.NewStringNull()
	}

	if value, ok := jsonData["standard"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastDisableSendCommunityStandard = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
