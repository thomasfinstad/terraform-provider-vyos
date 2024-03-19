// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnectedMetric   types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnectedRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
`,
			Description: `Metric for redistributed routes

    |  Format        |  Description                      |
    |----------------------|-----------------------------------|
    |  1-4294967295  |  Metric for redistributed routes  |
`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------------|------------------|
    |  txt     &emsp;|  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// Nodes

	}
}
