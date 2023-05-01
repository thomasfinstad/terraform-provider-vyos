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

// InterfacesOpenvpnServerMfa describes the resource data model.
type InterfacesOpenvpnServerMfa struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeInterfacesOpenvpnServerMfaTotp types.Object `tfsdk:"totp"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesOpenvpnServerMfa) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "openvpn", "server", "mfa"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeInterfacesOpenvpnServerMfaTotp.IsNull() || o.NodeInterfacesOpenvpnServerMfaTotp.IsUnknown()) {
		var subModel InterfacesOpenvpnServerMfaTotp
		diags.Append(o.NodeInterfacesOpenvpnServerMfaTotp.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["totp"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesOpenvpnServerMfa) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "openvpn", "server", "mfa"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["totp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesOpenvpnServerMfaTotp{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesOpenvpnServerMfaTotp = data

	} else {
		o.NodeInterfacesOpenvpnServerMfaTotp = basetypes.NewObjectNull(InterfacesOpenvpnServerMfaTotp{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "openvpn", "server", "mfa"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesOpenvpnServerMfa) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"totp": types.ObjectType{AttrTypes: InterfacesOpenvpnServerMfaTotp{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerMfa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"totp": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServerMfaTotp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time-based one-time passwords

`,
		},
	}
}