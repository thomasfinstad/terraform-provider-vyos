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

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce types.String `tfsdk:"force"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "nexthop-self"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce.IsUnknown()) {
		vyosData["force"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "nexthop-self"}})

	// Leafs
	if value, ok := vyosData["force"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelfForce = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-unicast", "nexthop-self"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"force": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"force": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set the next hop to self for reflected routes

`,
		},

		// TagNodes

		// Nodes

	}
}
