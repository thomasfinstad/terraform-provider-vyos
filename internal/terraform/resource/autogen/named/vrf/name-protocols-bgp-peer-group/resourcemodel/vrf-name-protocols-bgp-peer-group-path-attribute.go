// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupPathAttribute{}

// VrfNameProtocolsBgpPeerGroupPathAttribute describes the resource data model.
type VrfNameProtocolsBgpPeerGroupPathAttribute struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupPathAttributeDiscard         types.List   `tfsdk:"discard" vyos:"discard,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPathAttributeTreatAsWithdraw types.Number `tfsdk:"treat_as_withdraw" vyos:"treat-as-withdraw,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupPathAttribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"discard": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Drop specified attributes from incoming UPDATE messages

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
			Description: `Drop specified attributes from incoming UPDATE messages

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
		},

		"treat_as_withdraw": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
			Description: `Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-255   |  Attribute number  |
`,
		},

		// Nodes

	}
}
