// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn `tfsdk:"vpn" vyos:"vpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Between current VRF and VPN

`,
			Description: `Between current VRF and VPN

`,
		},
	}
}
