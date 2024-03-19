// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `For routes leaked from current address-family to VPN

    |  Format                   &emsp;|  Description                                   |
    |---------------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  &emsp;|  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `For routes leaked from current address-family to VPN

    |  Format                   |  Description                                   |
    |---------------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		// Nodes

	}
}
