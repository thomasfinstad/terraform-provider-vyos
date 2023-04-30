// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes

	// Nodes
	ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList types.Object `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrf) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixUnicastCapabilityOrfPrefixList{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
