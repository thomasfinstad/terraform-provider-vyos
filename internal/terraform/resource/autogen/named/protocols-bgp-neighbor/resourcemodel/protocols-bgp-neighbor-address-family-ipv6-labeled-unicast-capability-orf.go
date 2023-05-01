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

// ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList types.Object `tfsdk:"prefix_list"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "capability", "orf"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "capability", "orf"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv6-labeled-unicast", "capability", "orf"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"prefix_list": types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}
