// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsIsisSegmentRoutingGlobalBlock describes the resource data model.
type VrfNameProtocolsIsisSegmentRoutingGlobalBlock struct {
	// LeafNodes
	VrfNameProtocolsIsisSegmentRoutingGlobalBlockLowLabelValue  customtypes.CustomStringValue `tfsdk:"low_label_value" json:"low-label-value,omitempty"`
	VrfNameProtocolsIsisSegmentRoutingGlobalBlockHighLabelValue customtypes.CustomStringValue `tfsdk:"high_label_value" json:"high-label-value,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsIsisSegmentRoutingGlobalBlock) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"low_label_value": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MPLS label lower bound

|  Format  |  Description  |
|----------|---------------|
|  u32:16-1048575  |  Label value (recommended minimum value: 300)  |
`,
		},

		"high_label_value": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MPLS label upper bound

|  Format  |  Description  |
|----------|---------------|
|  u32:16-1048575  |  Label value  |
`,
		},

		// TagNodes

		// Nodes

	}
}
