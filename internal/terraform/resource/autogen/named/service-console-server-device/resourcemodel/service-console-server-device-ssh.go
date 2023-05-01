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

// ServiceConsoleServerDeviceTCP describes the resource data model.
type ServiceConsoleServerDeviceTCP struct {
	// LeafNodes
	LeafServiceConsoleServerDeviceTCPPort types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceConsoleServerDeviceTCP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "console-server", "device", "ssh"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceConsoleServerDeviceTCPPort.IsNull() || o.LeafServiceConsoleServerDeviceTCPPort.IsUnknown()) {
		vyosData["port"] = o.LeafServiceConsoleServerDeviceTCPPort.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceConsoleServerDeviceTCP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "console-server", "device", "ssh"}})

	// Leafs
	if value, ok := vyosData["port"]; ok {
		o.LeafServiceConsoleServerDeviceTCPPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceConsoleServerDeviceTCPPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "console-server", "device", "ssh"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceConsoleServerDeviceTCP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"port": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConsoleServerDeviceTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		// TagNodes

		// Nodes

	}
}