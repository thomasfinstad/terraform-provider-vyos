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

// VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface types.String `tfsdk:"interface"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-flowspec", "local-install"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-flowspec", "local-install"}})

	// Leafs
	if value, ok := vyosData["interface"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv4-flowspec", "local-install"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"interface": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface

`,
		},

		// TagNodes

		// Nodes

	}
}
