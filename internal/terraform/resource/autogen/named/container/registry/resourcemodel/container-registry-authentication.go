// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ContainerRegistryAuthentication{}

// ContainerRegistryAuthentication describes the resource data model.
type ContainerRegistryAuthentication struct {
	// LeafNodes
	LeafContainerRegistryAuthenticationUsername types.String `tfsdk:"username" vyos:"username,omitempty"`
	LeafContainerRegistryAuthenticationPassword types.String `tfsdk:"password" vyos:"password,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerRegistryAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
			Description: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Username     |
`,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
			Description: `Password used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Password     |
`,
		},

		// Nodes

	}
}
