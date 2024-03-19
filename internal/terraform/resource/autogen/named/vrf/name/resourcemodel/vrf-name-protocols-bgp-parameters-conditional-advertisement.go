// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersConditionalAdvertisement{}

// VrfNameProtocolsBgpParametersConditionalAdvertisement describes the resource data model.
type VrfNameProtocolsBgpParametersConditionalAdvertisement struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersConditionalAdvertisementTimer types.Number `tfsdk:"timer" vyos:"timer,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersConditionalAdvertisement) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timer": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set period to rescan BGP table to check if condition is met

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  5-240   |  Period to rerun the conditional advertisement scanner process  |
`,
			Description: `Set period to rescan BGP table to check if condition is met

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  5-240   |  Period to rerun the conditional advertisement scanner process  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		// Nodes

	}
}
