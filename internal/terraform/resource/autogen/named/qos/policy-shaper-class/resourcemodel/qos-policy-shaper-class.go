// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyShaperClass describes the resource data model.
type QosPolicyShaperClass struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:",self-id"`

	ParentIDQosPolicyShaper types.String `tfsdk:"shaper" vyos:"shaper,parent-id"`

	// LeafNodes
	LeafQosPolicyShaperClassDescrIPtion  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperClassBandwIDth    types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyShaperClassBurst        types.String `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafQosPolicyShaperClassCeiling      types.String `tfsdk:"ceiling" vyos:"ceiling,omitempty"`
	LeafQosPolicyShaperClassCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyShaperClassFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyShaperClassInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyShaperClassPriority     types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafQosPolicyShaperClassQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyShaperClassQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyShaperClassSetDscp      types.String `tfsdk:"set_dscp" vyos:"set-dscp,omitempty"`
	LeafQosPolicyShaperClassTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyShaperClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyShaperClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"qos",

		"policy",

		"shaper",
		o.ParentIDQosPolicyShaper.ValueString(),

		"class",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClass) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"class_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-4095  &emsp; |  Class Identifier  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"shaper_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic shaping based policy (Hierarchy Token Bucket)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
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

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  auto  &emsp; |  Bandwidth matches interface speed  |
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bytes  |
    |  <number><suffix>  &emsp; |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"ceiling": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bandwidth limit for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Rate in kbit (kilobit per second)  |
    |  <number>%%  &emsp; |  Percentage of overall rate  |
    |  <number>bit  &emsp; |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
    |  <number>ibit  &emsp; |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  &emsp; |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps  &emsp; |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

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

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority for rule evaluation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-20  &emsp; |  Priority for match rule evaluation  |

`,
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

			// Default:          stringdefault.StaticString(`fq-codel`),
			Computed: true,
		},

		"set_dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Change the Differentiated Services (DiffServ) field in the IP header

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  Priority order for bandwidth pool  |
    |  default  &emsp; |  match DSCP (000000)  |
    |  reliability  &emsp; |  match DSCP (000001)  |
    |  throughput  &emsp; |  match DSCP (000010)  |
    |  lowdelay  &emsp; |  match DSCP (000100)  |
    |  priority  &emsp; |  match DSCP (001000)  |
    |  immediate  &emsp; |  match DSCP (010000)  |
    |  flash  &emsp; |  match DSCP (011000)  |
    |  flash-override  &emsp; |  match DSCP (100000)  |
    |  critical  &emsp; |  match DSCP (101000)  |
    |  internet  &emsp; |  match DSCP (110000)  |
    |  network  &emsp; |  match DSCP (111000)  |
    |  AF11  &emsp; |  High-throughput data  |
    |  AF12  &emsp; |  High-throughput data  |
    |  AF13  &emsp; |  High-throughput data  |
    |  AF21  &emsp; |  Low-latency data  |
    |  AF22  &emsp; |  Low-latency data  |
    |  AF23  &emsp; |  Low-latency data  |
    |  AF31  &emsp; |  Multimedia streaming  |
    |  AF32  &emsp; |  Multimedia streaming  |
    |  AF33  &emsp; |  Multimedia streaming  |
    |  AF41  &emsp; |  Multimedia conferencing  |
    |  AF42  &emsp; |  Multimedia conferencing  |
    |  AF43  &emsp; |  Multimedia conferencing  |
    |  CS1  &emsp; |  Low-priority data  |
    |  CS2  &emsp; |  OAM  |
    |  CS3  &emsp; |  Broadcast video  |
    |  CS4  &emsp; |  Real-time interactive  |
    |  CS5  &emsp; |  Signaling  |
    |  CS6  &emsp; |  Network control  |
    |  CS7  &emsp; |    |
    |  EF  &emsp; |  Expedited Forwarding  |

`,
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
