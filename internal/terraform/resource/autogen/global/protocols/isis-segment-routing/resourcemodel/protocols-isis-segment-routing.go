// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsIsisSegmentRouting describes the resource data model.
type ProtocolsIsisSegmentRouting struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsIsisSegmentRoutingMaximumLabelDepth types.Number `tfsdk:"maximum_label_depth" vyos:"maximum-label-depth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsIsisSegmentRoutingPrefix bool `tfsdk:"-" vyos:"prefix,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsIsisSegmentRoutingGlobalBlock bool `tfsdk:"-" vyos:"global-block,ignore,omitempty"`
	ExistsNodeProtocolsIsisSegmentRoutingLocalBlock  bool `tfsdk:"-" vyos:"local-block,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsIsisSegmentRouting) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIsisSegmentRouting) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"isis",

		"segment-routing",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisSegmentRouting) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"maximum_label_depth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum MPLS labels allowed for this router

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-16  &emsp; |  MPLS label depth  |

`,
		},
	}
}
