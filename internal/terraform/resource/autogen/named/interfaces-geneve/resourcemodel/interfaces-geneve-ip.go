// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// InterfacesGeneveIP describes the resource data model.
type InterfacesGeneveIP struct {
	// LeafNodes
	LeafInterfacesGeneveIPAdjustMss               types.String `tfsdk:"adjust_mss" json:"adjust-mss,omitempty"`
	LeafInterfacesGeneveIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout" json:"arp-cache-timeout,omitempty"`
	LeafInterfacesGeneveIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter" json:"disable-arp-filter,omitempty"`
	LeafInterfacesGeneveIPDisableForwarding       types.String `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`
	LeafInterfacesGeneveIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast" json:"enable-directed-broadcast,omitempty"`
	LeafInterfacesGeneveIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept" json:"enable-arp-accept,omitempty"`
	LeafInterfacesGeneveIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce" json:"enable-arp-announce,omitempty"`
	LeafInterfacesGeneveIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore" json:"enable-arp-ignore,omitempty"`
	LeafInterfacesGeneveIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp" json:"enable-proxy-arp,omitempty"`
	LeafInterfacesGeneveIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan" json:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesGeneveIPSourceValIDation        types.String `tfsdk:"source_validation" json:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesGeneveIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *InterfacesGeneveIP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafInterfacesGeneveIPAdjustMss.IsNull() && !o.LeafInterfacesGeneveIPAdjustMss.IsUnknown() {
		jsonData["adjust-mss"] = o.LeafInterfacesGeneveIPAdjustMss.ValueString()
	}

	if !o.LeafInterfacesGeneveIPArpCacheTimeout.IsNull() && !o.LeafInterfacesGeneveIPArpCacheTimeout.IsUnknown() {
		jsonData["arp-cache-timeout"] = o.LeafInterfacesGeneveIPArpCacheTimeout.ValueString()
	}

	if !o.LeafInterfacesGeneveIPDisableArpFilter.IsNull() && !o.LeafInterfacesGeneveIPDisableArpFilter.IsUnknown() {
		jsonData["disable-arp-filter"] = o.LeafInterfacesGeneveIPDisableArpFilter.ValueString()
	}

	if !o.LeafInterfacesGeneveIPDisableForwarding.IsNull() && !o.LeafInterfacesGeneveIPDisableForwarding.IsUnknown() {
		jsonData["disable-forwarding"] = o.LeafInterfacesGeneveIPDisableForwarding.ValueString()
	}

	if !o.LeafInterfacesGeneveIPEnableDirectedBroadcast.IsNull() && !o.LeafInterfacesGeneveIPEnableDirectedBroadcast.IsUnknown() {
		jsonData["enable-directed-broadcast"] = o.LeafInterfacesGeneveIPEnableDirectedBroadcast.ValueString()
	}

	if !o.LeafInterfacesGeneveIPEnableArpAccept.IsNull() && !o.LeafInterfacesGeneveIPEnableArpAccept.IsUnknown() {
		jsonData["enable-arp-accept"] = o.LeafInterfacesGeneveIPEnableArpAccept.ValueString()
	}

	if !o.LeafInterfacesGeneveIPEnableArpAnnounce.IsNull() && !o.LeafInterfacesGeneveIPEnableArpAnnounce.IsUnknown() {
		jsonData["enable-arp-announce"] = o.LeafInterfacesGeneveIPEnableArpAnnounce.ValueString()
	}

	if !o.LeafInterfacesGeneveIPEnableArpIgnore.IsNull() && !o.LeafInterfacesGeneveIPEnableArpIgnore.IsUnknown() {
		jsonData["enable-arp-ignore"] = o.LeafInterfacesGeneveIPEnableArpIgnore.ValueString()
	}

	if !o.LeafInterfacesGeneveIPEnableProxyArp.IsNull() && !o.LeafInterfacesGeneveIPEnableProxyArp.IsUnknown() {
		jsonData["enable-proxy-arp"] = o.LeafInterfacesGeneveIPEnableProxyArp.ValueString()
	}

	if !o.LeafInterfacesGeneveIPProxyArpPvlan.IsNull() && !o.LeafInterfacesGeneveIPProxyArpPvlan.IsUnknown() {
		jsonData["proxy-arp-pvlan"] = o.LeafInterfacesGeneveIPProxyArpPvlan.ValueString()
	}

	if !o.LeafInterfacesGeneveIPSourceValIDation.IsNull() && !o.LeafInterfacesGeneveIPSourceValIDation.IsUnknown() {
		jsonData["source-validation"] = o.LeafInterfacesGeneveIPSourceValIDation.ValueString()
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
func (o *InterfacesGeneveIP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["adjust-mss"]; ok {
		o.LeafInterfacesGeneveIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPAdjustMss = basetypes.NewStringNull()
	}

	if value, ok := jsonData["arp-cache-timeout"]; ok {
		o.LeafInterfacesGeneveIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPArpCacheTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-arp-filter"]; ok {
		o.LeafInterfacesGeneveIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPDisableArpFilter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable-forwarding"]; ok {
		o.LeafInterfacesGeneveIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPDisableForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesGeneveIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-accept"]; ok {
		o.LeafInterfacesGeneveIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPEnableArpAccept = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-announce"]; ok {
		o.LeafInterfacesGeneveIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPEnableArpAnnounce = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-arp-ignore"]; ok {
		o.LeafInterfacesGeneveIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPEnableArpIgnore = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-proxy-arp"]; ok {
		o.LeafInterfacesGeneveIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPEnableProxyArp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesGeneveIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPProxyArpPvlan = basetypes.NewStringNull()
	}

	if value, ok := jsonData["source-validation"]; ok {
		o.LeafInterfacesGeneveIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesGeneveIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
