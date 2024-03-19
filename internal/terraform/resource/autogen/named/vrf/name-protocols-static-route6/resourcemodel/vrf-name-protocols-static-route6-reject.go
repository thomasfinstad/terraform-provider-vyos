// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsStaticRoutesixReject{}

// VrfNameProtocolsStaticRoutesixReject describes the resource data model.
type VrfNameProtocolsStaticRoutesixReject struct {
	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixRejectDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixRejectTag      types.Number `tfsdk:"tag" vyos:"tag,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixReject) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
			Description: `Distance for this route

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-255   |  Distance for this route  |
`,
		},

		"tag": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Tag value for this route

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967295  |  Tag value for this route  |
`,
			Description: `Tag value for this route

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967295  |  Tag value for this route  |
`,
		},

		// Nodes

	}
}
