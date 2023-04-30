// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDisableSendCommunity describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDisableSendCommunity struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDisableSendCommunityExtended customtypes.CustomStringValue `tfsdk:"extended" json:"extended,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDisableSendCommunityStandard customtypes.CustomStringValue `tfsdk:"standard" json:"standard,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastDisableSendCommunity) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"extended": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable sending extended community attributes to this peer

`,
		},

		"standard": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable sending standard community attributes to this peer

`,
		},

		// TagNodes

		// Nodes

	}
}
