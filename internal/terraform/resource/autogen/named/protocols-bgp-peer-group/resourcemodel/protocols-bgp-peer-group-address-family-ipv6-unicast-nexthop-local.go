// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal struct {
	// LeafNodes
	ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocalUnchanged customtypes.CustomStringValue `tfsdk:"unchanged" json:"unchanged,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastNexthopLocal) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unchanged": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
		},

		// TagNodes

		// Nodes

	}
}
