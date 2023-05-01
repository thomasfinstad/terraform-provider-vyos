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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended types.String `tfsdk:"extended"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard types.String `tfsdk:"standard"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "disable-send-community"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended.IsUnknown()) {
		vyosData["extended"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard.IsUnknown()) {
		vyosData["standard"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "disable-send-community"}})

	// Leafs
	if value, ok := vyosData["extended"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityExtended = basetypes.NewStringNull()
	}
	if value, ok := vyosData["standard"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunityStandard = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-unicast", "disable-send-community"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"extended": types.StringType,
		"standard": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixUnicastDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"extended": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending extended community attributes to this peer

`,
		},

		"standard": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending standard community attributes to this peer

`,
		},

		// TagNodes

		// Nodes

	}
}
