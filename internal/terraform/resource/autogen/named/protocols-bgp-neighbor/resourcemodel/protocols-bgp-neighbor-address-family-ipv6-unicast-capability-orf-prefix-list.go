// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixListReceive customtypes.CustomStringValue `tfsdk:"receive" json:"receive,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixListSend    customtypes.CustomStringValue `tfsdk:"send" json:"send,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Capability to receive the ORF

`,
		},

		"send": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Capability to send the ORF

`,
		},

		// TagNodes

		// Nodes

	}
}
