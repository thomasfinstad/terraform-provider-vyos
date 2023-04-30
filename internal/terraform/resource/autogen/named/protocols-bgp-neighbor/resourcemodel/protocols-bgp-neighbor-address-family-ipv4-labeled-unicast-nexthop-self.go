// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelfForce customtypes.CustomStringValue `tfsdk:"force" json:"force,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastNexthopSelf) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"force": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set the next hop to self for reflected routes

`,
		},

		// TagNodes

		// Nodes

	}
}
