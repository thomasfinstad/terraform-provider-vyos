// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap struct {
	// LeafNodes
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMapExport customtypes.CustomStringValue `tfsdk:"export" json:"export,omitempty"`
	VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMapImport customtypes.CustomStringValue `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-map to filter outgoing route updates

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		"import": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-map to filter incoming route updates

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
