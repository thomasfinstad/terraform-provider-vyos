// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityExtended types.Bool `tfsdk:"extended" vyos:"extended,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunityStandard types.Bool `tfsdk:"standard" vyos:"standard,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDisableSendCommunity) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"extended": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending extended community attributes to this peer

`,
			Description: `Disable sending extended community attributes to this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"standard": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending standard community attributes to this peer

`,
			Description: `Disable sending standard community attributes to this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
