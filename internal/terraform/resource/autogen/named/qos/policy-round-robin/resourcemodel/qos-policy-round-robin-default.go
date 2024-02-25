// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyRoundRobinDefault describes the resource data model.
type QosPolicyRoundRobinDefault struct {
	// LeafNodes
	LeafQosPolicyRoundRobinDefaultCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyRoundRobinDefaultFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyRoundRobinDefaultInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyRoundRobinDefaultQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyRoundRobinDefaultQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyRoundRobinDefaultTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"codel_quantum": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-1048576  &emsp; |  Number of bytes used as 'deficit'  |

`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65536  &emsp; |  Number of flows  |

`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Queue size in packets  |

`,
		},

		"queue_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  drop-tail  &emsp; |  First-In-First-Out (FIFO)  |
    |  fair-queue  &emsp; |  Stochastic Fair Queue (SFQ)  |
    |  fq-codel  &emsp; |  Fair Queue Codel  |
    |  priority  &emsp; |  Priority queuing  |
    |  random-detect  &emsp; |  Random Early Detection (RED)  |

`,

			// Default:          stringdefault.StaticString(`fair-queue`),
			Computed: true,
		},

		"target": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Queue delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}