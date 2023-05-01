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

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport types.String `tfsdk:"export"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-labeled-unicast", "distribute-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport.IsUnknown()) {
		vyosData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport.IsUnknown()) {
		vyosData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-labeled-unicast", "distribute-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-labeled-unicast", "distribute-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
