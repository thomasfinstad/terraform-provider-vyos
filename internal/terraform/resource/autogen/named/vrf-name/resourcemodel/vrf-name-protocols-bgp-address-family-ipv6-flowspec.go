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

// VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall types.Object `tfsdk:"local_install"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-flowspec"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["local-install"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-flowspec"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["local-install"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-flowspec"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"local_install": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"local_install": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Apply local policy routing to interface

`,
		},
	}
}