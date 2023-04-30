// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast struct {
	// LeafNodes

	// TagNodes
	VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress types.Map `tfsdk:"aggregate_address" json:"aggregate-address,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork          types.Map `tfsdk:"network" json:"network,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		"aggregate_address": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP aggregate network/prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  BGP aggregate network/prefix  |
`,
		},

		"network": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Labeled Unicast IPv6 BGP network/prefix  |
`,
		},

		// Nodes

	}
}
