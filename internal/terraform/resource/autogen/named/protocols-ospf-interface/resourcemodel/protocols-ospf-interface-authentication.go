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

// ProtocolsOspfInterfaceAuthentication describes the resource data model.
type ProtocolsOspfInterfaceAuthentication struct {
	// LeafNodes
	LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password"`

	// TagNodes

	// Nodes
	NodeProtocolsOspfInterfaceAuthenticationMdfive types.Object `tfsdk:"md5"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsOspfInterfaceAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword.IsNull() || o.LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword.IsUnknown()) {
		vyosData["plaintext-password"] = o.LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsOspfInterfaceAuthenticationMdfive.IsNull() || o.NodeProtocolsOspfInterfaceAuthenticationMdfive.IsUnknown()) {
		var subModel ProtocolsOspfInterfaceAuthenticationMdfive
		diags.Append(o.NodeProtocolsOspfInterfaceAuthenticationMdfive.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["md5"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsOspfInterfaceAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface", "authentication"}})

	// Leafs
	if value, ok := vyosData["plaintext-password"]; ok {
		o.LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfInterfaceAuthenticationPlaintextPassword = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["md5"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsOspfInterfaceAuthenticationMdfive{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsOspfInterfaceAuthenticationMdfive = data

	} else {
		o.NodeProtocolsOspfInterfaceAuthenticationMdfive = basetypes.NewObjectNull(ProtocolsOspfInterfaceAuthenticationMdfive{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ospf", "interface", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsOspfInterfaceAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"plaintext_password": types.StringType,

		// Tags

		// Nodes
		"md5": types.ObjectType{AttrTypes: ProtocolsOspfInterfaceAuthenticationMdfive{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfInterfaceAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain text password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Plain text password (8 characters or less)  |

`,
		},

		// TagNodes

		// Nodes

		"md5": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfInterfaceAuthenticationMdfive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `MD5 key id

`,
		},
	}
}