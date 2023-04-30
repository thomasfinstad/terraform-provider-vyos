// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpParametersBestpathPeerType describes the resource data model.
type VrfNameProtocolsBgpParametersBestpathPeerType struct {
	// LeafNodes
	VrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax customtypes.CustomStringValue `tfsdk:"multipath_relax" json:"multipath-relax,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersBestpathPeerType) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"multipath_relax": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Allow load sharing across routes learned from different peer types

`,
		},

		// TagNodes

		// Nodes

	}
}
