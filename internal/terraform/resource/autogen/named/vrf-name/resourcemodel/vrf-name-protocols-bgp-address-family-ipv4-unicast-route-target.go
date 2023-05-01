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

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn types.Object `tfsdk:"vpn"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "route-target"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["vpn"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "route-target"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["vpn"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-unicast", "route-target"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"vpn": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
		},
	}
}
