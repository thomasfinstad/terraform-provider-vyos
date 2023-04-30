// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ContainerNameNetwork describes the resource data model.
type ContainerNameNetwork struct {
	// LeafNodes
	ContainerNameNetworkAddress customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ContainerNameNetwork) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Assign static IP address to container

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
