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

// PkiKeyPairPrivate describes the resource data model.
type PkiKeyPairPrivate struct {
	// LeafNodes
	LeafPkiKeyPairPrivateKey               types.String `tfsdk:"key"`
	LeafPkiKeyPairPrivatePasswordProtected types.String `tfsdk:"password_protected"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PkiKeyPairPrivate) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"pki", "key-pair", "private"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPkiKeyPairPrivateKey.IsNull() || o.LeafPkiKeyPairPrivateKey.IsUnknown()) {
		vyosData["key"] = o.LeafPkiKeyPairPrivateKey.ValueString()
	}
	if !(o.LeafPkiKeyPairPrivatePasswordProtected.IsNull() || o.LeafPkiKeyPairPrivatePasswordProtected.IsUnknown()) {
		vyosData["password-protected"] = o.LeafPkiKeyPairPrivatePasswordProtected.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PkiKeyPairPrivate) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"pki", "key-pair", "private"}})

	// Leafs
	if value, ok := vyosData["key"]; ok {
		o.LeafPkiKeyPairPrivateKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiKeyPairPrivateKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["password-protected"]; ok {
		o.LeafPkiKeyPairPrivatePasswordProtected = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiKeyPairPrivatePasswordProtected = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"pki", "key-pair", "private"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PkiKeyPairPrivate) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"key":                types.StringType,
		"password_protected": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPairPrivate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Private key in PEM format

`,
		},

		"password_protected": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Private key is password protected

`,
		},

		// TagNodes

		// Nodes

	}
}