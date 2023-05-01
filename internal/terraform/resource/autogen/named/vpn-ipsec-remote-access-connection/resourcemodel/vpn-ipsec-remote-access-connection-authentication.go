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

// VpnIPsecRemoteAccessConnectionAuthentication describes the resource data model.
type VpnIPsecRemoteAccessConnectionAuthentication struct {
	// LeafNodes
	LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID         types.String `tfsdk:"local_id"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode      types.String `tfsdk:"client_mode"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode      types.String `tfsdk:"server_mode"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret"`

	// TagNodes

	// Nodes
	NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine types.Object `tfsdk:"x509"`
	NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers    types.Object `tfsdk:"local_users"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnIPsecRemoteAccessConnectionAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "ipsec", "remote-access", "connection", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.IsNull() || o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.IsUnknown()) {
		vyosData["local-id"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID.ValueString()
	}
	if !(o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.IsNull() || o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.IsUnknown()) {
		vyosData["client-mode"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode.ValueString()
	}
	if !(o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.IsNull() || o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.IsUnknown()) {
		vyosData["server-mode"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode.ValueString()
	}
	if !(o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.IsNull() || o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.IsUnknown()) {
		vyosData["pre-shared-secret"] = o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine.IsNull() || o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine.IsUnknown()) {
		var subModel VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine
		diags.Append(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["x509"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers.IsNull() || o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers.IsUnknown()) {
		var subModel VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers
		diags.Append(o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["local-users"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnIPsecRemoteAccessConnectionAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "ipsec", "remote-access", "connection", "authentication"}})

	// Leafs
	if value, ok := vyosData["local-id"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["client-mode"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["server-mode"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["pre-shared-secret"]; ok {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["x509"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine = data

	} else {
		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine = basetypes.NewObjectNull(VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.AttributeTypes())
	}
	if value, ok := vyosData["local-users"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers = data

	} else {
		o.NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers = basetypes.NewObjectNull(VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "ipsec", "remote-access", "connection", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnIPsecRemoteAccessConnectionAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"local_id":          types.StringType,
		"client_mode":       types.StringType,
		"server_mode":       types.StringType,
		"pre_shared_secret": types.StringType,

		// Tags

		// Nodes
		"x509":        types.ObjectType{AttrTypes: VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.AttributeTypes()},
		"local_users": types.ObjectType{AttrTypes: VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessConnectionAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"client_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client authentication mode

|  Format  |  Description  |
|----------|---------------|
|  eap-tls  |  Use EAP-TLS authentication  |
|  eap-mschapv2  |  Use EAP-MSCHAPv2 authentication  |
|  eap-radius  |  Use EAP-RADIUS authentication  |

`,

			// Default:          stringdefault.StaticString(`eap-mschapv2`),
			Computed: true,
		},

		"server_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use a pre-shared secret key  |
|  x509  |  Use x.509 certificate  |

`,

			// Default:          stringdefault.StaticString(`x509`),
			Computed: true,
		},

		"pre_shared_secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Pre-shared secret key  |

`,
		},

		// TagNodes

		// Nodes

		"x509": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `X.509 certificate

`,
		},

		"local_users": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Local user authentication

`,
		},
	}
}