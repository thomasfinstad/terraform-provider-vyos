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

// VrfNameProtocolsOspfRedistributeStatic describes the resource data model.
type VrfNameProtocolsOspfRedistributeStatic struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfRedistributeStaticMetric     types.String `tfsdk:"metric"`
	LeafVrfNameProtocolsOspfRedistributeStaticMetricType types.String `tfsdk:"metric_type"`
	LeafVrfNameProtocolsOspfRedistributeStaticRouteMap   types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfRedistributeStatic) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "redistribute", "static"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfRedistributeStaticMetric.IsNull() || o.LeafVrfNameProtocolsOspfRedistributeStaticMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafVrfNameProtocolsOspfRedistributeStaticMetric.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfRedistributeStaticMetricType.IsNull() || o.LeafVrfNameProtocolsOspfRedistributeStaticMetricType.IsUnknown()) {
		vyosData["metric-type"] = o.LeafVrfNameProtocolsOspfRedistributeStaticMetricType.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfRedistributeStaticRouteMap.IsNull() || o.LeafVrfNameProtocolsOspfRedistributeStaticRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsOspfRedistributeStaticRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfRedistributeStatic) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "redistribute", "static"}})

	// Leafs
	if value, ok := vyosData["metric"]; ok {
		o.LeafVrfNameProtocolsOspfRedistributeStaticMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfRedistributeStaticMetric = basetypes.NewStringNull()
	}
	if value, ok := vyosData["metric-type"]; ok {
		o.LeafVrfNameProtocolsOspfRedistributeStaticMetricType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfRedistributeStaticMetricType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsOspfRedistributeStaticRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfRedistributeStaticRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "redistribute", "static"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfRedistributeStatic) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"metric":      types.StringType,
		"metric_type": types.StringType,
		"route_map":   types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRedistributeStatic) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF default metric

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777214  |  Default metric  |

`,
		},

		"metric_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF metric type for default routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2  |  Set OSPF External Type 1/2 metrics  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}