// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended types.String `tfsdk:"extended" json:"extended,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard types.String `tfsdk:"standard" json:"standard,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended.IsUnknown() {
		jsonData["extended"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard.IsUnknown() {
		jsonData["standard"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["extended"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended = basetypes.NewStringNull()
	}

	if value, ok := jsonData["standard"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
