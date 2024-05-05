// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvsixConnected{}

// VrfNameProtocolsIsisRedistributeIPvsixConnected describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixConnected struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne *VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixConnected) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
			Description: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixConnectedLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
			Description: `Redistribute into level-2

`,
		},
	}
}
