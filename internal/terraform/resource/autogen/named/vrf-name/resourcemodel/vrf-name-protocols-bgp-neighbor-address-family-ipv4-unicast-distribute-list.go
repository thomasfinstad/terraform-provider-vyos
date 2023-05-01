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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport types.String `tfsdk:"export"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "distribute-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport.IsUnknown()) {
		vyosData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport.IsUnknown()) {
		vyosData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "distribute-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "distribute-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter outgoing route updates to this peer-group  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter incoming route updates from this peer-group  |

`,
		},

		// TagNodes

		// Nodes

	}
}
