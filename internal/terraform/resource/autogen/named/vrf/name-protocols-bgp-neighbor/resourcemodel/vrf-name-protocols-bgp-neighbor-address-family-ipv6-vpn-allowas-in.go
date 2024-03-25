// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
			Description: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
		},

		// Nodes

	}
}
