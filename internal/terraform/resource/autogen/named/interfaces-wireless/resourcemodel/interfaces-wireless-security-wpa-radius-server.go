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

// InterfacesWirelessSecURItyWpaRadiusServer describes the resource data model.
type InterfacesWirelessSecURItyWpaRadiusServer struct {
	// LeafNodes
	LeafInterfacesWirelessSecURItyWpaRadiusServerDisable    types.String `tfsdk:"disable"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerKey        types.String `tfsdk:"key"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerPort       types.String `tfsdk:"port"`
	LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting types.String `tfsdk:"accounting"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessSecURItyWpaRadiusServer) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "security", "wpa", "radius", "server"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessSecURItyWpaRadiusServerDisable.IsNull() || o.LeafInterfacesWirelessSecURItyWpaRadiusServerDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesWirelessSecURItyWpaRadiusServerDisable.ValueString()
	}
	if !(o.LeafInterfacesWirelessSecURItyWpaRadiusServerKey.IsNull() || o.LeafInterfacesWirelessSecURItyWpaRadiusServerKey.IsUnknown()) {
		vyosData["key"] = o.LeafInterfacesWirelessSecURItyWpaRadiusServerKey.ValueString()
	}
	if !(o.LeafInterfacesWirelessSecURItyWpaRadiusServerPort.IsNull() || o.LeafInterfacesWirelessSecURItyWpaRadiusServerPort.IsUnknown()) {
		vyosData["port"] = o.LeafInterfacesWirelessSecURItyWpaRadiusServerPort.ValueString()
	}
	if !(o.LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting.IsNull() || o.LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting.IsUnknown()) {
		vyosData["accounting"] = o.LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessSecURItyWpaRadiusServer) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "security", "wpa", "radius", "server"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["key"]; ok {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerPort = basetypes.NewStringNull()
	}
	if value, ok := vyosData["accounting"]; ok {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessSecURItyWpaRadiusServerAccounting = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "security", "wpa", "radius", "server"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessSecURItyWpaRadiusServer) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":    types.StringType,
		"key":        types.StringType,
		"port":       types.StringType,
		"accounting": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURItyWpaRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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

		"accounting": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable RADIUS server to receive accounting info

`,
		},

		// TagNodes

		// Nodes

	}
}
