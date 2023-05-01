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

// VpnIPsecSiteToSitePeerTunnelLocal describes the resource data model.
type VpnIPsecSiteToSitePeerTunnelLocal struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerTunnelLocalPort   types.String `tfsdk:"port"`
	LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix types.String `tfsdk:"prefix"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecSiteToSitePeerTunnelLocal) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "tunnel", "local"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecSiteToSitePeerTunnelLocalPort.IsNull() || o.LeafVpnIPsecSiteToSitePeerTunnelLocalPort.IsUnknown()) {
		vyosData["port"] = o.LeafVpnIPsecSiteToSitePeerTunnelLocalPort.ValueString()
	}
	if !(o.LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix.IsNull() || o.LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix.IsUnknown()) {
		vyosData["prefix"] = o.LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecSiteToSitePeerTunnelLocal) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "tunnel", "local"}})

	// Leafs
	if value, ok := vyosData["port"]; ok {
		o.LeafVpnIPsecSiteToSitePeerTunnelLocalPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerTunnelLocalPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix"]; ok {
		o.LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "tunnel", "local"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecSiteToSitePeerTunnelLocal) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"port":   types.StringType,
		"prefix": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerTunnelLocal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv4 or IPv6 prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 prefix  |
|  ipv6net  |  Local IPv6 prefix  |

`,
		},

		// TagNodes

		// Nodes

	}
}
