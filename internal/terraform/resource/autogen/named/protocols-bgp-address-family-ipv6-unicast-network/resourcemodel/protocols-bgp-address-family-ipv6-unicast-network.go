// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpAddressFamilyIPvsixUnicastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixUnicastNetwork struct {
	// LeafNodes
	ProtocolsBgpAddressFamilyIPvsixUnicastNetworkPathLimit customtypes.CustomStringValue `tfsdk:"path_limit" json:"path-limit,omitempty"`
	ProtocolsBgpAddressFamilyIPvsixUnicastNetworkRouteMap  customtypes.CustomStringValue `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicastNetwork) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"path_limit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `AS-path hopcount limit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  AS path hop count limit  |
`,
		},

		"route_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
