// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical{}

// VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCriticalLevelOne types.Bool `tfsdk:"level_1" vyos:"level-1,omitempty"`
	LeafVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCriticalLevelTwo types.Bool `tfsdk:"level_2" vyos:"level-2,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"level_1": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match on IS-IS level-1 routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"level_2": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match on IS-IS level-2 routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
