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

// ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce types.String `tfsdk:"force"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "nexthop-self"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce.IsUnknown()) {
		vyosData["force"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "nexthop-self"}})

	// Leafs
	if value, ok := vyosData["force"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelfForce = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "nexthop-self"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"force": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf) ResourceSchemaAttributes() map[string]schema.Attribute {
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
