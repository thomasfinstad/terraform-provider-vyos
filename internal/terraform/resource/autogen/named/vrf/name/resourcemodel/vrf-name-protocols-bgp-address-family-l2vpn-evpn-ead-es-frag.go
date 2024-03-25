// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFragEviLimit types.Number `tfsdk:"evi_limit" vyos:"evi-limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"evi_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `EVIs per-fragment

    |  Format  |  Description  |
    |----------|---------------|
    |  1-1000  |  limit        |
`,
			Description: `EVIs per-fragment

    |  Format  |  Description  |
    |----------|---------------|
    |  1-1000  |  limit        |
`,
		},

		// Nodes

	}
}
