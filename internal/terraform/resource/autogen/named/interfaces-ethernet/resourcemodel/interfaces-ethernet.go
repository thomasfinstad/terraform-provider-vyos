// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesEthernet describes the resource data model.
type InterfacesEthernet struct {
	// LeafNodes
	InterfacesEthernetAddress            customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesEthernetDescrIPtion        customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesEthernetDisableFlowControl customtypes.CustomStringValue `tfsdk:"disable_flow_control" json:"disable-flow-control,omitempty"`
	InterfacesEthernetDisableLinkDetect  customtypes.CustomStringValue `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	InterfacesEthernetDisable            customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesEthernetDuplex             customtypes.CustomStringValue `tfsdk:"duplex" json:"duplex,omitempty"`
	InterfacesEthernetHwID               customtypes.CustomStringValue `tfsdk:"hw_id" json:"hw-id,omitempty"`
	InterfacesEthernetMac                customtypes.CustomStringValue `tfsdk:"mac" json:"mac,omitempty"`
	InterfacesEthernetMtu                customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesEthernetSpeed              customtypes.CustomStringValue `tfsdk:"speed" json:"speed,omitempty"`
	InterfacesEthernetRedirect           customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesEthernetVrf                customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`
	InterfacesEthernetXdp                customtypes.CustomStringValue `tfsdk:"xdp" json:"xdp,omitempty"`

	// TagNodes
	InterfacesEthernetVifS types.Map `tfsdk:"vif_s" json:"vif-s,omitempty"`
	InterfacesEthernetVif  types.Map `tfsdk:"vif" json:"vif,omitempty"`

	// Nodes
	InterfacesEthernetDhcpOptions     types.Object `tfsdk:"dhcp_options" json:"dhcp-options,omitempty"`
	InterfacesEthernetDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options" json:"dhcpv6-options,omitempty"`
	InterfacesEthernetEapol           types.Object `tfsdk:"eapol" json:"eapol,omitempty"`
	InterfacesEthernetIP              types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesEthernetIPvsix          types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesEthernetMirror          types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
	InterfacesEthernetOffload         types.Object `tfsdk:"offload" json:"offload,omitempty"`
	InterfacesEthernetRingBuffer      types.Object `tfsdk:"ring_buffer" json:"ring-buffer,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesEthernet) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"disable_flow_control": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable Ethernet flow control (pause frames)

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

		"duplex": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Duplex mode

|  Format  |  Description  |
|----------|---------------|
|  auto  |  Auto negotiation  |
|  half  |  Half duplex  |
|  full  |  Full duplex  |
`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"hw_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Associate Ethernet Interface with given Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |
`,
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

		"speed": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Link speed

|  Format  |  Description  |
|----------|---------------|
|  auto  |  Auto negotiation  |
|  10  |  10 Mbit/sec  |
|  100  |  100 Mbit/sec  |
|  1000  |  1 Gbit/sec  |
|  2500  |  2.5 Gbit/sec  |
|  5000  |  5 Gbit/sec  |
|  10000  |  10 Gbit/sec  |
|  25000  |  25 Gbit/sec  |
|  40000  |  40 Gbit/sec  |
|  50000  |  50 Gbit/sec  |
|  100000  |  100 Gbit/sec  |
`,

			// Default:          stringdefault.StaticString(`auto`),
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

		"xdp": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable eXpress Data Path

`,
		},

		// TagNodes

		"vif_s": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesEthernetVifS{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |
`,
		},

		"vif": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesEthernetVif{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |
`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetDhcpOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetDhcpvsixOptions{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"eapol": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetEapol{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Extensible Authentication Protocol over Local Area Network

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"offload": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetOffload{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Configurable offload options

`,
		},

		"ring_buffer": schema.SingleNestedAttribute{
			Attributes: InterfacesEthernetRingBuffer{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Shared buffer between the device driver and NIC

`,
		},
	}
}
