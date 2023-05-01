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

// InterfacesBrIDgeVifIP describes the resource data model.
type InterfacesBrIDgeVifIP struct {
	// LeafNodes
	LeafInterfacesBrIDgeVifIPAdjustMss               types.String `tfsdk:"adjust_mss"`
	LeafInterfacesBrIDgeVifIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout"`
	LeafInterfacesBrIDgeVifIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter"`
	LeafInterfacesBrIDgeVifIPDisableForwarding       types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast"`
	LeafInterfacesBrIDgeVifIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept"`
	LeafInterfacesBrIDgeVifIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce"`
	LeafInterfacesBrIDgeVifIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore"`
	LeafInterfacesBrIDgeVifIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp"`
	LeafInterfacesBrIDgeVifIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan"`
	LeafInterfacesBrIDgeVifIPSourceValIDation        types.String `tfsdk:"source_validation"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBrIDgeVifIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBrIDgeVifIPAdjustMss.IsNull() || o.LeafInterfacesBrIDgeVifIPAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesBrIDgeVifIPAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPArpCacheTimeout.IsNull() || o.LeafInterfacesBrIDgeVifIPArpCacheTimeout.IsUnknown()) {
		vyosData["arp-cache-timeout"] = o.LeafInterfacesBrIDgeVifIPArpCacheTimeout.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPDisableArpFilter.IsNull() || o.LeafInterfacesBrIDgeVifIPDisableArpFilter.IsUnknown()) {
		vyosData["disable-arp-filter"] = o.LeafInterfacesBrIDgeVifIPDisableArpFilter.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPDisableForwarding.IsNull() || o.LeafInterfacesBrIDgeVifIPDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesBrIDgeVifIPDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast.IsNull() || o.LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast.IsUnknown()) {
		vyosData["enable-directed-broadcast"] = o.LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPEnableArpAccept.IsNull() || o.LeafInterfacesBrIDgeVifIPEnableArpAccept.IsUnknown()) {
		vyosData["enable-arp-accept"] = o.LeafInterfacesBrIDgeVifIPEnableArpAccept.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPEnableArpAnnounce.IsNull() || o.LeafInterfacesBrIDgeVifIPEnableArpAnnounce.IsUnknown()) {
		vyosData["enable-arp-announce"] = o.LeafInterfacesBrIDgeVifIPEnableArpAnnounce.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPEnableArpIgnore.IsNull() || o.LeafInterfacesBrIDgeVifIPEnableArpIgnore.IsUnknown()) {
		vyosData["enable-arp-ignore"] = o.LeafInterfacesBrIDgeVifIPEnableArpIgnore.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPEnableProxyArp.IsNull() || o.LeafInterfacesBrIDgeVifIPEnableProxyArp.IsUnknown()) {
		vyosData["enable-proxy-arp"] = o.LeafInterfacesBrIDgeVifIPEnableProxyArp.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPProxyArpPvlan.IsNull() || o.LeafInterfacesBrIDgeVifIPProxyArpPvlan.IsUnknown()) {
		vyosData["proxy-arp-pvlan"] = o.LeafInterfacesBrIDgeVifIPProxyArpPvlan.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeVifIPSourceValIDation.IsNull() || o.LeafInterfacesBrIDgeVifIPSourceValIDation.IsUnknown()) {
		vyosData["source-validation"] = o.LeafInterfacesBrIDgeVifIPSourceValIDation.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBrIDgeVifIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "ip"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesBrIDgeVifIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["arp-cache-timeout"]; ok {
		o.LeafInterfacesBrIDgeVifIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPArpCacheTimeout = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-arp-filter"]; ok {
		o.LeafInterfacesBrIDgeVifIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPDisableArpFilter = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesBrIDgeVifIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-accept"]; ok {
		o.LeafInterfacesBrIDgeVifIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPEnableArpAccept = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-announce"]; ok {
		o.LeafInterfacesBrIDgeVifIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPEnableArpAnnounce = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-ignore"]; ok {
		o.LeafInterfacesBrIDgeVifIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPEnableArpIgnore = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-proxy-arp"]; ok {
		o.LeafInterfacesBrIDgeVifIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPEnableProxyArp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesBrIDgeVifIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPProxyArpPvlan = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-validation"]; ok {
		o.LeafInterfacesBrIDgeVifIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeVifIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bridge", "vif", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBrIDgeVifIP) AttributeTypes() map[string]attr.Type {
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
func (o InterfacesBrIDgeVifIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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