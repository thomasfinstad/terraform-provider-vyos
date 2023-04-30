// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpPeerGroupTTLSecURIty describes the resource data model.
type VrfNameProtocolsBgpPeerGroupTTLSecURIty struct {
	// LeafNodes
	VrfNameProtocolsBgpPeerGroupTTLSecURItyHops customtypes.CustomStringValue `tfsdk:"hops" json:"hops,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupTTLSecURIty) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hops": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Number of the maximum number of hops to the BGP peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-254  |  Number of hops  |
`,
		},

		// TagNodes

		// Nodes

	}
}
