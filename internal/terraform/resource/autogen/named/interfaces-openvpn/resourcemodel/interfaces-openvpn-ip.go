// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesOpenvpnIP describes the resource data model.
type InterfacesOpenvpnIP struct {
	// LeafNodes
	LeafInterfacesOpenvpnIPAdjustMss               types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesOpenvpnIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	LeafInterfacesOpenvpnIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	LeafInterfacesOpenvpnIPDisableForwarding       types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesOpenvpnIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	LeafInterfacesOpenvpnIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	LeafInterfacesOpenvpnIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	LeafInterfacesOpenvpnIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	LeafInterfacesOpenvpnIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	LeafInterfacesOpenvpnIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesOpenvpnIPSourceValIDation        types.String `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesOpenvpnIP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesOpenvpnIPAdjustMss.IsNull() && !o.LeafInterfacesOpenvpnIPAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesOpenvpnIPAdjustMss.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPArpCacheTimeout.IsNull() && !o.LeafInterfacesOpenvpnIPArpCacheTimeout.IsUnknown() {
		jsonData["arp-cache-timeout"] = o.LeafInterfacesOpenvpnIPArpCacheTimeout.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPDisableArpFilter.IsNull() && !o.LeafInterfacesOpenvpnIPDisableArpFilter.IsUnknown() {
		jsonData["disable-arp-filter"] = o.LeafInterfacesOpenvpnIPDisableArpFilter.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPDisableForwarding.IsNull() && !o.LeafInterfacesOpenvpnIPDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesOpenvpnIPDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPEnableDirectedBroadcast.IsNull() && !o.LeafInterfacesOpenvpnIPEnableDirectedBroadcast.IsUnknown() {
		jsonData["enable-directed-broadcast"] = o.LeafInterfacesOpenvpnIPEnableDirectedBroadcast.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPEnableArpAccept.IsNull() && !o.LeafInterfacesOpenvpnIPEnableArpAccept.IsUnknown() {
		jsonData["enable-arp-accept"] = o.LeafInterfacesOpenvpnIPEnableArpAccept.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPEnableArpAnnounce.IsNull() && !o.LeafInterfacesOpenvpnIPEnableArpAnnounce.IsUnknown() {
		jsonData["enable-arp-announce"] = o.LeafInterfacesOpenvpnIPEnableArpAnnounce.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPEnableArpIgnore.IsNull() && !o.LeafInterfacesOpenvpnIPEnableArpIgnore.IsUnknown() {
		jsonData["enable-arp-ignore"] = o.LeafInterfacesOpenvpnIPEnableArpIgnore.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPEnableProxyArp.IsNull() && !o.LeafInterfacesOpenvpnIPEnableProxyArp.IsUnknown() {
		jsonData["enable-proxy-arp"] = o.LeafInterfacesOpenvpnIPEnableProxyArp.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPProxyArpPvlan.IsNull() && !o.LeafInterfacesOpenvpnIPProxyArpPvlan.IsUnknown() {
		jsonData["proxy-arp-pvlan"] = o.LeafInterfacesOpenvpnIPProxyArpPvlan.ValueString()
	}

	if !o.LeafInterfacesOpenvpnIPSourceValIDation.IsNull() && !o.LeafInterfacesOpenvpnIPSourceValIDation.IsUnknown() {
		jsonData["source-validation"] = o.LeafInterfacesOpenvpnIPSourceValIDation.ValueString()
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
func (o *InterfacesOpenvpnIP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesOpenvpnIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["arp-cache-timeout"]; ok {
		o.LeafInterfacesOpenvpnIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPArpCacheTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-arp-filter"]; ok {
		o.LeafInterfacesOpenvpnIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPDisableArpFilter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesOpenvpnIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesOpenvpnIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-accept"]; ok {
		o.LeafInterfacesOpenvpnIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPEnableArpAccept = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-announce"]; ok {
		o.LeafInterfacesOpenvpnIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPEnableArpAnnounce = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-ignore"]; ok {
		o.LeafInterfacesOpenvpnIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPEnableArpIgnore = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-proxy-arp"]; ok {
		o.LeafInterfacesOpenvpnIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPEnableProxyArp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesOpenvpnIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPProxyArpPvlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-validation"]; ok {
		o.LeafInterfacesOpenvpnIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesOpenvpnIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
