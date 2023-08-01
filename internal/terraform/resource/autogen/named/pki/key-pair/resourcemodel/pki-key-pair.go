// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PkiKeyPair describes the resource data model.
type PkiKeyPair struct {
	SelfIdentifier types.String `tfsdk:"key_pair_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePkiKeyPairPublic  *PkiKeyPairPublic  `tfsdk:"public" vyos:"public,omitempty"`
	NodePkiKeyPairPrivate *PkiKeyPairPrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiKeyPair) GetVyosPath() []string {
	return []string{
		"pki",

		"key-pair",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPair) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `key_pair_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"key_pair_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Public and private keys

`,
		},

		// LeafNodes

		// Nodes

		"public": schema.SingleNestedAttribute{
			Attributes: PkiKeyPairPublic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Public key

`,
		},

		"private": schema.SingleNestedAttribute{
			Attributes: PkiKeyPairPrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Private key

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiKeyPair) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PkiKeyPair) UnmarshalJSON(_ []byte) error {
	return nil
}
