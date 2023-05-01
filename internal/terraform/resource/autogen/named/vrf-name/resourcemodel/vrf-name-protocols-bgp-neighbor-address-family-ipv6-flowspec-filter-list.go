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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport types.String `tfsdk:"export"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-flowspec", "filter-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport.IsUnknown()) {
		vyosData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport.IsUnknown()) {
		vyosData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-flowspec", "filter-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-flowspec", "filter-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecFilterList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
		},

		// TagNodes

		// Nodes

	}
}
