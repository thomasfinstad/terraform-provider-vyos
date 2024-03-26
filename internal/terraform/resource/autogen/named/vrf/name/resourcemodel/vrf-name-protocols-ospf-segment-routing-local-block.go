// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfSegmentRoutingLocalBlock{}

// VrfNameProtocolsOspfSegmentRoutingLocalBlock describes the resource data model.
type VrfNameProtocolsOspfSegmentRoutingLocalBlock struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockLowLabelValue  types.Number `tfsdk:"low_label_value" vyos:"low-label-value,omitempty"`
	LeafVrfNameProtocolsOspfSegmentRoutingLocalBlockHighLabelValue types.Number `tfsdk:"high_label_value" vyos:"high-label-value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfSegmentRoutingLocalBlock) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
