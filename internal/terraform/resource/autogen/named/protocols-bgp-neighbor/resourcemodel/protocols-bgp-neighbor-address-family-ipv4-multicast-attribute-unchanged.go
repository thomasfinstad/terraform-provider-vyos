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

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath  types.String `tfsdk:"as_path"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed     types.String `tfsdk:"med"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop types.String `tfsdk:"next_hop"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "attribute-unchanged"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.IsUnknown()) {
		vyosData["as-path"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.IsUnknown()) {
		vyosData["med"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.IsUnknown()) {
		vyosData["next-hop"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "attribute-unchanged"}})

	// Leafs
	if value, ok := vyosData["as-path"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}
	if value, ok := vyosData["med"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedMed = basetypes.NewStringNull()
	}
	if value, ok := vyosData["next-hop"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-multicast", "attribute-unchanged"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"as_path":  types.StringType,
		"med":      types.StringType,
		"next_hop": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send AS path unchanged

`,
		},

		"med": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
		},

		"next_hop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
		},

		// TagNodes

		// Nodes

	}
}