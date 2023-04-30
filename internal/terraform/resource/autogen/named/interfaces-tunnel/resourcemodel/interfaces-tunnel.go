// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesTunnel describes the resource data model.
type InterfacesTunnel struct {
	// LeafNodes
	InterfacesTunnelDescrIPtion       customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesTunnelAddress           customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesTunnelDisable           customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesTunnelDisableLinkDetect customtypes.CustomStringValue `tfsdk:"disable_link_detect" json:"disable-link-detect,omitempty"`
	InterfacesTunnelMtu               customtypes.CustomStringValue `tfsdk:"mtu" json:"mtu,omitempty"`
	InterfacesTunnelSourceAddress     customtypes.CustomStringValue `tfsdk:"source_address" json:"source-address,omitempty"`
	InterfacesTunnelRemote            customtypes.CustomStringValue `tfsdk:"remote" json:"remote,omitempty"`
	InterfacesTunnelSourceInterface   customtypes.CustomStringValue `tfsdk:"source_interface" json:"source-interface,omitempty"`
	InterfacesTunnelSixrdPrefix       customtypes.CustomStringValue `tfsdk:"6rd_prefix" json:"6rd-prefix,omitempty"`
	InterfacesTunnelSixrdRelayPrefix  customtypes.CustomStringValue `tfsdk:"6rd_relay_prefix" json:"6rd-relay-prefix,omitempty"`
	InterfacesTunnelEncapsulation     customtypes.CustomStringValue `tfsdk:"encapsulation" json:"encapsulation,omitempty"`
	InterfacesTunnelEnableMulticast   customtypes.CustomStringValue `tfsdk:"enable_multicast" json:"enable-multicast,omitempty"`
	InterfacesTunnelVrf               customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`
	InterfacesTunnelRedirect          customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`

	// TagNodes

	// Nodes
	InterfacesTunnelIP         types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesTunnelIPvsix     types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesTunnelMirror     types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
	InterfacesTunnelParameters types.Object `tfsdk:"parameters" json:"parameters,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesTunnel) ResourceAttributes() map[string]schema.Attribute {
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
`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"disable_link_detect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:64-8024  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1476`),
			Computed: true,
		},

		"source_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Source IP address used to initiate connection

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 source address  |
|  ipv6  |  IPv6 source address  |
`,
		},

		"remote": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Tunnel remote address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Tunnel remote IPv4 address  |
|  ipv6  |  Tunnel remote IPv6 address  |
`,
		},

		"source_interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface used to establish connection

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Interface name  |
`,
		},

		"6rd_prefix": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `6rd network prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address and prefix length  |
`,
		},

		"6rd_relay_prefix": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `6rd relay prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 prefix of interface for 6rd  |
`,
		},

		"encapsulation": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Encapsulation of this tunnel interface

|  Format  |  Description  |
|----------|---------------|
|  erspan  |  Encapsulated Remote Switched Port Analyzer  |
|  gre  |  Generic Routing Encapsulation (network layer)  |
|  gretap  |  Generic Routing Encapsulation (datalink layer)  |
|  ip6erspan  |  Encapsulated Remote Switched Port Analyzer over IPv6  |
|  ip6gre  |  GRE over IPv6 (network layer)  |
|  ip6gretap  |  GRE over IPv6 (datalink layer)  |
|  ip6ip6  |  IPv6 in IPv6 encapsulation  |
|  ipip  |  IPv4 in IPv4 encapsulation  |
|  ipip6  |  IPv4 in IP6 encapsulation  |
|  sit  |  Simple Internet Transition (IPv6 in IPv4)  |
`,
		},

		"enable_multicast": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable multicast operation over tunnel

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

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
		},

		// TagNodes

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParameters{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Tunnel parameters

`,
		},
	}
}
