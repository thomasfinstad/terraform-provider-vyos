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

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn types.Object `tfsdk:"vpn"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "label"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["vpn"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "label"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["vpn"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "label"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"vpn": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabel) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
		},
	}
}