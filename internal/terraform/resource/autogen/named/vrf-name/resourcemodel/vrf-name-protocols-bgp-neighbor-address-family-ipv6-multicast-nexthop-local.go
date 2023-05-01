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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged types.String `tfsdk:"unchanged"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "nexthop-local"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged.IsUnknown()) {
		vyosData["unchanged"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "nexthop-local"}})

	// Leafs
	if value, ok := vyosData["unchanged"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocalUnchanged = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-multicast", "nexthop-local"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"unchanged": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixMulticastNexthopLocal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unchanged": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
		},

		// TagNodes

		// Nodes

	}
}