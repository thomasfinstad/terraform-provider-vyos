// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfRefresh{}

// VrfNameProtocolsOspfRefresh describes the resource data model.
type VrfNameProtocolsOspfRefresh struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfRefreshTimers types.Number `tfsdk:"timers" vyos:"timers,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRefresh) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timers": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Refresh timer

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  10-1800  |  Timer value in seconds  |
`,
			Description: `Refresh timer

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  10-1800  |  Timer value in seconds  |
`,
		},

		// Nodes

	}
}
