// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsIsisSpfDelayIetf describes the resource data model.
type ProtocolsIsisSpfDelayIetf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsIsisSpfDelayIetfInitDelay   types.Number `tfsdk:"init_delay" vyos:"init-delay,omitempty"`
	LeafProtocolsIsisSpfDelayIetfShortDelay  types.Number `tfsdk:"short_delay" vyos:"short-delay,omitempty"`
	LeafProtocolsIsisSpfDelayIetfLongDelay   types.Number `tfsdk:"long_delay" vyos:"long-delay,omitempty"`
	LeafProtocolsIsisSpfDelayIetfHolddown    types.Number `tfsdk:"holddown" vyos:"holddown,omitempty"`
	LeafProtocolsIsisSpfDelayIetfTimeToLearn types.Number `tfsdk:"time_to_learn" vyos:"time-to-learn,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsIsisSpfDelayIetf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIsisSpfDelayIetf) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"isis",

		"spf-delay-ietf",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisSpfDelayIetf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"init_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in QUIET state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in QUIET state (in ms)  |

`,
		},

		"short_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in SHORT_WAIT state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in SHORT_WAIT state (in ms)  |

`,
		},

		"long_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in LONG_WAIT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in LONG_WAIT state in ms  |

`,
		},

		"holddown": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time with no received IGP events before considering IGP stable

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Time with no received IGP events before considering IGP stable in ms  |

`,
		},

		"time_to_learn": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum duration needed to learn all the events related to a single failure

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Maximum duration needed to learn all the events related to a single failure in ms  |

`,
		},
	}
}
