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

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf types.Object `tfsdk:"orf"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "capability"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf.IsNull() || o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf.IsUnknown()) {
		var subModel ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf
		diags.Append(o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["orf"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "capability"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["orf"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf = data

	} else {
		o.NodeProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf = basetypes.NewObjectNull(ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "peer-group", "address-family", "ipv4-unicast", "capability"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"orf": types.ObjectType{AttrTypes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"orf": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise ORF capability to this peer

`,
		},
	}
}
