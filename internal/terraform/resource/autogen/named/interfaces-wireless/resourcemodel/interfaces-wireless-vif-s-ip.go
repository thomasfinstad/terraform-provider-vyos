// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InterfacesWirelessVifSIP describes the resource data model.
type InterfacesWirelessVifSIP struct {
	// LeafNodes
	LeafInterfacesWirelessVifSIPAdjustMss               types.String `tfsdk:"adjust_mss"`
	LeafInterfacesWirelessVifSIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout"`
	LeafInterfacesWirelessVifSIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter"`
	LeafInterfacesWirelessVifSIPDisableForwarding       types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesWirelessVifSIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast"`
	LeafInterfacesWirelessVifSIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept"`
	LeafInterfacesWirelessVifSIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce"`
	LeafInterfacesWirelessVifSIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore"`
	LeafInterfacesWirelessVifSIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp"`
	LeafInterfacesWirelessVifSIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan"`
	LeafInterfacesWirelessVifSIPSourceValIDation        types.String `tfsdk:"source_validation"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessVifSIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessVifSIPAdjustMss.IsNull() || o.LeafInterfacesWirelessVifSIPAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesWirelessVifSIPAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPArpCacheTimeout.IsNull() || o.LeafInterfacesWirelessVifSIPArpCacheTimeout.IsUnknown()) {
		vyosData["arp-cache-timeout"] = o.LeafInterfacesWirelessVifSIPArpCacheTimeout.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPDisableArpFilter.IsNull() || o.LeafInterfacesWirelessVifSIPDisableArpFilter.IsUnknown()) {
		vyosData["disable-arp-filter"] = o.LeafInterfacesWirelessVifSIPDisableArpFilter.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPDisableForwarding.IsNull() || o.LeafInterfacesWirelessVifSIPDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesWirelessVifSIPDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPEnableDirectedBroadcast.IsNull() || o.LeafInterfacesWirelessVifSIPEnableDirectedBroadcast.IsUnknown()) {
		vyosData["enable-directed-broadcast"] = o.LeafInterfacesWirelessVifSIPEnableDirectedBroadcast.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPEnableArpAccept.IsNull() || o.LeafInterfacesWirelessVifSIPEnableArpAccept.IsUnknown()) {
		vyosData["enable-arp-accept"] = o.LeafInterfacesWirelessVifSIPEnableArpAccept.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPEnableArpAnnounce.IsNull() || o.LeafInterfacesWirelessVifSIPEnableArpAnnounce.IsUnknown()) {
		vyosData["enable-arp-announce"] = o.LeafInterfacesWirelessVifSIPEnableArpAnnounce.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPEnableArpIgnore.IsNull() || o.LeafInterfacesWirelessVifSIPEnableArpIgnore.IsUnknown()) {
		vyosData["enable-arp-ignore"] = o.LeafInterfacesWirelessVifSIPEnableArpIgnore.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPEnableProxyArp.IsNull() || o.LeafInterfacesWirelessVifSIPEnableProxyArp.IsUnknown()) {
		vyosData["enable-proxy-arp"] = o.LeafInterfacesWirelessVifSIPEnableProxyArp.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPProxyArpPvlan.IsNull() || o.LeafInterfacesWirelessVifSIPProxyArpPvlan.IsUnknown()) {
		vyosData["proxy-arp-pvlan"] = o.LeafInterfacesWirelessVifSIPProxyArpPvlan.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSIPSourceValIDation.IsNull() || o.LeafInterfacesWirelessVifSIPSourceValIDation.IsUnknown()) {
		vyosData["source-validation"] = o.LeafInterfacesWirelessVifSIPSourceValIDation.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessVifSIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "ip"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesWirelessVifSIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["arp-cache-timeout"]; ok {
		o.LeafInterfacesWirelessVifSIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPArpCacheTimeout = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-arp-filter"]; ok {
		o.LeafInterfacesWirelessVifSIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPDisableArpFilter = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesWirelessVifSIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesWirelessVifSIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-accept"]; ok {
		o.LeafInterfacesWirelessVifSIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPEnableArpAccept = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-announce"]; ok {
		o.LeafInterfacesWirelessVifSIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPEnableArpAnnounce = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-ignore"]; ok {
		o.LeafInterfacesWirelessVifSIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPEnableArpIgnore = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-proxy-arp"]; ok {
		o.LeafInterfacesWirelessVifSIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPEnableProxyArp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesWirelessVifSIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPProxyArpPvlan = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-validation"]; ok {
		o.LeafInterfacesWirelessVifSIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessVifSIP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"adjust_mss":                types.StringType,
		"arp_cache_timeout":         types.StringType,
		"disable_arp_filter":        types.StringType,
		"disable_forwarding":        types.StringType,
		"enable_directed_broadcast": types.StringType,
		"enable_arp_accept":         types.StringType,
		"enable_arp_announce":       types.StringType,
		"enable_arp_ignore":         types.StringType,
		"enable_proxy_arp":          types.StringType,
		"proxy_arp_pvlan":           types.StringType,
		"source_validation":         types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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