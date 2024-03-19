// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyShaperDefault{}

// QosPolicyShaperDefault describes the resource data model.
type QosPolicyShaperDefault struct {
	// LeafNodes
	LeafQosPolicyShaperDefaultBandwIDth    types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyShaperDefaultBurst        types.String `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafQosPolicyShaperDefaultCeiling      types.String `tfsdk:"ceiling" vyos:"ceiling,omitempty"`
	LeafQosPolicyShaperDefaultCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyShaperDefaultFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyShaperDefaultInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyShaperDefaultPriority     types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafQosPolicyShaperDefaultQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyShaperDefaultQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyShaperDefaultSetDscp      types.String `tfsdk:"set_dscp" vyos:"set-dscp,omitempty"`
	LeafQosPolicyShaperDefaultTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
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
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
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
			MarkdownDescription: `Priority for usage of excess bandwidth

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-7     |  Priority order for bandwidth pool  |
`,
			Description: `Priority for usage of excess bandwidth

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-7     |  Priority order for bandwidth pool  |
`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
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
