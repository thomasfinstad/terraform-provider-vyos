// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfSegmentRoutingGlobalBlock{}

// VrfNameProtocolsOspfSegmentRoutingGlobalBlock describes the resource data model.
type VrfNameProtocolsOspfSegmentRoutingGlobalBlock struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfSegmentRoutingGlobalBlockLowLabelValue  types.Number `tfsdk:"low_label_value" vyos:"low-label-value,omitempty"`
	LeafVrfNameProtocolsOspfSegmentRoutingGlobalBlockHighLabelValue types.Number `tfsdk:"high_label_value" vyos:"high-label-value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfSegmentRoutingGlobalBlock) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"low_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label lower bound

    |  Format      |  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  |  Label value (recommended minimum value: 300)  |
`,
			Description: `MPLS label lower bound

    |  Format      |  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  |  Label value (recommended minimum value: 300)  |
`,
		},

		"high_label_value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label upper bound

    |  Format      |  Description  |
    |--------------|---------------|
    |  16-1048575  |  Label value  |
`,
			Description: `MPLS label upper bound

    |  Format      |  Description  |
    |--------------|---------------|
    |  16-1048575  |  Label value  |
`,
		},

		// Nodes

	}
}
