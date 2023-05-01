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

// InterfacesWirelessIP describes the resource data model.
type InterfacesWirelessIP struct {
	// LeafNodes
	LeafInterfacesWirelessIPAdjustMss               types.String `tfsdk:"adjust_mss"`
	LeafInterfacesWirelessIPArpCacheTimeout         types.String `tfsdk:"arp_cache_timeout"`
	LeafInterfacesWirelessIPDisableArpFilter        types.String `tfsdk:"disable_arp_filter"`
	LeafInterfacesWirelessIPDisableForwarding       types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesWirelessIPEnableDirectedBroadcast types.String `tfsdk:"enable_directed_broadcast"`
	LeafInterfacesWirelessIPEnableArpAccept         types.String `tfsdk:"enable_arp_accept"`
	LeafInterfacesWirelessIPEnableArpAnnounce       types.String `tfsdk:"enable_arp_announce"`
	LeafInterfacesWirelessIPEnableArpIgnore         types.String `tfsdk:"enable_arp_ignore"`
	LeafInterfacesWirelessIPEnableProxyArp          types.String `tfsdk:"enable_proxy_arp"`
	LeafInterfacesWirelessIPProxyArpPvlan           types.String `tfsdk:"proxy_arp_pvlan"`
	LeafInterfacesWirelessIPSourceValIDation        types.String `tfsdk:"source_validation"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessIP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ip"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessIPAdjustMss.IsNull() || o.LeafInterfacesWirelessIPAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesWirelessIPAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPArpCacheTimeout.IsNull() || o.LeafInterfacesWirelessIPArpCacheTimeout.IsUnknown()) {
		vyosData["arp-cache-timeout"] = o.LeafInterfacesWirelessIPArpCacheTimeout.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPDisableArpFilter.IsNull() || o.LeafInterfacesWirelessIPDisableArpFilter.IsUnknown()) {
		vyosData["disable-arp-filter"] = o.LeafInterfacesWirelessIPDisableArpFilter.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPDisableForwarding.IsNull() || o.LeafInterfacesWirelessIPDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesWirelessIPDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPEnableDirectedBroadcast.IsNull() || o.LeafInterfacesWirelessIPEnableDirectedBroadcast.IsUnknown()) {
		vyosData["enable-directed-broadcast"] = o.LeafInterfacesWirelessIPEnableDirectedBroadcast.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPEnableArpAccept.IsNull() || o.LeafInterfacesWirelessIPEnableArpAccept.IsUnknown()) {
		vyosData["enable-arp-accept"] = o.LeafInterfacesWirelessIPEnableArpAccept.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPEnableArpAnnounce.IsNull() || o.LeafInterfacesWirelessIPEnableArpAnnounce.IsUnknown()) {
		vyosData["enable-arp-announce"] = o.LeafInterfacesWirelessIPEnableArpAnnounce.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPEnableArpIgnore.IsNull() || o.LeafInterfacesWirelessIPEnableArpIgnore.IsUnknown()) {
		vyosData["enable-arp-ignore"] = o.LeafInterfacesWirelessIPEnableArpIgnore.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPEnableProxyArp.IsNull() || o.LeafInterfacesWirelessIPEnableProxyArp.IsUnknown()) {
		vyosData["enable-proxy-arp"] = o.LeafInterfacesWirelessIPEnableProxyArp.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPProxyArpPvlan.IsNull() || o.LeafInterfacesWirelessIPProxyArpPvlan.IsUnknown()) {
		vyosData["proxy-arp-pvlan"] = o.LeafInterfacesWirelessIPProxyArpPvlan.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPSourceValIDation.IsNull() || o.LeafInterfacesWirelessIPSourceValIDation.IsUnknown()) {
		vyosData["source-validation"] = o.LeafInterfacesWirelessIPSourceValIDation.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessIP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ip"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesWirelessIPAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["arp-cache-timeout"]; ok {
		o.LeafInterfacesWirelessIPArpCacheTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPArpCacheTimeout = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-arp-filter"]; ok {
		o.LeafInterfacesWirelessIPDisableArpFilter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPDisableArpFilter = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesWirelessIPDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-directed-broadcast"]; ok {
		o.LeafInterfacesWirelessIPEnableDirectedBroadcast = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPEnableDirectedBroadcast = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-accept"]; ok {
		o.LeafInterfacesWirelessIPEnableArpAccept = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPEnableArpAccept = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-announce"]; ok {
		o.LeafInterfacesWirelessIPEnableArpAnnounce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPEnableArpAnnounce = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-arp-ignore"]; ok {
		o.LeafInterfacesWirelessIPEnableArpIgnore = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPEnableArpIgnore = basetypes.NewStringNull()
	}
	if value, ok := vyosData["enable-proxy-arp"]; ok {
		o.LeafInterfacesWirelessIPEnableProxyArp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPEnableProxyArp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["proxy-arp-pvlan"]; ok {
		o.LeafInterfacesWirelessIPProxyArpPvlan = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPProxyArpPvlan = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-validation"]; ok {
		o.LeafInterfacesWirelessIPSourceValIDation = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPSourceValIDation = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ip"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessIP) AttributeTypes() map[string]attr.Type {
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
func (o InterfacesWirelessIP) ResourceSchemaAttributes() map[string]schema.Attribute {
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
