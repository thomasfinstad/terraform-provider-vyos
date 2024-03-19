// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths{}

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsEbgp types.Number `tfsdk:"ebgp" vyos:"ebgp,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPathsIbgp types.Number `tfsdk:"ibgp" vyos:"ibgp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastMaximumPaths) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ebgp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |
`,
			Description: `eBGP maximum paths

    |  Format  |  Description                  |
    |----------------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
		},

		"ibgp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |
`,
			Description: `iBGP maximum paths

    |  Format  |  Description                  |
    |----------------|-------------------------------|
    |  1-256   |  Number of paths to consider  |
`,
		},

		// Nodes

	}
}
