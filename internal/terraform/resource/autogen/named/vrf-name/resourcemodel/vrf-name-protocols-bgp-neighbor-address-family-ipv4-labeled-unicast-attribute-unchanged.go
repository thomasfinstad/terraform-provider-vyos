// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedAsPath  customtypes.CustomStringValue `tfsdk:"as_path" json:"as-path,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedMed     customtypes.CustomStringValue `tfsdk:"med" json:"med,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchangedNextHop customtypes.CustomStringValue `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastAttributeUnchanged) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Send AS path unchanged

`,
		},

		"med": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
		},

		"next_hop": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Send nexthop unchanged

`,
		},

		// TagNodes

		// Nodes

	}
}
