// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisRedistributeIPvfourRIP describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfourRIP struct {
	// LeafNodes

	// TagNodes

	// Nodes
	VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne types.Object `tfsdk:"level_1" json:"level-1,omitempty"`
	VrfNameProtocolsIsisRedistributeIPvfourRIPLevelTwo types.Object `tfsdk:"level_2" json:"level-2,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourRIP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourRIPLevelOne{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourRIPLevelTwo{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}
