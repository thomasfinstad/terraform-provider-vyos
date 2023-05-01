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

// ServicePppoeServerAuthenticationRadiusServer describes the resource data model.
type ServicePppoeServerAuthenticationRadiusServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafServicePppoeServerAuthenticationRadiusServerDisable           types.String `tfsdk:"disable"`
	LeafServicePppoeServerAuthenticationRadiusServerKey               types.String `tfsdk:"key"`
	LeafServicePppoeServerAuthenticationRadiusServerPort              types.String `tfsdk:"port"`
	LeafServicePppoeServerAuthenticationRadiusServerAcctPort          types.String `tfsdk:"acct_port"`
	LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting types.String `tfsdk:"disable_accounting"`
	LeafServicePppoeServerAuthenticationRadiusServerFailTime          types.String `tfsdk:"fail_time"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"service",
		"pppoe-server",
		"authentication",
		"radius",
		"server",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServicePppoeServerAuthenticationRadiusServer) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "radius", "server"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerDisable.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafServicePppoeServerAuthenticationRadiusServerDisable.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerKey.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerKey.IsUnknown()) {
		vyosData["key"] = o.LeafServicePppoeServerAuthenticationRadiusServerKey.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerPort.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerPort.IsUnknown()) {
		vyosData["port"] = o.LeafServicePppoeServerAuthenticationRadiusServerPort.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerAcctPort.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerAcctPort.IsUnknown()) {
		vyosData["acct-port"] = o.LeafServicePppoeServerAuthenticationRadiusServerAcctPort.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting.IsUnknown()) {
		vyosData["disable-accounting"] = o.LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting.ValueString()
	}
	if !(o.LeafServicePppoeServerAuthenticationRadiusServerFailTime.IsNull() || o.LeafServicePppoeServerAuthenticationRadiusServerFailTime.IsUnknown()) {
		vyosData["fail-time"] = o.LeafServicePppoeServerAuthenticationRadiusServerFailTime.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServicePppoeServerAuthenticationRadiusServer) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "radius", "server"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["key"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["acct-port"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerAcctPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerAcctPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-accounting"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerDisableAccounting = basetypes.NewStringNull()
	}
	if value, ok := vyosData["fail-time"]; ok {
		o.LeafServicePppoeServerAuthenticationRadiusServerFailTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServicePppoeServerAuthenticationRadiusServerFailTime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "pppoe-server", "authentication", "radius", "server"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServicePppoeServerAuthenticationRadiusServer) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":            types.StringType,
		"key":                types.StringType,
		"port":               types.StringType,
		"acct_port":          types.StringType,
		"disable_accounting": types.StringType,
		"fail_time":          types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			MarkdownDescription: `Authentication port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		"disable_accounting": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable accounting

`,
		},

		"fail_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mark server unavailable for <n> seconds on failure

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600  |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}