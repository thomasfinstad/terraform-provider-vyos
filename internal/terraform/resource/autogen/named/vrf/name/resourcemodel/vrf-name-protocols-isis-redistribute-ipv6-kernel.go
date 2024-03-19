// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvsixKernel{}

// VrfNameProtocolsIsisRedistributeIPvsixKernel describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixKernel struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixKernelLevelOne *VrfNameProtocolsIsisRedistributeIPvsixKernelLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixKernelLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixKernelLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixKernel) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixKernelLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
			Description: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixKernelLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
			Description: `Redistribute into level-2

`,
		},
	}
}
