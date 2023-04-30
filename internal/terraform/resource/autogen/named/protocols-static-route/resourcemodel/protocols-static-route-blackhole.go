// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsStaticRouteBlackhole describes the resource data model.
type ProtocolsStaticRouteBlackhole struct {
	// LeafNodes
	ProtocolsStaticRouteBlackholeDistance customtypes.CustomStringValue `tfsdk:"distance" json:"distance,omitempty"`
	ProtocolsStaticRouteBlackholeTag      customtypes.CustomStringValue `tfsdk:"tag" json:"tag,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsStaticRouteBlackhole) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"distance": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Distance for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Distance for this route  |
`,
		},

		"tag": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Tag value for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Tag value for this route  |
`,
		},

		// TagNodes

		// Nodes

	}
}
