// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfvthreeAreaRange describes the resource data model.
type VrfNameProtocolsOspfvthreeAreaRange struct {
	// LeafNodes
	VrfNameProtocolsOspfvthreeAreaRangeAdvertise    customtypes.CustomStringValue `tfsdk:"advertise" json:"advertise,omitempty"`
	VrfNameProtocolsOspfvthreeAreaRangeNotAdvertise customtypes.CustomStringValue `tfsdk:"not_advertise" json:"not-advertise,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAreaRange) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Advertise this range

`,
		},

		"not_advertise": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not advertise this range

`,
		},

		// TagNodes

		// Nodes

	}
}
