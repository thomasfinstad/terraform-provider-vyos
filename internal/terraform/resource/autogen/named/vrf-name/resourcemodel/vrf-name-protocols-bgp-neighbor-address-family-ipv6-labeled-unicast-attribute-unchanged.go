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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath  types.String `tfsdk:"as_path"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed     types.String `tfsdk:"med"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop types.String `tfsdk:"next_hop"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "attribute-unchanged"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.IsUnknown()) {
		vyosData["as-path"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.IsUnknown()) {
		vyosData["med"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.IsUnknown()) {
		vyosData["next-hop"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "attribute-unchanged"}})

	// Leafs
	if value, ok := vyosData["as-path"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedAsPath = basetypes.NewStringNull()
	}
	if value, ok := vyosData["med"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedMed = basetypes.NewStringNull()
	}
	if value, ok := vyosData["next-hop"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchangedNextHop = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "attribute-unchanged"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) AttributeTypes() map[string]attr.Type {
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
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastAttributeUnchanged) ResourceSchemaAttributes() map[string]schema.Attribute {
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
