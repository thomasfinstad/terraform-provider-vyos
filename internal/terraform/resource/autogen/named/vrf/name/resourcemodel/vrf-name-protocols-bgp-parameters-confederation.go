// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpParametersConfederation{}

// VrfNameProtocolsBgpParametersConfederation describes the resource data model.
type VrfNameProtocolsBgpParametersConfederation struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersConfederationIDentifier types.Number `tfsdk:"identifier" vyos:"identifier,omitempty"`
	LeafVrfNameProtocolsBgpParametersConfederationPeers      types.List   `tfsdk:"peers" vyos:"peers,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersConfederation) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"identifier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Confederation AS identifier

    |  Format        |  Description          |
    |----------------|-----------------------|
    |  1-4294967294  |  Confederation AS id  |
`,
			Description: `Confederation AS identifier

    |  Format        |  Description          |
    |----------------|-----------------------|
    |  1-4294967294  |  Confederation AS id  |
`,
		},

		"peers": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Peer ASs in the BGP confederation

    |  Format        |  Description     |
    |----------------|------------------|
    |  1-4294967294  |  Peer AS number  |
`,
			Description: `Peer ASs in the BGP confederation

    |  Format        |  Description     |
    |----------------|------------------|
    |  1-4294967294  |  Peer AS number  |
`,
		},

		// Nodes

	}
}
