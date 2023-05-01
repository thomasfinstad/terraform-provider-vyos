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

// VpnIPsecSiteToSitePeerVti describes the resource data model.
type VpnIPsecSiteToSitePeerVti struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerVtiBind     types.String `tfsdk:"bind"`
	LeafVpnIPsecSiteToSitePeerVtiEspGroup types.String `tfsdk:"esp_group"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecSiteToSitePeerVti) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "vti"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecSiteToSitePeerVtiBind.IsNull() || o.LeafVpnIPsecSiteToSitePeerVtiBind.IsUnknown()) {
		vyosData["bind"] = o.LeafVpnIPsecSiteToSitePeerVtiBind.ValueString()
	}
	if !(o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.IsNull() || o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.IsUnknown()) {
		vyosData["esp-group"] = o.LeafVpnIPsecSiteToSitePeerVtiEspGroup.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecSiteToSitePeerVti) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "vti"}})

	// Leafs
	if value, ok := vyosData["bind"]; ok {
		o.LeafVpnIPsecSiteToSitePeerVtiBind = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerVtiBind = basetypes.NewStringNull()
	}
	if value, ok := vyosData["esp-group"]; ok {
		o.LeafVpnIPsecSiteToSitePeerVtiEspGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerVtiEspGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "vti"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecSiteToSitePeerVti) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"bind":      types.StringType,
		"esp_group": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerVti) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bind": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VTI tunnel interface associated with this configuration

`,
		},

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		// TagNodes

		// Nodes

	}
}