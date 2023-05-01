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

// ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath  types.String `tfsdk:"as_path"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed     types.String `tfsdk:"med"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop types.String `tfsdk:"next_hop"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "attribute-unchanged"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath.IsUnknown()) {
		vyosData["as-path"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed.IsUnknown()) {
		vyosData["med"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop.IsUnknown()) {
		vyosData["next-hop"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "attribute-unchanged"}})

	// Leafs
	if value, ok := vyosData["as-path"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedAsPath = basetypes.NewStringNull()
	}
	if value, ok := vyosData["med"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedMed = basetypes.NewStringNull()
	}
	if value, ok := vyosData["next-hop"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn", "attribute-unchanged"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged) AttributeTypes() map[string]attr.Type {
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
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
