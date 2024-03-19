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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationModePerNexthop types.Bool `tfsdk:"per_nexthop" vyos:"per-nexthop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"per_nexthop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allocate a label per connected next-hop in the VRF

`,
			Description: `Allocate a label per connected next-hop in the VRF

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
