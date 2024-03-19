// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisAreaPassword{}

// VrfNameProtocolsIsisAreaPassword describes the resource data model.
type VrfNameProtocolsIsisAreaPassword struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisAreaPasswordPlaintextPassword types.String `tfsdk:"plaintext_password" vyos:"plaintext-password,omitempty"`
	LeafVrfNameProtocolsIsisAreaPasswordMdfive            types.String `tfsdk:"md5" vyos:"md5,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisAreaPassword) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Plain-text authentication type

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Circuit password  |
`,
			Description: `Plain-text authentication type

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  Circuit password  |
`,
		},

		"md5": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Level-wide password  |
`,
			Description: `MD5 authentication type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Level-wide password  |
`,
		},

		// Nodes

	}
}
