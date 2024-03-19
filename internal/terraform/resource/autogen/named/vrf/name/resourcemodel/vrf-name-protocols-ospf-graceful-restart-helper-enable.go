// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfGracefulRestartHelperEnable{}

// VrfNameProtocolsOspfGracefulRestartHelperEnable describes the resource data model.
type VrfNameProtocolsOspfGracefulRestartHelperEnable struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfGracefulRestartHelperEnableRouterID types.List `tfsdk:"router_id" vyos:"router-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfGracefulRestartHelperEnable) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"router_id": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Advertising Router-ID

    |  Format  &emsp;|  Description                     |
    |----------------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |
`,
			Description: `Advertising Router-ID

    |  Format  |  Description                     |
    |----------------|----------------------------------|
    |  ipv4    |  Router-ID in IP address format  |
`,
		},

		// Nodes

	}
}
