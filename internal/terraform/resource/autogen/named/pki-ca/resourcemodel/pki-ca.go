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

// PkiCa describes the resource data model.
type PkiCa struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPkiCaCertificate types.String `tfsdk:"certificate"`
	LeafPkiCaDescrIPtion types.String `tfsdk:"description"`
	LeafPkiCaCrl         types.String `tfsdk:"crl"`
	LeafPkiCaRevoke      types.String `tfsdk:"revoke"`

	// TagNodes

	// Nodes
	NodePkiCaPrivate types.Object `tfsdk:"private"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCa) GetVyosPath() []string {
	return []string{
		"pki",
		"ca",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *PkiCa) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"pki", "ca"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPkiCaCertificate.IsNull() || o.LeafPkiCaCertificate.IsUnknown()) {
		vyosData["certificate"] = o.LeafPkiCaCertificate.ValueString()
	}
	if !(o.LeafPkiCaDescrIPtion.IsNull() || o.LeafPkiCaDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafPkiCaDescrIPtion.ValueString()
	}
	if !(o.LeafPkiCaCrl.IsNull() || o.LeafPkiCaCrl.IsUnknown()) {
		vyosData["crl"] = o.LeafPkiCaCrl.ValueString()
	}
	if !(o.LeafPkiCaRevoke.IsNull() || o.LeafPkiCaRevoke.IsUnknown()) {
		vyosData["revoke"] = o.LeafPkiCaRevoke.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePkiCaPrivate.IsNull() || o.NodePkiCaPrivate.IsUnknown()) {
		var subModel PkiCaPrivate
		diags.Append(o.NodePkiCaPrivate.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["private"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PkiCa) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"pki", "ca"}})

	// Leafs
	if value, ok := vyosData["certificate"]; ok {
		o.LeafPkiCaCertificate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCaCertificate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafPkiCaDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCaDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["crl"]; ok {
		o.LeafPkiCaCrl = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCaCrl = basetypes.NewStringNull()
	}
	if value, ok := vyosData["revoke"]; ok {
		o.LeafPkiCaRevoke = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiCaRevoke = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["private"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PkiCaPrivate{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePkiCaPrivate = data

	} else {
		o.NodePkiCaPrivate = basetypes.NewObjectNull(PkiCaPrivate{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"pki", "ca"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PkiCa) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"certificate": types.StringType,
		"description": types.StringType,
		"crl":         types.StringType,
		"revoke":      types.StringType,

		// Tags

		// Nodes
		"private": types.ObjectType{AttrTypes: PkiCaPrivate{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Certificate Authority

`,
		},

		// LeafNodes

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `CA certificate in PEM format

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

`,
		},

		"crl": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate revocation list in PEM format

`,
		},

		"revoke": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `If parent CA is present, this CA certificate will be included in generated CRLs

`,
		},

		// TagNodes

		// Nodes

		"private": schema.SingleNestedAttribute{
			Attributes: PkiCaPrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `CA private key in PEM format

`,
		},
	}
}