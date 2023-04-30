// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesEthernetVifS describes the resource data model.
type InterfacesEthernetVifS struct {
	// LeafNodes
	InterfacesEthernetVifSDescrIPtion       customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesEthernetVifSAddress           customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesEthernetVifSDisableLinkDetect customtypes.CustomStringValue `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	InterfacesEthernetVifSDisable           customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesEthernetVifSProtocol          customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	InterfacesEthernetVifSMac               customtypes.CustomStringValue `tfsdk:"mac" json:"mac,omitempty"`
	InterfacesEthernetVifSMtu               customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesEthernetVifSRedirect          customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesEthernetVifSVrf               customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes
	InterfacesEthernetVifSVifC types.Map `tfsdk:"vif_c" json:"vif-c,omitempty"`

	// Nodes
	InterfacesEthernetVifSDhcpOptions     types.Object `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	InterfacesEthernetVifSDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesEthernetVifSIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesEthernetVifSIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesEthernetVifSMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesEthernetVifS) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |
`,
		},

		"disable_link_detect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"protocol": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Protocol used for service VLAN (default: 802.1ad)

|  Format  |  Description  |
|----------|---------------|
|  802.1ad  |  Provider Bridging (IEEE 802.1ad, Q-inQ), ethertype 0x88a8  |
|  802.1q  |  VLAN-tagged frame (IEEE 802.1q), ethertype 0x8100  |
`,

			// Default:          stringdefault.StaticString(`802.1ad`),
			Computed: true,
		},

		"mac": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |
`,
		},

		"mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |
`,
		},

		// TagNodes

		"vif_c": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesEthernetVifSVifC{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSDhcpOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetVifSMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}
