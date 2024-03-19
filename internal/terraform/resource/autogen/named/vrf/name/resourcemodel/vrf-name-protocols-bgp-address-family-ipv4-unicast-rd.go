// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn `tfsdk:"vpn" vyos:"vpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRdVpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
			Description: `Between current address-family and VPN

`,
		},
	}
}
