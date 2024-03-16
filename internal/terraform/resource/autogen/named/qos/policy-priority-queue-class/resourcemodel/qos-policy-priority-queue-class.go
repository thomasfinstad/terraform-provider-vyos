// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// QosPolicyPriorityQueueClass describes the resource data model.
type QosPolicyPriorityQueueClass struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:"-,self-id"`

	ParentIDQosPolicyPriorityQueue types.String `tfsdk:"priority_queue_id" vyos:"priority-queue,parent-id"`

	// LeafNodes
	LeafQosPolicyPriorityQueueClassDescrIPtion  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyPriorityQueueClassCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyPriorityQueueClassFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyPriorityQueueClassInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyPriorityQueueClassQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyPriorityQueueClassQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyPriorityQueueClassTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyPriorityQueueClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyPriorityQueueClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyPriorityQueueClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"qos",

		"policy",

		"priority-queue",
		o.ParentIDQosPolicyPriorityQueue.ValueString(),

		"class",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClass) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class Handle

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-7  &emsp; |  Priority  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"priority_queue_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Priority queuing based policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in priority_queue_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  priority_queue_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

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

			// Default:          stringdefault.StaticString(`drop-tail`),
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
