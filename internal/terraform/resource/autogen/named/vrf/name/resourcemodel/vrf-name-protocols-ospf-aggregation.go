// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfAggregation{}

// VrfNameProtocolsOspfAggregation describes the resource data model.
type VrfNameProtocolsOspfAggregation struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAggregationTimer types.Number `tfsdk:"timer" vyos:"timer,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAggregation) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timer": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay timer

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  5-1800  |  Timer interval in seconds  |
`,
			Description: `Delay timer

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  5-1800  |  Timer interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}
