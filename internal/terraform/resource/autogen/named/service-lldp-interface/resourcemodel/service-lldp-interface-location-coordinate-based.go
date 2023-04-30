// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceLldpInterfaceLocationCoordinateBased describes the resource data model.
type ServiceLldpInterfaceLocationCoordinateBased struct {
	// LeafNodes
	ServiceLldpInterfaceLocationCoordinateBasedAltitude  customtypes.CustomStringValue `tfsdk:"altitude" json:"altitude,omitempty"`
	ServiceLldpInterfaceLocationCoordinateBasedDatum     customtypes.CustomStringValue `tfsdk:"datum" json:"datum,omitempty"`
	ServiceLldpInterfaceLocationCoordinateBasedLatitude  customtypes.CustomStringValue `tfsdk:"latitude" json:"latitude,omitempty"`
	ServiceLldpInterfaceLocationCoordinateBasedLongitude customtypes.CustomStringValue `tfsdk:"longitude" json:"longitude,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceLldpInterfaceLocationCoordinateBased) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"altitude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Altitude in meters

|  Format  |  Description  |
|----------|---------------|
|  0  |  No altitude  |
|  [+-]<meters>  |  Altitude in meters  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"datum": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Coordinate datum type

|  Format  |  Description  |
|----------|---------------|
|  WGS84  |  WGS84  |
|  NAD83  |  NAD83  |
|  MLLW  |  NAD83/MLLW  |
`,

			// Default:          stringdefault.StaticString(`WGS84`),
			Computed: true,
		},

		"latitude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Latitude

|  Format  |  Description  |
|----------|---------------|
|  <latitude>  |  Latitude (example "37.524449N")  |
`,
		},

		"longitude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Longitude

|  Format  |  Description  |
|----------|---------------|
|  <longitude>  |  Longitude (example "122.267255W")  |
`,
		},

		// TagNodes

		// Nodes

	}
}
