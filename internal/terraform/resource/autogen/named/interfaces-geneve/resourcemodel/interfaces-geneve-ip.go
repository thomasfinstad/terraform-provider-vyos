// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesGeneveIP describes the resource data model.
type InterfacesGeneveIP struct {
	// LeafNodes
	InterfacesGeneveIPAdjustMss               customtypes.CustomStringValue `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	InterfacesGeneveIPArpCacheTimeout         customtypes.CustomStringValue `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	InterfacesGeneveIPDisableArpFilter        customtypes.CustomStringValue `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	InterfacesGeneveIPDisableForwarding       customtypes.CustomStringValue `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	InterfacesGeneveIPEnableDirectedBroadcast customtypes.CustomStringValue `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	InterfacesGeneveIPEnableArpAccept         customtypes.CustomStringValue `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	InterfacesGeneveIPEnableArpAnnounce       customtypes.CustomStringValue `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	InterfacesGeneveIPEnableArpIgnore         customtypes.CustomStringValue `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	InterfacesGeneveIPEnableProxyArp          customtypes.CustomStringValue `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	InterfacesGeneveIPProxyArpPvlan           customtypes.CustomStringValue `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	InterfacesGeneveIPSourceValIDation        customtypes.CustomStringValue `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesGeneveIP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |
`,
		},

		"arp_cache_timeout": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ARP cache entry timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  ARP cache entry timout in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"disable_arp_filter": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable ARP filter on this interface

`,
		},

		"disable_forwarding": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		"enable_directed_broadcast": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
		},

		"enable_arp_accept": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable ARP accept on this interface

`,
		},

		"enable_arp_announce": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable ARP announce on this interface

`,
		},

		"enable_arp_ignore": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable ARP ignore on this interface

`,
		},

		"enable_proxy_arp": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable proxy-arp on this interface

`,
		},

		"proxy_arp_pvlan": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
		},

		"source_validation": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |
`,
		},

		// TagNodes

		// Nodes

	}
}
