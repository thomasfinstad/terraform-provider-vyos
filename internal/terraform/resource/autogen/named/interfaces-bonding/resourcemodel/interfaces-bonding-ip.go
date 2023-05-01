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

// InterfacesBondingIP describes the resource data model.
type InterfacesBondingIP struct {
	// LeafNodes
	LeafInterfacesBondingIPAdjustMss               types.String `tfsdk:"adjust_mss"`
	LeafInterfacesBondingIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout"`
	LeafInterfacesBondingIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter"`
	LeafInterfacesBondingIPDisableForwarding       types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesBondingIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast"`
	LeafInterfacesBondingIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept"`
	LeafInterfacesBondingIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce"`
	LeafInterfacesBondingIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore"`
	LeafInterfacesBondingIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp"`
	LeafInterfacesBondingIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan"`
	LeafInterfacesBondingIPSourceValIDation        types.String `tfsdk:"source_validation"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBondingIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bonding", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBondingIPAdjustMss.IsNull() || o.LeafInterfacesBondingIPAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesBondingIPAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesBondingIPArpCacheTimeout.IsNull() || o.LeafInterfacesBondingIPArpCacheTimeout.IsUnknown()) {
		vyosData["arp-cache-timeout"] = o.LeafInterfacesBondingIPArpCacheTimeout.ValueString()
	}
	if !(o.LeafInterfacesBondingIPDisableArpFilter.IsNull() || o.LeafInterfacesBondingIPDisableArpFilter.IsUnknown()) {
		vyosData["disable-arp-filter"] = o.LeafInterfacesBondingIPDisableArpFilter.ValueString()
	}
	if !(o.LeafInterfacesBondingIPDisableForwarding.IsNull() || o.LeafInterfacesBondingIPDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesBondingIPDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesBondingIPEnableDirectedBroadcast.IsNull() || o.LeafInterfacesBondingIPEnableDirectedBroadcast.IsUnknown()) {
		vyosData["enable-directed-broadcast"] = o.LeafInterfacesBondingIPEnableDirectedBroadcast.ValueString()
	}
	if !(o.LeafInterfacesBondingIPEnableArpAccept.IsNull() || o.LeafInterfacesBondingIPEnableArpAccept.IsUnknown()) {
		vyosData["enable-arp-accept"] = o.LeafInterfacesBondingIPEnableArpAccept.ValueString()
	}
	if !(o.LeafInterfacesBondingIPEnableArpAnnounce.IsNull() || o.LeafInterfacesBondingIPEnableArpAnnounce.IsUnknown()) {
		vyosData["enable-arp-announce"] = o.LeafInterfacesBondingIPEnableArpAnnounce.ValueString()
	}
	if !(o.LeafInterfacesBondingIPEnableArpIgnore.IsNull() || o.LeafInterfacesBondingIPEnableArpIgnore.IsUnknown()) {
		vyosData["enable-arp-ignore"] = o.LeafInterfacesBondingIPEnableArpIgnore.ValueString()
	}
	if !(o.LeafInterfacesBondingIPEnableProxyArp.IsNull() || o.LeafInterfacesBondingIPEnableProxyArp.IsUnknown()) {
		vyosData["enable-proxy-arp"] = o.LeafInterfacesBondingIPEnableProxyArp.ValueString()
	}
	if !(o.LeafInterfacesBondingIPProxyArpPvlan.IsNull() || o.LeafInterfacesBondingIPProxyArpPvlan.IsUnknown()) {
		vyosData["proxy-arp-pvlan"] = o.LeafInterfacesBondingIPProxyArpPvlan.ValueString()
	}
	if !(o.LeafInterfacesBondingIPSourceValIDation.IsNull() || o.LeafInterfacesBondingIPSourceValIDation.IsUnknown()) {
		vyosData["source-validation"] = o.LeafInterfacesBondingIPSourceValIDation.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBondingIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bonding", "ip"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesBondingIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["arp-cache-timeout"]; ok {
		o.LeafInterfacesBondingIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPArpCacheTimeout = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-arp-filter"]; ok {
		o.LeafInterfacesBondingIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPDisableArpFilter = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesBondingIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesBondingIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-accept"]; ok {
		o.LeafInterfacesBondingIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPEnableArpAccept = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-announce"]; ok {
		o.LeafInterfacesBondingIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPEnableArpAnnounce = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-ignore"]; ok {
		o.LeafInterfacesBondingIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPEnableArpIgnore = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-proxy-arp"]; ok {
		o.LeafInterfacesBondingIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPEnableProxyArp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesBondingIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPProxyArpPvlan = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-validation"]; ok {
		o.LeafInterfacesBondingIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bonding", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBondingIP) AttributeTypes() map[string]attr.Type {
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
func (o InterfacesBondingIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
