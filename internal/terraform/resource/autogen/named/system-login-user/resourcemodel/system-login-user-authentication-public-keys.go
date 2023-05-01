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

// SystemLoginUserAuthenticationPublicKeys describes the resource data model.
type SystemLoginUserAuthenticationPublicKeys struct {
	// LeafNodes
	LeafSystemLoginUserAuthenticationPublicKeysKey     types.String `tfsdk:"key"`
	LeafSystemLoginUserAuthenticationPublicKeysOptions types.String `tfsdk:"options"`
	LeafSystemLoginUserAuthenticationPublicKeysType    types.String `tfsdk:"type"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *SystemLoginUserAuthenticationPublicKeys) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "public-keys"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafSystemLoginUserAuthenticationPublicKeysKey.IsNull() || o.LeafSystemLoginUserAuthenticationPublicKeysKey.IsUnknown()) {
		vyosData["key"] = o.LeafSystemLoginUserAuthenticationPublicKeysKey.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationPublicKeysOptions.IsNull() || o.LeafSystemLoginUserAuthenticationPublicKeysOptions.IsUnknown()) {
		vyosData["options"] = o.LeafSystemLoginUserAuthenticationPublicKeysOptions.ValueString()
	}
	if !(o.LeafSystemLoginUserAuthenticationPublicKeysType.IsNull() || o.LeafSystemLoginUserAuthenticationPublicKeysType.IsUnknown()) {
		vyosData["type"] = o.LeafSystemLoginUserAuthenticationPublicKeysType.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *SystemLoginUserAuthenticationPublicKeys) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "public-keys"}})

	// Leafs
	if value, ok := vyosData["key"]; ok {
		o.LeafSystemLoginUserAuthenticationPublicKeysKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationPublicKeysKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["options"]; ok {
		o.LeafSystemLoginUserAuthenticationPublicKeysOptions = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationPublicKeysOptions = basetypes.NewStringNull()
	}
	if value, ok := vyosData["type"]; ok {
		o.LeafSystemLoginUserAuthenticationPublicKeysType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemLoginUserAuthenticationPublicKeysType = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"system", "login", "user", "authentication", "public-keys"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o SystemLoginUserAuthenticationPublicKeys) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"key":     types.StringType,
		"options": types.StringType,
		"type":    types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUserAuthenticationPublicKeys) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Public key value (Base64 encoded)

`,
		},

		"options": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Optional public key options

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SSH public key type

|  Format  |  Description  |
|----------|---------------|
|  ssh-dss  |  Digital Signature Algorithm (DSA) key support  |
|  ssh-rsa  |  Key pair based on RSA algorithm  |
|  ecdsa-sha2-nistp256  |  Elliptic Curve DSA with NIST P-256 curve  |
|  ecdsa-sha2-nistp384  |  Elliptic Curve DSA with NIST P-384 curve  |
|  ecdsa-sha2-nistp521  |  Elliptic Curve DSA with NIST P-521 curve  |
|  ssh-ed25519  |  Edwards-curve DSA with elliptic curve 25519  |
|  sk-ecdsa-sha2-nistp256@openssh.com  |  Elliptic Curve DSA security key  |
|  sk-ssh-ed25519@openssh.com  |  Elliptic curve 25519 security key  |

`,
		},

		// TagNodes

		// Nodes

	}
}