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

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber types.String `tfsdk:"number"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "allowas-in"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.IsUnknown()) {
		vyosData["number"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "allowas-in"}})

	// Leafs
	if value, ok := vyosData["number"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasInNumber = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "allowas-in"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"number": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of times AS is allowed in path  |

`,
		},

		// TagNodes

		// Nodes

	}
}
