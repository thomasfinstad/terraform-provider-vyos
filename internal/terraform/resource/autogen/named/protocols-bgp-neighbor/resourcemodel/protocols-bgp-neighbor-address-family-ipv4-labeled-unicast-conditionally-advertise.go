// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise struct {
	// LeafNodes
	ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertiseAdvertiseMap customtypes.CustomStringValue `tfsdk:"advertise_map" json:"advertise-map,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertiseExistMap     customtypes.CustomStringValue `tfsdk:"exist_map" json:"exist-map,omitempty"`
	ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertiseNonExistMap  customtypes.CustomStringValue `tfsdk:"non_exist_map" json:"non-exist-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastConditionallyAdvertise) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-map to conditionally advertise routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		"exist_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		"non_exist_map": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
