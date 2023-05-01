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

// PkiCertificatePrivate describes the resource data model.
type PkiCertificatePrivate struct {
	// LeafNodes
	LeafPkiCertificatePrivateKey               types.String `tfsdk:"key"`
	LeafPkiCertificatePrivatePasswordProtected types.String `tfsdk:"password_protected"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PkiCertificatePrivate) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"pki", "certificate", "private"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPkiCertificatePrivateKey.IsNull() || o.LeafPkiCertificatePrivateKey.IsUnknown()) {
		vyosData["key"] = o.LeafPkiCertificatePrivateKey.ValueString()
	}
	if !(o.LeafPkiCertificatePrivatePasswordProtected.IsNull() || o.LeafPkiCertificatePrivatePasswordProtected.IsUnknown()) {
		vyosData["password-protected"] = o.LeafPkiCertificatePrivatePasswordProtected.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PkiCertificatePrivate) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"pki", "certificate", "private"}})

	// Leafs
	if value, ok := vyosData["key"]; ok {
		o.LeafPkiCertificatePrivateKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCertificatePrivateKey = basetypes.NewStringNull()
	}
	if value, ok := vyosData["password-protected"]; ok {
		o.LeafPkiCertificatePrivatePasswordProtected = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCertificatePrivatePasswordProtected = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"pki", "certificate", "private"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PkiCertificatePrivate) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"key":                types.StringType,
		"password_protected": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCertificatePrivate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate private key in PEM format

`,
		},

		"password_protected": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate private key is password protected

`,
		},

		// TagNodes

		// Nodes

	}
}