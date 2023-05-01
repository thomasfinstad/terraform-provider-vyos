// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesBondingVifSVifCIP describes the resource data model.
type InterfacesBondingVifSVifCIP struct {
	// LeafNodes
	LeafInterfacesBondingVifSVifCIPAdjustMss               types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesBondingVifSVifCIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	LeafInterfacesBondingVifSVifCIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	LeafInterfacesBondingVifSVifCIPDisableForwarding       types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	LeafInterfacesBondingVifSVifCIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	LeafInterfacesBondingVifSVifCIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	LeafInterfacesBondingVifSVifCIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	LeafInterfacesBondingVifSVifCIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	LeafInterfacesBondingVifSVifCIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesBondingVifSVifCIPSourceValIDation        types.String `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSVifCIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesBondingVifSVifCIP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesBondingVifSVifCIPAdjustMss.IsNull() && !o.LeafInterfacesBondingVifSVifCIPAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesBondingVifSVifCIPAdjustMss.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPArpCacheTimeout.IsNull() && !o.LeafInterfacesBondingVifSVifCIPArpCacheTimeout.IsUnknown() {
		jsonData["arp-cache-timeout"] = o.LeafInterfacesBondingVifSVifCIPArpCacheTimeout.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPDisableArpFilter.IsNull() && !o.LeafInterfacesBondingVifSVifCIPDisableArpFilter.IsUnknown() {
		jsonData["disable-arp-filter"] = o.LeafInterfacesBondingVifSVifCIPDisableArpFilter.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPDisableForwarding.IsNull() && !o.LeafInterfacesBondingVifSVifCIPDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesBondingVifSVifCIPDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast.IsNull() && !o.LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast.IsUnknown() {
		jsonData["enable-directed-broadcast"] = o.LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPEnableArpAccept.IsNull() && !o.LeafInterfacesBondingVifSVifCIPEnableArpAccept.IsUnknown() {
		jsonData["enable-arp-accept"] = o.LeafInterfacesBondingVifSVifCIPEnableArpAccept.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPEnableArpAnnounce.IsNull() && !o.LeafInterfacesBondingVifSVifCIPEnableArpAnnounce.IsUnknown() {
		jsonData["enable-arp-announce"] = o.LeafInterfacesBondingVifSVifCIPEnableArpAnnounce.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPEnableArpIgnore.IsNull() && !o.LeafInterfacesBondingVifSVifCIPEnableArpIgnore.IsUnknown() {
		jsonData["enable-arp-ignore"] = o.LeafInterfacesBondingVifSVifCIPEnableArpIgnore.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPEnableProxyArp.IsNull() && !o.LeafInterfacesBondingVifSVifCIPEnableProxyArp.IsUnknown() {
		jsonData["enable-proxy-arp"] = o.LeafInterfacesBondingVifSVifCIPEnableProxyArp.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPProxyArpPvlan.IsNull() && !o.LeafInterfacesBondingVifSVifCIPProxyArpPvlan.IsUnknown() {
		jsonData["proxy-arp-pvlan"] = o.LeafInterfacesBondingVifSVifCIPProxyArpPvlan.ValueString()
	}

	if !o.LeafInterfacesBondingVifSVifCIPSourceValIDation.IsNull() && !o.LeafInterfacesBondingVifSVifCIPSourceValIDation.IsUnknown() {
		jsonData["source-validation"] = o.LeafInterfacesBondingVifSVifCIPSourceValIDation.ValueString()
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
func (o *InterfacesBondingVifSVifCIP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesBondingVifSVifCIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["arp-cache-timeout"]; ok {
		o.LeafInterfacesBondingVifSVifCIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPArpCacheTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-arp-filter"]; ok {
		o.LeafInterfacesBondingVifSVifCIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPDisableArpFilter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesBondingVifSVifCIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-accept"]; ok {
		o.LeafInterfacesBondingVifSVifCIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPEnableArpAccept = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-announce"]; ok {
		o.LeafInterfacesBondingVifSVifCIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPEnableArpAnnounce = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-ignore"]; ok {
		o.LeafInterfacesBondingVifSVifCIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPEnableArpIgnore = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-proxy-arp"]; ok {
		o.LeafInterfacesBondingVifSVifCIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPEnableProxyArp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesBondingVifSVifCIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPProxyArpPvlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-validation"]; ok {
		o.LeafInterfacesBondingVifSVifCIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifSVifCIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
