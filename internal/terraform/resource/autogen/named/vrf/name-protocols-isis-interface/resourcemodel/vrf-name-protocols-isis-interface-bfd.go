// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisInterfaceBfd{}

// VrfNameProtocolsIsisInterfaceBfd describes the resource data model.
type VrfNameProtocolsIsisInterfaceBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceBfdProfile types.String `tfsdk:"profile" vyos:"profile,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
			Description: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
		},

		// Nodes

	}
}
