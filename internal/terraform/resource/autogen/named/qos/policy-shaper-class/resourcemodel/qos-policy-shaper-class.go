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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperClass{}

// QosPolicyShaperClass describes the resource data model.
type QosPolicyShaperClass struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:"-,self-id"`

	ParentIDQosPolicyShaper types.String `tfsdk:"shaper_id" vyos:"shaper,parent-id"`

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

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyShaperClass) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"class",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyShaperClass) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper",
		o.ParentIDQosPolicyShaper.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyShaperClass) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper",
		o.ParentIDQosPolicyShaper.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClass) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  2-4095  |  Class Identifier  |
`,
			Description: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  2-4095  |  Class Identifier  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"shaper_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic shaping based policy (Hierarchy Token Bucket)

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Traffic shaping based policy (Hierarchy Token Bucket)

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in shaper_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  shaper_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  auto          |  Bandwidth matches interface speed   |
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
			Description: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  auto          |  Bandwidth matches interface speed   |
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

    |  Format            |  Description                             |
    |--------------------|------------------------------------------|
    |  <number>          |  Bytes                                   |
    |  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,
			Description: `Burst size for this class

    |  Format            |  Description                             |
    |--------------------|------------------------------------------|
    |  <number>          |  Bytes                                   |
    |  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"ceiling": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bandwidth limit for this class

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
			Description: `Bandwidth limit for this class

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
		},

		"codel_quantum": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

    |  Format     |  Description                        |
    |-------------|-------------------------------------|
    |  0-1048576  |  Number of bytes used as 'deficit'  |
`,
			Description: `Deficit in the fair queuing algorithm

    |  Format     |  Description                        |
    |-------------|-------------------------------------|
    |  0-1048576  |  Number of bytes used as 'deficit'  |
`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65536  |  Number of flows  |
`,
			Description: `Number of flows into which the incoming packets are classified

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65536  |  Number of flows  |
`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Interval in milliseconds  |
`,
			Description: `Interval used to measure the delay

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority for rule evaluation

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  0-20    |  Priority for match rule evaluation  |
`,
			Description: `Priority for rule evaluation

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  0-20    |  Priority for match rule evaluation  |
`,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
			Description: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
		},

		"queue_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

    |  Format         |  Description                   |
    |-----------------|--------------------------------|
    |  drop-tail      |  First-In-First-Out (FIFO)     |
    |  fair-queue     |  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       |  Fair Queue Codel              |
    |  priority       |  Priority queuing              |
    |  random-detect  |  Random Early Detection (RED)  |
`,
			Description: `Queue type for default traffic

    |  Format         |  Description                   |
    |-----------------|--------------------------------|
    |  drop-tail      |  First-In-First-Out (FIFO)     |
    |  fair-queue     |  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       |  Fair Queue Codel              |
    |  priority       |  Priority queuing              |
    |  random-detect  |  Random Early Detection (RED)  |
`,

			// Default:          stringdefault.StaticString(`fq-codel`),
			Computed: true,
		},

		"set_dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Change the Differentiated Services (DiffServ) field in the IP header

    |  Format          |  Description                        |
    |------------------|-------------------------------------|
    |  0-63            |  Priority order for bandwidth pool  |
    |  default         |  match DSCP (000000)                |
    |  reliability     |  match DSCP (000001)                |
    |  throughput      |  match DSCP (000010)                |
    |  lowdelay        |  match DSCP (000100)                |
    |  priority        |  match DSCP (001000)                |
    |  immediate       |  match DSCP (010000)                |
    |  flash           |  match DSCP (011000)                |
    |  flash-override  |  match DSCP (100000)                |
    |  critical        |  match DSCP (101000)                |
    |  internet        |  match DSCP (110000)                |
    |  network         |  match DSCP (111000)                |
    |  AF11            |  High-throughput data               |
    |  AF12            |  High-throughput data               |
    |  AF13            |  High-throughput data               |
    |  AF21            |  Low-latency data                   |
    |  AF22            |  Low-latency data                   |
    |  AF23            |  Low-latency data                   |
    |  AF31            |  Multimedia streaming               |
    |  AF32            |  Multimedia streaming               |
    |  AF33            |  Multimedia streaming               |
    |  AF41            |  Multimedia conferencing            |
    |  AF42            |  Multimedia conferencing            |
    |  AF43            |  Multimedia conferencing            |
    |  CS1             |  Low-priority data                  |
    |  CS2             |  OAM                                |
    |  CS3             |  Broadcast video                    |
    |  CS4             |  Real-time interactive              |
    |  CS5             |  Signaling                          |
    |  CS6             |  Network control                    |
    |  CS7             |  N/A                                |
    |  EF              |  Expedited Forwarding               |
`,
			Description: `Change the Differentiated Services (DiffServ) field in the IP header

    |  Format          |  Description                        |
    |------------------|-------------------------------------|
    |  0-63            |  Priority order for bandwidth pool  |
    |  default         |  match DSCP (000000)                |
    |  reliability     |  match DSCP (000001)                |
    |  throughput      |  match DSCP (000010)                |
    |  lowdelay        |  match DSCP (000100)                |
    |  priority        |  match DSCP (001000)                |
    |  immediate       |  match DSCP (010000)                |
    |  flash           |  match DSCP (011000)                |
    |  flash-override  |  match DSCP (100000)                |
    |  critical        |  match DSCP (101000)                |
    |  internet        |  match DSCP (110000)                |
    |  network         |  match DSCP (111000)                |
    |  AF11            |  High-throughput data               |
    |  AF12            |  High-throughput data               |
    |  AF13            |  High-throughput data               |
    |  AF21            |  Low-latency data                   |
    |  AF22            |  Low-latency data                   |
    |  AF23            |  Low-latency data                   |
    |  AF31            |  Multimedia streaming               |
    |  AF32            |  Multimedia streaming               |
    |  AF33            |  Multimedia streaming               |
    |  AF41            |  Multimedia conferencing            |
    |  AF42            |  Multimedia conferencing            |
    |  AF43            |  Multimedia conferencing            |
    |  CS1             |  Low-priority data                  |
    |  CS2             |  OAM                                |
    |  CS3             |  Broadcast video                    |
    |  CS4             |  Real-time interactive              |
    |  CS5             |  Signaling                          |
    |  CS6             |  Network control                    |
    |  CS7             |  N/A                                |
    |  EF              |  Expedited Forwarding               |
`,
		},

		"target": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Queue delay in milliseconds  |
`,
			Description: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Queue delay in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}
