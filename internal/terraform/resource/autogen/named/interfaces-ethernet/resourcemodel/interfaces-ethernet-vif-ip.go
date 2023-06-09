// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesEthernetVifIP describes the resource data model.
type InterfacesEthernetVifIP struct {
	// LeafNodes
	LeafInterfacesEthernetVifIPAdjustMss               types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesEthernetVifIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	LeafInterfacesEthernetVifIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	LeafInterfacesEthernetVifIPDisableForwarding       types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesEthernetVifIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	LeafInterfacesEthernetVifIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	LeafInterfacesEthernetVifIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	LeafInterfacesEthernetVifIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	LeafInterfacesEthernetVifIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	LeafInterfacesEthernetVifIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesEthernetVifIPSourceValIDation        types.String `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |

`,
		},

		"arp_cache_timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ARP cache entry timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  ARP cache entry timout in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"disable_arp_filter": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable ARP filter on this interface

`,
		},

		"disable_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		"enable_directed_broadcast": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
		},

		"enable_arp_accept": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP accept on this interface

`,
		},

		"enable_arp_announce": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP announce on this interface

`,
		},

		"enable_arp_ignore": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP ignore on this interface

`,
		},

		"enable_proxy_arp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable proxy-arp on this interface

`,
		},

		"proxy_arp_pvlan": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
		},

		"source_validation": schema.StringAttribute{
			Optional: true,
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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifIP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesEthernetVifIPAdjustMss.IsNull() && !o.LeafInterfacesEthernetVifIPAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesEthernetVifIPAdjustMss.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPArpCacheTimeout.IsNull() && !o.LeafInterfacesEthernetVifIPArpCacheTimeout.IsUnknown() {
		jsonData["arp-cache-timeout"] = o.LeafInterfacesEthernetVifIPArpCacheTimeout.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPDisableArpFilter.IsNull() && !o.LeafInterfacesEthernetVifIPDisableArpFilter.IsUnknown() {
		jsonData["disable-arp-filter"] = o.LeafInterfacesEthernetVifIPDisableArpFilter.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPDisableForwarding.IsNull() && !o.LeafInterfacesEthernetVifIPDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesEthernetVifIPDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPEnableDirectedBroadcast.IsNull() && !o.LeafInterfacesEthernetVifIPEnableDirectedBroadcast.IsUnknown() {
		jsonData["enable-directed-broadcast"] = o.LeafInterfacesEthernetVifIPEnableDirectedBroadcast.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPEnableArpAccept.IsNull() && !o.LeafInterfacesEthernetVifIPEnableArpAccept.IsUnknown() {
		jsonData["enable-arp-accept"] = o.LeafInterfacesEthernetVifIPEnableArpAccept.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPEnableArpAnnounce.IsNull() && !o.LeafInterfacesEthernetVifIPEnableArpAnnounce.IsUnknown() {
		jsonData["enable-arp-announce"] = o.LeafInterfacesEthernetVifIPEnableArpAnnounce.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPEnableArpIgnore.IsNull() && !o.LeafInterfacesEthernetVifIPEnableArpIgnore.IsUnknown() {
		jsonData["enable-arp-ignore"] = o.LeafInterfacesEthernetVifIPEnableArpIgnore.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPEnableProxyArp.IsNull() && !o.LeafInterfacesEthernetVifIPEnableProxyArp.IsUnknown() {
		jsonData["enable-proxy-arp"] = o.LeafInterfacesEthernetVifIPEnableProxyArp.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPProxyArpPvlan.IsNull() && !o.LeafInterfacesEthernetVifIPProxyArpPvlan.IsUnknown() {
		jsonData["proxy-arp-pvlan"] = o.LeafInterfacesEthernetVifIPProxyArpPvlan.ValueString()
	}

	if !o.LeafInterfacesEthernetVifIPSourceValIDation.IsNull() && !o.LeafInterfacesEthernetVifIPSourceValIDation.IsUnknown() {
		jsonData["source-validation"] = o.LeafInterfacesEthernetVifIPSourceValIDation.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVifIP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesEthernetVifIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["arp-cache-timeout"]; ok {
		o.LeafInterfacesEthernetVifIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPArpCacheTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-arp-filter"]; ok {
		o.LeafInterfacesEthernetVifIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPDisableArpFilter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesEthernetVifIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesEthernetVifIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-accept"]; ok {
		o.LeafInterfacesEthernetVifIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPEnableArpAccept = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-announce"]; ok {
		o.LeafInterfacesEthernetVifIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPEnableArpAnnounce = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-ignore"]; ok {
		o.LeafInterfacesEthernetVifIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPEnableArpIgnore = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-proxy-arp"]; ok {
		o.LeafInterfacesEthernetVifIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPEnableProxyArp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesEthernetVifIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPProxyArpPvlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-validation"]; ok {
		o.LeafInterfacesEthernetVifIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
