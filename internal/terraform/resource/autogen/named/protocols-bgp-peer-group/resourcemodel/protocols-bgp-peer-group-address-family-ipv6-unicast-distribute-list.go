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

// ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport types.String `tfsdk:"export"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport types.String `tfsdk:"import"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "distribute-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport.IsNull() || o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport.IsUnknown()) {
		vyosData["export"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport.ValueString()
	}
	if !(o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport.IsNull() || o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport.IsUnknown()) {
		vyosData["import"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "distribute-list"}})

	// Leafs
	if value, ok := vyosData["export"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListExport = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "distribute-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export": types.StringType,
		"import": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
