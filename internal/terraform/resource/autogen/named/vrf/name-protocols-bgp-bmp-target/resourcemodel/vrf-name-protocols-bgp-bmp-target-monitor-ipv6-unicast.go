// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast describes the resource data model.
type VrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicastPrePolicy  types.Bool `tfsdk:"pre_policy" vyos:"pre-policy,omitempty"`
	LeafVrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicastPostPolicy types.Bool `tfsdk:"post_policy" vyos:"post-policy,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpBmpTargetMonitorIPvsixUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"pre_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send state before policy and filter processing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"post_policy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send state with policy and filters applied

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}