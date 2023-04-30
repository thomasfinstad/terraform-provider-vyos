// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicast struct {
	// LeafNodes

	// TagNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress types.Map `tfsdk:"aggregate_address" json:"aggregate-address,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork          types.Map `tfsdk:"network" json:"network,omitempty"`

	// Nodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance     types.Object `tfsdk:"distance" json:"distance,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastExport       types.Object `tfsdk:"export" json:"export,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport       types.Object `tfsdk:"import" json:"import,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabel        types.Object `tfsdk:"label" json:"label,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths types.Object `tfsdk:"maximum_paths" json:"maximum-paths,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd           types.Object `tfsdk:"rd" json:"rd,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteMap     types.Object `tfsdk:"route_map" json:"route-map,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget  types.Object `tfsdk:"route_target" json:"route-target,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute types.Object `tfsdk:"redistribute" json:"redistribute,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicast) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		"aggregate_address": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP aggregate network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  BGP aggregate network  |
`,
		},

		"network": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastNetwork{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  BGP network  |
`,
		},

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
		},

		"export": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastExport{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Export routes from this address-family

`,
		},

		"import": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastImport{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Import routes to this address-family

`,
		},

		"label": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabel{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Label value for VRF

`,
		},

		"maximum_paths": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Forward packets over multiple paths

`,
		},

		"rd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify route distinguisher

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteMap{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"route_target": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify route target list

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute routes from other protocols into BGP

`,
		},
	}
}
