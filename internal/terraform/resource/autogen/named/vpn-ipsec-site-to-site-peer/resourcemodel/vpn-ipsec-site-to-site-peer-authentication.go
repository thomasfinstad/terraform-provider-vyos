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

// VpnIPsecSiteToSitePeerAuthentication describes the resource data model.
type VpnIPsecSiteToSitePeerAuthentication struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerAuthenticationLocalID            types.String `tfsdk:"local_id"`
	LeafVpnIPsecSiteToSitePeerAuthenticationMode               types.String `tfsdk:"mode"`
	LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID           types.String `tfsdk:"remote_id"`
	LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID types.String `tfsdk:"use_x509_id"`

	// TagNodes

	// Nodes
	NodeVpnIPsecSiteToSitePeerAuthenticationRsa           types.Object `tfsdk:"rsa"`
	NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine types.Object `tfsdk:"x509"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecSiteToSitePeerAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecSiteToSitePeerAuthenticationLocalID.IsNull() || o.LeafVpnIPsecSiteToSitePeerAuthenticationLocalID.IsUnknown()) {
		vyosData["local-id"] = o.LeafVpnIPsecSiteToSitePeerAuthenticationLocalID.ValueString()
	}
	if !(o.LeafVpnIPsecSiteToSitePeerAuthenticationMode.IsNull() || o.LeafVpnIPsecSiteToSitePeerAuthenticationMode.IsUnknown()) {
		vyosData["mode"] = o.LeafVpnIPsecSiteToSitePeerAuthenticationMode.ValueString()
	}
	if !(o.LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID.IsNull() || o.LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID.IsUnknown()) {
		vyosData["remote-id"] = o.LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID.ValueString()
	}
	if !(o.LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID.IsNull() || o.LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID.IsUnknown()) {
		vyosData["use-x509-id"] = o.LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVpnIPsecSiteToSitePeerAuthenticationRsa.IsNull() || o.NodeVpnIPsecSiteToSitePeerAuthenticationRsa.IsUnknown()) {
		var subModel VpnIPsecSiteToSitePeerAuthenticationRsa
		diags.Append(o.NodeVpnIPsecSiteToSitePeerAuthenticationRsa.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["rsa"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine.IsNull() || o.NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine.IsUnknown()) {
		var subModel VpnIPsecSiteToSitePeerAuthenticationXfivezeronine
		diags.Append(o.NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["x509"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecSiteToSitePeerAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "authentication"}})

	// Leafs
	if value, ok := vyosData["local-id"]; ok {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationLocalID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationLocalID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mode"]; ok {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationMode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remote-id"]; ok {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationRemoteID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["use-x509-id"]; ok {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecSiteToSitePeerAuthenticationUseXfivezeronineID = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["rsa"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecSiteToSitePeerAuthenticationRsa{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecSiteToSitePeerAuthenticationRsa = data

	} else {
		o.NodeVpnIPsecSiteToSitePeerAuthenticationRsa = basetypes.NewObjectNull(VpnIPsecSiteToSitePeerAuthenticationRsa{}.AttributeTypes())
	}
	if value, ok := vyosData["x509"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecSiteToSitePeerAuthenticationXfivezeronine{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine = data

	} else {
		o.NodeVpnIPsecSiteToSitePeerAuthenticationXfivezeronine = basetypes.NewObjectNull(VpnIPsecSiteToSitePeerAuthenticationXfivezeronine{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "site-to-site", "peer", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecSiteToSitePeerAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"local_id":    types.StringType,
		"mode":        types.StringType,
		"remote_id":   types.StringType,
		"use_x509_id": types.StringType,

		// Tags

		// Nodes
		"rsa":  types.ObjectType{AttrTypes: VpnIPsecSiteToSitePeerAuthenticationRsa{}.AttributeTypes()},
		"x509": types.ObjectType{AttrTypes: VpnIPsecSiteToSitePeerAuthenticationXfivezeronine{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local ID for peer authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Local ID used for peer authentication  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use pre-shared secret key  |
|  rsa  |  Use RSA key  |
|  x509  |  Use x.509 certificate  |

`,
		},

		"remote_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ID for remote authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  ID used for peer authentication  |

`,

			// Default:          stringdefault.StaticString(`%any`),
			Computed: true,
		},

		"use_x509_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use certificate common name as ID

`,
		},

		// TagNodes

		// Nodes

		"rsa": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerAuthenticationRsa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `RSA keys

`,
		},

		"x509": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerAuthenticationXfivezeronine{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `X.509 certificate

`,
		},
	}
}
