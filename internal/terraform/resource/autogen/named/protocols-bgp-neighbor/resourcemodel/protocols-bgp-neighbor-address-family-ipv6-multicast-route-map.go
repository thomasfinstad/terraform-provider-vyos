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

// ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport types.String `tfsdk:"export"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "route-map"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport.IsUnknown()) {
		vyosData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport.IsUnknown()) {
		vyosData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "route-map"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMapImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "route-map"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixMulticastRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter outgoing route updates

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter incoming route updates

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}
