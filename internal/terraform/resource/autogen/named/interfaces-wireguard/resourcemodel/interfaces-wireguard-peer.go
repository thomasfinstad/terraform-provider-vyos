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

// InterfacesWireguardPeer describes the resource data model.
type InterfacesWireguardPeer struct {
	// LeafNodes
	LeafInterfacesWireguardPeerDisable             types.String `tfsdk:"disable"`
	LeafInterfacesWireguardPeerPublicKey           types.String `tfsdk:"public_key"`
	LeafInterfacesWireguardPeerPresharedKey        types.String `tfsdk:"preshared_key"`
	LeafInterfacesWireguardPeerAllowedIPs          types.String `tfsdk:"allowed_ips"`
	LeafInterfacesWireguardPeerAddress             types.String `tfsdk:"address"`
	LeafInterfacesWireguardPeerPort                types.String `tfsdk:"port"`
	LeafInterfacesWireguardPeerPersistentKeepalive types.String `tfsdk:"persistent_keepalive"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWireguardPeer) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireguard", "peer"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWireguardPeerDisable.IsNull() || o.LeafInterfacesWireguardPeerDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesWireguardPeerDisable.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerPublicKey.IsNull() || o.LeafInterfacesWireguardPeerPublicKey.IsUnknown()) {
		vyosData["public-key"] = o.LeafInterfacesWireguardPeerPublicKey.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerPresharedKey.IsNull() || o.LeafInterfacesWireguardPeerPresharedKey.IsUnknown()) {
		vyosData["preshared-key"] = o.LeafInterfacesWireguardPeerPresharedKey.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerAllowedIPs.IsNull() || o.LeafInterfacesWireguardPeerAllowedIPs.IsUnknown()) {
		vyosData["allowed-ips"] = o.LeafInterfacesWireguardPeerAllowedIPs.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerAddress.IsNull() || o.LeafInterfacesWireguardPeerAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesWireguardPeerAddress.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerPort.IsNull() || o.LeafInterfacesWireguardPeerPort.IsUnknown()) {
		vyosData["port"] = o.LeafInterfacesWireguardPeerPort.ValueString()
	}
	if !(o.LeafInterfacesWireguardPeerPersistentKeepalive.IsNull() || o.LeafInterfacesWireguardPeerPersistentKeepalive.IsUnknown()) {
		vyosData["persistent-keepalive"] = o.LeafInterfacesWireguardPeerPersistentKeepalive.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWireguardPeer) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireguard", "peer"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesWireguardPeerDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["public-key"]; ok {
		o.LeafInterfacesWireguardPeerPublicKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerPublicKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["preshared-key"]; ok {
		o.LeafInterfacesWireguardPeerPresharedKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerPresharedKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["allowed-ips"]; ok {
		o.LeafInterfacesWireguardPeerAllowedIPs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerAllowedIPs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesWireguardPeerAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafInterfacesWireguardPeerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["persistent-keepalive"]; ok {
		o.LeafInterfacesWireguardPeerPersistentKeepalive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWireguardPeerPersistentKeepalive = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireguard", "peer"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWireguardPeer) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":              types.StringType,
		"public_key":           types.StringType,
		"preshared_key":        types.StringType,
		"allowed_ips":          types.StringType,
		"address":              types.StringType,
		"port":                 types.StringType,
		"persistent_keepalive": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWireguardPeer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"public_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded public key

`,
		},

		"preshared_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded preshared key

`,
		},

		"allowed_ips": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP addresses allowed to traverse the peer

`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of tunnel endpoint

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of remote tunnel endpoint  |
|  ipv6  |  IPv6 address of remote tunnel endpoint  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		"persistent_keepalive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval to send keepalive messages

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Interval in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}
