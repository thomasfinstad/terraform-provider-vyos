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

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn types.String `tfsdk:"vpn"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "import"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn.IsUnknown()) {
		vyosData["vpn"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf.IsNull() || o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "import"}})

	// Leafs
	if value, ok := vyosData["vpn"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVpn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastImportVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family", "ipv6-unicast", "import"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"vpn": types.StringType,
		"vrf": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastImport) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"vpn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `to/from default instance VPN RIB

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to import from

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		// Nodes

	}
}