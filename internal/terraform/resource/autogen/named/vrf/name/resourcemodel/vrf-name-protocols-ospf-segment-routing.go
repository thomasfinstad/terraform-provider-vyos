// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfSegmentRouting{}

// VrfNameProtocolsOspfSegmentRouting describes the resource data model.
type VrfNameProtocolsOspfSegmentRouting struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfSegmentRoutingMaximumLabelDepth types.Number `tfsdk:"maximum_label_depth" vyos:"maximum-label-depth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsOspfSegmentRoutingPrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
	NodeVrfNameProtocolsOspfSegmentRoutingGlobalBlock *VrfNameProtocolsOspfSegmentRoutingGlobalBlock `tfsdk:"global_block" vyos:"global-block,omitempty"`
	NodeVrfNameProtocolsOspfSegmentRoutingLocalBlock  *VrfNameProtocolsOspfSegmentRoutingLocalBlock  `tfsdk:"local_block" vyos:"local-block,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfSegmentRouting) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"maximum_label_depth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum MPLS labels allowed for this router

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-16    |  MPLS label depth  |
`,
			Description: `Maximum MPLS labels allowed for this router

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-16    |  MPLS label depth  |
`,
		},

		// Nodes

		"global_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfSegmentRoutingGlobalBlock{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment Routing Global Block label range

`,
			Description: `Segment Routing Global Block label range

`,
		},

		"local_block": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfSegmentRoutingLocalBlock{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment Routing Local Block label range

`,
			Description: `Segment Routing Local Block label range

`,
		},
	}
}
