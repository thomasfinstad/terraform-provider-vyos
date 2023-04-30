// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// SystemFlowAccountingSflowServer describes the resource data model.
type SystemFlowAccountingSflowServer struct {
	// LeafNodes
	SystemFlowAccountingSflowServerPort customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o SystemFlowAccountingSflowServer) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `sFlow port number

|  Format  |  Description  |
|----------|---------------|
|  u32:1025-65535  |  sFlow port number  |
`,

			// Default:          stringdefault.StaticString(`6343`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
