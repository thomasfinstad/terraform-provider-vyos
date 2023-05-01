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

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric   types.String `tfsdk:"metric"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "babel"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "babel"}})

	// Leafs
	if value, ok := vyosData["metric"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelMetric = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabelRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "babel"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"metric":    types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Metric for redistributed routes

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Metric for redistributed routes  |

`,
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
