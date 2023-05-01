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

// VpnOpenconnectAccountingRadiusServer describes the resource data model.
type VpnOpenconnectAccountingRadiusServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVpnOpenconnectAccountingRadiusServerDisable types.String `tfsdk:"disable"`
	LeafVpnOpenconnectAccountingRadiusServerKey     types.String `tfsdk:"key"`
	LeafVpnOpenconnectAccountingRadiusServerPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectAccountingRadiusServer) GetVyosPath() []string {
	return []string{
		"vpn",
		"openconnect",
		"accounting",
		"radius",
		"server",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnOpenconnectAccountingRadiusServer) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "openconnect", "accounting", "radius", "server"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnOpenconnectAccountingRadiusServerDisable.IsNull() || o.LeafVpnOpenconnectAccountingRadiusServerDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafVpnOpenconnectAccountingRadiusServerDisable.ValueString()
	}
	if !(o.LeafVpnOpenconnectAccountingRadiusServerKey.IsNull() || o.LeafVpnOpenconnectAccountingRadiusServerKey.IsUnknown()) {
		vyosData["key"] = o.LeafVpnOpenconnectAccountingRadiusServerKey.ValueString()
	}
	if !(o.LeafVpnOpenconnectAccountingRadiusServerPort.IsNull() || o.LeafVpnOpenconnectAccountingRadiusServerPort.IsUnknown()) {
		vyosData["port"] = o.LeafVpnOpenconnectAccountingRadiusServerPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnOpenconnectAccountingRadiusServer) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "openconnect", "accounting", "radius", "server"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafVpnOpenconnectAccountingRadiusServerDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnOpenconnectAccountingRadiusServerDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["key"]; ok {
		o.LeafVpnOpenconnectAccountingRadiusServerKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnOpenconnectAccountingRadiusServerKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafVpnOpenconnectAccountingRadiusServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnOpenconnectAccountingRadiusServerPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "openconnect", "accounting", "radius", "server"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnOpenconnectAccountingRadiusServer) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable": types.StringType,
		"key":     types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectAccountingRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  RADIUS server IPv4 address  |

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}