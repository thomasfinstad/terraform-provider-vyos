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

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound types.String `tfsdk:"inbound"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "soft-reconfiguration"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound.IsNull() || o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound.IsUnknown()) {
		vyosData["inbound"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "soft-reconfiguration"}})

	// Leafs
	if value, ok := vyosData["inbound"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfigurationInbound = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "soft-reconfiguration"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"inbound": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
		},

		// TagNodes

		// Nodes

	}
}
