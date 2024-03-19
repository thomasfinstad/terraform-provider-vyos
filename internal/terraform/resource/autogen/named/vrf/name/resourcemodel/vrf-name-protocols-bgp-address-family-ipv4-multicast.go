// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}

// VrfNameProtocolsBgpAddressFamilyIPvfourMulticast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourMulticast struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress bool `tfsdk:"aggregate_address" vyos:"aggregate-address,child"`
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourMulticastNetwork          bool `tfsdk:"network" vyos:"network,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance *VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance `tfsdk:"distance" vyos:"distance,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourMulticast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticastDistance{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
			Description: `Administrative distances for BGP routes

`,
		},
	}
}
