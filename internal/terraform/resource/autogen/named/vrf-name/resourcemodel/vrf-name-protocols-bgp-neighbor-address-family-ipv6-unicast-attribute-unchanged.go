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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed     types.String `tfsdk:"med"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "attribute-unchanged"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath.IsUnknown()) {
		vyosData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed.IsUnknown()) {
		vyosData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop.IsUnknown()) {
		vyosData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "attribute-unchanged"}})

	// Leafs
	if value, ok := vyosData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}
	if value, ok := vyosData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}
	if value, ok := vyosData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "attribute-unchanged"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged) AttributeTypes() map[string]attr.Type {
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
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
