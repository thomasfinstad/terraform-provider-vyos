// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceRouterAdvertInterface describes the resource data model.
type ServiceRouterAdvertInterface struct {
	// LeafNodes
	ServiceRouterAdvertInterfaceHopLimit           customtypes.CustomStringValue `tfsdk:"hop_limit" json:"hop-limit,omitempty"`
	ServiceRouterAdvertInterfaceDefaultLifetime    customtypes.CustomStringValue `tfsdk:"default_lifetime" json:"default-lifetime,omitempty"`
	ServiceRouterAdvertInterfaceDefaultPreference  customtypes.CustomStringValue `tfsdk:"default_preference" json:"default-preference,omitempty"`
	ServiceRouterAdvertInterfaceDNSsl              customtypes.CustomStringValue `tfsdk:"dnssl" json:"dnssl,omitempty"`
	ServiceRouterAdvertInterfaceLinkMtu            customtypes.CustomStringValue `tfsdk:"link_mtu" json:"link-mtu,omitempty"`
	ServiceRouterAdvertInterfaceManagedFlag        customtypes.CustomStringValue `tfsdk:"managed_flag" json:"managed-flag,omitempty"`
	ServiceRouterAdvertInterfaceNameServer         customtypes.CustomStringValue `tfsdk:"name_server" json:"name-server,omitempty"`
	ServiceRouterAdvertInterfaceNameServerLifetime customtypes.CustomStringValue `tfsdk:"name_server_lifetime" json:"name-server-lifetime,omitempty"`
	ServiceRouterAdvertInterfaceOtherConfigFlag    customtypes.CustomStringValue `tfsdk:"other_config_flag" json:"other-config-flag,omitempty"`
	ServiceRouterAdvertInterfaceSourceAddress      customtypes.CustomStringValue `tfsdk:"source_address" json:"source-address,omitempty"`
	ServiceRouterAdvertInterfaceReachableTime      customtypes.CustomStringValue `tfsdk:"reachable_time" json:"reachable-time,omitempty"`
	ServiceRouterAdvertInterfaceRetransTimer       customtypes.CustomStringValue `tfsdk:"retrans_timer" json:"retrans-timer,omitempty"`
	ServiceRouterAdvertInterfaceNoSendAdvert       customtypes.CustomStringValue `tfsdk:"no_send_advert" json:"no-send-advert,omitempty"`

	// TagNodes
	ServiceRouterAdvertInterfaceRoute  types.Map `tfsdk:"route" json:"route,omitempty"`
	ServiceRouterAdvertInterfacePrefix types.Map `tfsdk:"prefix" json:"prefix,omitempty"`

	// Nodes
	ServiceRouterAdvertInterfaceInterval types.Object `tfsdk:"interval" json:"interval,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceRouterAdvertInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hop_limit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set Hop Count field of the IP header for outgoing packets

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Unspecified (by this router)  |
|  u32:1-255  |  Value should represent current diameter of the Internet  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"default_lifetime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Lifetime associated with the default router in units of seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:4-9000  |  Router Lifetime in seconds  |
|  0  |  Not a default router  |
`,
		},

		"default_preference": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Preference associated with the default router,

|  Format  |  Description  |
|----------|---------------|
|  low  |  Default router has low preference  |
|  medium  |  Default router has medium preference  |
|  high  |  Default router has high preference  |
`,

			// Default:          stringdefault.StaticString(`medium`),
			Computed: true,
		},

		"dnssl": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DNS search list

`,
		},

		"link_mtu": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Link MTU value placed in RAs, exluded in RAs if unset

|  Format  |  Description  |
|----------|---------------|
|  u32:1280-9000  |  Link MTU value in RAs  |
`,
		},

		"managed_flag": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Hosts use the administered (stateful) protocol for address autoconfiguration in addition to any addresses autoconfigured using SLAAC

`,
		},

		"name_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |
`,
		},

		"name_server_lifetime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum duration how long the RDNSS entries are used

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Name-servers should no longer be used  |
|  u32:1-7200  |  Maximum interval in seconds  |
`,
		},

		"other_config_flag": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Hosts use the administered (stateful) protocol for autoconfiguration of other (non-address) information

`,
		},

		"source_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use IPv6 address as source address. Useful with VRRP.

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address to be advertized (must be configured on interface)  |
`,
		},

		"reachable_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time, in milliseconds, that a node assumes a neighbor is reachable after having received a reachability confirmation

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Reachable Time unspecified by this router  |
|  u32:1-3600000  |  Reachable Time value in RAs (in milliseconds)  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"retrans_timer": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time in milliseconds between retransmitted Neighbor Solicitation messages

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Time, in milliseconds, between retransmitted Neighbor Solicitation messages  |
|  u32:1-4294967295  |  Minimum interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"no_send_advert": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not send router adverts

`,
		},

		// TagNodes

		"route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceRouterAdvertInterfaceRoute{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 route to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 route to be advertized  |
`,
		},

		"prefix": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceRouterAdvertInterfacePrefix{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `IPv6 prefix to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 prefix to be advertized  |
`,
		},

		// Nodes

		"interval": schema.SingleNestedAttribute{
			Attributes: ServiceRouterAdvertInterfaceInterval{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Set interval between unsolicited multicast RAs

`,
		},
	}
}
