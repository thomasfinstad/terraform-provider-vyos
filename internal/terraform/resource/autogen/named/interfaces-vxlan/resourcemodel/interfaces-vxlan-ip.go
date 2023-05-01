// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesVxlanIP describes the resource data model.
type InterfacesVxlanIP struct {
	// LeafNodes
	LeafInterfacesVxlanIPAdjustMss               types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesVxlanIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	LeafInterfacesVxlanIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	LeafInterfacesVxlanIPDisableForwarding       types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesVxlanIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	LeafInterfacesVxlanIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	LeafInterfacesVxlanIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	LeafInterfacesVxlanIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	LeafInterfacesVxlanIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	LeafInterfacesVxlanIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesVxlanIPSourceValIDation        types.String `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVxlanIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesVxlanIP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesVxlanIPAdjustMss.IsNull() && !o.LeafInterfacesVxlanIPAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesVxlanIPAdjustMss.ValueString()
	}

	if !o.LeafInterfacesVxlanIPArpCacheTimeout.IsNull() && !o.LeafInterfacesVxlanIPArpCacheTimeout.IsUnknown() {
		jsonData["arp-cache-timeout"] = o.LeafInterfacesVxlanIPArpCacheTimeout.ValueString()
	}

	if !o.LeafInterfacesVxlanIPDisableArpFilter.IsNull() && !o.LeafInterfacesVxlanIPDisableArpFilter.IsUnknown() {
		jsonData["disable-arp-filter"] = o.LeafInterfacesVxlanIPDisableArpFilter.ValueString()
	}

	if !o.LeafInterfacesVxlanIPDisableForwarding.IsNull() && !o.LeafInterfacesVxlanIPDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesVxlanIPDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesVxlanIPEnableDirectedBroadcast.IsNull() && !o.LeafInterfacesVxlanIPEnableDirectedBroadcast.IsUnknown() {
		jsonData["enable-directed-broadcast"] = o.LeafInterfacesVxlanIPEnableDirectedBroadcast.ValueString()
	}

	if !o.LeafInterfacesVxlanIPEnableArpAccept.IsNull() && !o.LeafInterfacesVxlanIPEnableArpAccept.IsUnknown() {
		jsonData["enable-arp-accept"] = o.LeafInterfacesVxlanIPEnableArpAccept.ValueString()
	}

	if !o.LeafInterfacesVxlanIPEnableArpAnnounce.IsNull() && !o.LeafInterfacesVxlanIPEnableArpAnnounce.IsUnknown() {
		jsonData["enable-arp-announce"] = o.LeafInterfacesVxlanIPEnableArpAnnounce.ValueString()
	}

	if !o.LeafInterfacesVxlanIPEnableArpIgnore.IsNull() && !o.LeafInterfacesVxlanIPEnableArpIgnore.IsUnknown() {
		jsonData["enable-arp-ignore"] = o.LeafInterfacesVxlanIPEnableArpIgnore.ValueString()
	}

	if !o.LeafInterfacesVxlanIPEnableProxyArp.IsNull() && !o.LeafInterfacesVxlanIPEnableProxyArp.IsUnknown() {
		jsonData["enable-proxy-arp"] = o.LeafInterfacesVxlanIPEnableProxyArp.ValueString()
	}

	if !o.LeafInterfacesVxlanIPProxyArpPvlan.IsNull() && !o.LeafInterfacesVxlanIPProxyArpPvlan.IsUnknown() {
		jsonData["proxy-arp-pvlan"] = o.LeafInterfacesVxlanIPProxyArpPvlan.ValueString()
	}

	if !o.LeafInterfacesVxlanIPSourceValIDation.IsNull() && !o.LeafInterfacesVxlanIPSourceValIDation.IsUnknown() {
		jsonData["source-validation"] = o.LeafInterfacesVxlanIPSourceValIDation.ValueString()
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
func (o *InterfacesVxlanIP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesVxlanIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["arp-cache-timeout"]; ok {
		o.LeafInterfacesVxlanIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPArpCacheTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-arp-filter"]; ok {
		o.LeafInterfacesVxlanIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPDisableArpFilter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesVxlanIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesVxlanIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-accept"]; ok {
		o.LeafInterfacesVxlanIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPEnableArpAccept = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-announce"]; ok {
		o.LeafInterfacesVxlanIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPEnableArpAnnounce = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-ignore"]; ok {
		o.LeafInterfacesVxlanIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPEnableArpIgnore = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-proxy-arp"]; ok {
		o.LeafInterfacesVxlanIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPEnableProxyArp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesVxlanIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPProxyArpPvlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-validation"]; ok {
		o.LeafInterfacesVxlanIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesVxlanIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
