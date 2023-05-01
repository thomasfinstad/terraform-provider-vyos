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

// ServiceRouterAdvertInterfaceInterval describes the resource data model.
type ServiceRouterAdvertInterfaceInterval struct {
	// LeafNodes
	LeafServiceRouterAdvertInterfaceIntervalMax types.String `tfsdk:"max"`
	LeafServiceRouterAdvertInterfaceIntervalMin types.String `tfsdk:"min"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ServiceRouterAdvertInterfaceInterval) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"service", "router-advert", "interface", "interval"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafServiceRouterAdvertInterfaceIntervalMax.IsNull() || o.LeafServiceRouterAdvertInterfaceIntervalMax.IsUnknown()) {
		vyosData["max"] = o.LeafServiceRouterAdvertInterfaceIntervalMax.ValueString()
	}
	if !(o.LeafServiceRouterAdvertInterfaceIntervalMin.IsNull() || o.LeafServiceRouterAdvertInterfaceIntervalMin.IsUnknown()) {
		vyosData["min"] = o.LeafServiceRouterAdvertInterfaceIntervalMin.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ServiceRouterAdvertInterfaceInterval) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"service", "router-advert", "interface", "interval"}})

	// Leafs
	if value, ok := vyosData["max"]; ok {
		o.LeafServiceRouterAdvertInterfaceIntervalMax = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceIntervalMax = basetypes.NewStringNull()
	}
	if value, ok := vyosData["min"]; ok {
		o.LeafServiceRouterAdvertInterfaceIntervalMin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceRouterAdvertInterfaceIntervalMin = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"service", "router-advert", "interface", "interval"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ServiceRouterAdvertInterfaceInterval) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"max": types.StringType,
		"min": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceRouterAdvertInterfaceInterval) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"max": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum interval between unsolicited multicast RAs

|  Format  |  Description  |
|----------|---------------|
|  u32:4-1800  |  Maximum interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`600`),
			Computed: true,
		},

		"min": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between unsolicited multicast RAs

|  Format  |  Description  |
|----------|---------------|
|  u32:3-1350  |  Minimum interval in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}