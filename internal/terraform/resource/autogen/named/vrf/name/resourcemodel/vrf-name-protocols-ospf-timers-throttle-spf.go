// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfTimersThroTTLeSpf{}

// VrfNameProtocolsOspfTimersThroTTLeSpf describes the resource data model.
type VrfNameProtocolsOspfTimersThroTTLeSpf struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfDelay           types.Number `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfInitialHoldtime types.Number `tfsdk:"initial_holdtime" vyos:"initial-holdtime,omitempty"`
	LeafVrfNameProtocolsOspfTimersThroTTLeSpfMaxHoldtime     types.Number `tfsdk:"max_holdtime" vyos:"max-holdtime,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfTimersThroTTLeSpf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay from the first change received to SPF calculation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`200`),
			Computed: true,
		},

		"initial_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Initial hold time between consecutive SPF calculations

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Initial hold time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"max_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Max hold time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		// Nodes

	}
}
