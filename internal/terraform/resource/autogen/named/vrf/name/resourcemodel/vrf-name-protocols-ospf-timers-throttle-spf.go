// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfTimersThroTTLeSpf{}

// VrfNameProtocolsOspfTimersThroTTLeSpf describes the resource data model.
type VrfNameProtocolsOspfTimersThroTTLeSpf struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay           types.Number `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime types.Number `tfsdk:"initial_holdtime" vyos:"initial-holdtime,omitempty"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime     types.Number `tfsdk:"max_holdtime" vyos:"max-holdtime,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfTimersThroTTLeSpf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay from the first change received to SPF calculation

    |  Format    |  Description            |
    |------------|-------------------------|
    |  0-600000  |  Delay in milliseconds  |
`,
			Description: `Delay from the first change received to SPF calculation

    |  Format    |  Description            |
    |------------|-------------------------|
    |  0-600000  |  Delay in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`200`),
			Computed: true,
		},

		"initial_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Initial hold time between consecutive SPF calculations

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  0-600000  |  Initial hold time in milliseconds  |
`,
			Description: `Initial hold time between consecutive SPF calculations

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  0-600000  |  Initial hold time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"max_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum hold time

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  0-600000  |  Max hold time in milliseconds  |
`,
			Description: `Maximum hold time

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  0-600000  |  Max hold time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		// Nodes

	}
}
