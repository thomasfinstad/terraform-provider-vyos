// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne{}

// VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOneAlways   types.Bool   `tfsdk:"always" vyos:"always,omitempty"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOneMetric   types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOneRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelOne) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"always": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Always advertise default route

`,
			Description: `Always advertise default route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
			Description: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// Nodes

	}
}
