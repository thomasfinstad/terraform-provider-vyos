// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchangedAsPath  types.Bool `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchangedMed     types.Bool `tfsdk:"med" vyos:"med,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchangedNextHop types.Bool `tfsdk:"next_hop" vyos:"next-hop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send AS path unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"med": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"next_hop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
