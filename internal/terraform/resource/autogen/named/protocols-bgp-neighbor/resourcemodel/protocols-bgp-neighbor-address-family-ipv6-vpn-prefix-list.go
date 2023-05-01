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

// ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport types.String `tfsdk:"export"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-vpn", "prefix-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport.IsUnknown()) {
		vyosData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport.IsUnknown()) {
		vyosData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-vpn", "prefix-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-vpn", "prefix-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |

`,
		},

		// TagNodes

		// Nodes

	}
}