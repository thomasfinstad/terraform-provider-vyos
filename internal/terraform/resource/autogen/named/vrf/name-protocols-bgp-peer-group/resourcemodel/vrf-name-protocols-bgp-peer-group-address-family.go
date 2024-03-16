// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamily{}

// VrfNameProtocolsBgpPeerGroupAddressFamily describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamily struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast        *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast        `tfsdk:"ipv4_unicast" vyos:"ipv4-unicast,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast `tfsdk:"ipv4_labeled_unicast" vyos:"ipv4-labeled-unicast,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn            *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn            `tfsdk:"ipv4_vpn" vyos:"ipv4-vpn,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast         *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast         `tfsdk:"ipv6_unicast" vyos:"ipv6-unicast,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast  *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast  `tfsdk:"ipv6_labeled_unicast" vyos:"ipv6-labeled-unicast,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn             *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn             `tfsdk:"ipv6_vpn" vyos:"ipv6-vpn,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn           *VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn           `tfsdk:"l2vpn_evpn" vyos:"l2vpn-evpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamily) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP neighbor parameters

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 VPN BGP neighbor parameters

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP neighbor parameters

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 Labeled Unicast BGP neighbor parameters

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 VPN BGP neighbor parameters

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
		},
	}
}
