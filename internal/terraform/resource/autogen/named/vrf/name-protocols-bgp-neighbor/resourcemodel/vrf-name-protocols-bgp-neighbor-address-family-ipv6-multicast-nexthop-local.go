// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged types.Bool `tfsdk:"unchanged" vyos:"unchanged,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unchanged": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
			Description: `Leave link-local nexthop unchanged for this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
