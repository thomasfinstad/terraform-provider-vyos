// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfLogAdjacencyChanges describes the resource data model.
type VrfNameProtocolsOspfLogAdjacencyChanges struct {
	// LeafNodes
	VrfNameProtocolsOspfLogAdjacencyChangesDetail customtypes.CustomStringValue `tfsdk:"detail" json:"detail,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfLogAdjacencyChanges) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"detail": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Log all state changes

`,
		},

		// TagNodes

		// Nodes

	}
}
