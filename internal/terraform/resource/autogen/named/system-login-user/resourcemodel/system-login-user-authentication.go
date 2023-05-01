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

// SystemLoginUserAuthentication describes the resource data model.
type SystemLoginUserAuthentication struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationEncryptedPassword types.String `tfsdk:"encrypted_password"`
	LeafSystemLoginUserAuthenticationPlaintextPassword types.String `tfsdk:"plaintext_password"`

	// TagNodes
	TagSystemLoginUserAuthenticationPublicKeys types.Map `tfsdk:"public_keys"`

	// Nodes
	NodeSystemLoginUserAuthenticationOtp types.Object `tfsdk:"otp"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *SystemLoginUserAuthentication) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafSystemLoginUserAuthenticationEncryptedPassword.IsNull() || o.LeafSystemLoginUserAuthenticationEncryptedPassword.IsUnknown()) {
		vyosData["encrypted-password"] = o.LeafSystemLoginUserAuthenticationEncryptedPassword.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationPlaintextPassword.IsNull() || o.LeafSystemLoginUserAuthenticationPlaintextPassword.IsUnknown()) {
		vyosData["plaintext-password"] = o.LeafSystemLoginUserAuthenticationPlaintextPassword.ValueString()
	}

	// Tags
	if !(o.TagSystemLoginUserAuthenticationPublicKeys.IsNull() || o.TagSystemLoginUserAuthenticationPublicKeys.IsUnknown()) {
		subModel := make(map[string]SystemLoginUserAuthenticationPublicKeys)
		diags.Append(o.TagSystemLoginUserAuthenticationPublicKeys.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["public-keys"] = subData
	}

	// Nodes
	if !(o.NodeSystemLoginUserAuthenticationOtp.IsNull() || o.NodeSystemLoginUserAuthenticationOtp.IsUnknown()) {
		var subModel SystemLoginUserAuthenticationOtp
		diags.Append(o.NodeSystemLoginUserAuthenticationOtp.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["otp"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *SystemLoginUserAuthentication) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication"}})

	// Leafs
	if value, ok := vyosData["encrypted-password"]; ok {
		o.LeafSystemLoginUserAuthenticationEncryptedPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationEncryptedPassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["plaintext-password"]; ok {
		o.LeafSystemLoginUserAuthenticationPlaintextPassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationPlaintextPassword = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["public-keys"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: SystemLoginUserAuthenticationPublicKeys{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagSystemLoginUserAuthenticationPublicKeys = data
	} else {
		o.TagSystemLoginUserAuthenticationPublicKeys = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["otp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, SystemLoginUserAuthenticationOtp{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeSystemLoginUserAuthenticationOtp = data

	} else {
		o.NodeSystemLoginUserAuthenticationOtp = basetypes.NewObjectNull(SystemLoginUserAuthenticationOtp{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o SystemLoginUserAuthentication) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"encrypted_password": types.StringType,
		"plaintext_password": types.StringType,

		// Tags
		"public_keys": types.MapType{ElemType: types.ObjectType{AttrTypes: SystemLoginUserAuthenticationPublicKeys{}.AttributeTypes()}},

		// Nodes
		"otp": types.ObjectType{AttrTypes: SystemLoginUserAuthenticationOtp{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encrypted password

`,

			// Default:          stringdefault.StaticString(`!`),
			Computed: true,
		},

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plaintext password used for encryption

`,
		},

		// TagNodes

		"public_keys": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: SystemLoginUserAuthenticationPublicKeys{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Remote access public keys

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Key identifier used by ssh-keygen (usually of form user@host)  |

`,
		},

		// Nodes

		"otp": schema.SingleNestedAttribute{
			Attributes: SystemLoginUserAuthenticationOtp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `One-Time-Pad (two-factor) authentication parameters

`,
		},
	}
}