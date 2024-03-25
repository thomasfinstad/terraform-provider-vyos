// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PkiKeyPairPublic{}

// PkiKeyPairPublic describes the resource data model.
type PkiKeyPairPublic struct {
	// LeafNodes
	LeafPkiKeyPairPublicKey types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPairPublic) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Public key in PEM format

`,
			Description: `Public key in PEM format

`,
		},

		// Nodes

	}
}
