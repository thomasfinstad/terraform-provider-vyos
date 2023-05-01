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

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended types.String `tfsdk:"extended"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard types.String `tfsdk:"standard"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "disable-send-community"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended.IsNull() || o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended.IsUnknown()) {
		vyosData["extended"] = o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard.IsNull() || o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard.IsUnknown()) {
		vyosData["standard"] = o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "disable-send-community"}})

	// Leafs
	if value, ok := vyosData["extended"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityExtended = basetypes.NewStringNull()
	}
	if value, ok := vyosData["standard"]; ok {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunityStandard = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "peer-group", "address-family", "ipv6-unicast", "disable-send-community"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"extended": types.StringType,
		"standard": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastDisableSendCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
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
