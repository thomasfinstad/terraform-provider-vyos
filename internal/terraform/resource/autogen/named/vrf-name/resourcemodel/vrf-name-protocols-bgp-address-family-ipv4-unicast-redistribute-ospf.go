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

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric   types.String `tfsdk:"metric"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "ospf"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric.IsUnknown()) {
		vyosData["metric"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "ospf"}})

	// Leafs
	if value, ok := vyosData["metric"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfMetric = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspfRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "redistribute", "ospf"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"metric":    types.StringType,
		"route_map": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf) ResourceSchemaAttributes() map[string]schema.Attribute {
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
