// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpListen describes the resource data model.
type VrfNameProtocolsBgpListen struct {
	// LeafNodes
	VrfNameProtocolsBgpListenLimit customtypes.CustomStringValue `tfsdk:"limit" json:"limit,omitempty"`

	// TagNodes
	VrfNameProtocolsBgpListenRange types.Map `tfsdk:"range" json:"range,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpListen) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"limit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Maximum number of dynamic neighbors that can be created

|  Format  |  Description  |
|----------|---------------|
|  u32:1-5000  |  BGP neighbor limit  |
`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpListenRange{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `BGP dynamic neighbors listen range

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 dynamic neighbors listen range  |
|  ipv6net  |  IPv6 dynamic neighbors listen range  |
`,
		},

		// Nodes

	}
}
