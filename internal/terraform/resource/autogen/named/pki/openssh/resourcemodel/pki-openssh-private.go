// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PkiOpenTCPPrivate{}

// PkiOpenTCPPrivate describes the resource data model.
type PkiOpenTCPPrivate struct {
	// LeafNodes
	LeafPkiOpenTCPPrivateKey               types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafPkiOpenTCPPrivatePasswordProtected types.Bool   `tfsdk:"password_protected" vyos:"password-protected,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiOpenTCPPrivate) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Private key in PEM format

`,
			Description: `Private key in PEM format

`,
		},

		"password_protected": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Private key portion is password protected

`,
			Description: `Private key portion is password protected

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
