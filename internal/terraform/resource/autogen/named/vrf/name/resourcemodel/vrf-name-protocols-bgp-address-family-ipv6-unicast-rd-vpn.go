// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `For routes leaked from current address-family to VPN

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `For routes leaked from current address-family to VPN

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		// Nodes

	}
}
