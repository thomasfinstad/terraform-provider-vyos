// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeAutoCost{}

// VrfNameProtocolsOspfvthreeAutoCost describes the resource data model.
type VrfNameProtocolsOspfvthreeAutoCost struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAutoCostReferenceBandwIDth types.Number `tfsdk:"reference_bandwidth" vyos:"reference-bandwidth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAutoCost) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"reference_bandwidth": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Reference bandwidth method to assign cost

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967  &emsp; |  Reference bandwidth cost in Mbits/sec  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		// Nodes

	}
}
