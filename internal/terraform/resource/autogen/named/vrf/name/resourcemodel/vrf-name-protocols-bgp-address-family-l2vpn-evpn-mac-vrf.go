// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrfSoo types.String `tfsdk:"soo" vyos:"soo,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnMacVrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"soo": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Site-of-Origin extended community

    |  Format  &emsp;|  Description                                                               |
    |----------------|----------------------------------------------------------------------------|
    |  ASN:NN  &emsp;|  based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    |  IP:NN   &emsp;|  Based on a router-id IP address in format &lt;IP:0-65535&gt;              |
`,
			Description: `Site-of-Origin extended community

    |  Format  |  Description                                                               |
    |----------------|----------------------------------------------------------------------------|
    |  ASN:NN  |  based on autonomous system number in format <0-65535:0-4294967295>  |
    |  IP:NN   |  Based on a router-id IP address in format <IP:0-65535>              |
`,
		},

		// Nodes

	}
}
