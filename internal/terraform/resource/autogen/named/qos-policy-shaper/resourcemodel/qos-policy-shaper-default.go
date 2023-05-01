// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// QosPolicyShaperDefault describes the resource data model.
type QosPolicyShaperDefault struct {
	// LeafNodes
	LeafQosPolicyShaperDefaultBandwIDth    types.String `tfsdk:"bandwidth"`
	LeafQosPolicyShaperDefaultBurst        types.String `tfsdk:"burst"`
	LeafQosPolicyShaperDefaultCeiling      types.String `tfsdk:"ceiling"`
	LeafQosPolicyShaperDefaultCodelQuantum types.String `tfsdk:"codel_quantum"`
	LeafQosPolicyShaperDefaultFlows        types.String `tfsdk:"flows"`
	LeafQosPolicyShaperDefaultInterval     types.String `tfsdk:"interval"`
	LeafQosPolicyShaperDefaultPriority     types.String `tfsdk:"priority"`
	LeafQosPolicyShaperDefaultQueueLimit   types.String `tfsdk:"queue_limit"`
	LeafQosPolicyShaperDefaultQueueType    types.String `tfsdk:"queue_type"`
	LeafQosPolicyShaperDefaultSetDscp      types.String `tfsdk:"set_dscp"`
	LeafQosPolicyShaperDefaultTarget       types.String `tfsdk:"target"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperDefault) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "default"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperDefaultBandwIDth.IsNull() || o.LeafQosPolicyShaperDefaultBandwIDth.IsUnknown()) {
		vyosData["bandwidth"] = o.LeafQosPolicyShaperDefaultBandwIDth.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultBurst.IsNull() || o.LeafQosPolicyShaperDefaultBurst.IsUnknown()) {
		vyosData["burst"] = o.LeafQosPolicyShaperDefaultBurst.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultCeiling.IsNull() || o.LeafQosPolicyShaperDefaultCeiling.IsUnknown()) {
		vyosData["ceiling"] = o.LeafQosPolicyShaperDefaultCeiling.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultCodelQuantum.IsNull() || o.LeafQosPolicyShaperDefaultCodelQuantum.IsUnknown()) {
		vyosData["codel-quantum"] = o.LeafQosPolicyShaperDefaultCodelQuantum.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultFlows.IsNull() || o.LeafQosPolicyShaperDefaultFlows.IsUnknown()) {
		vyosData["flows"] = o.LeafQosPolicyShaperDefaultFlows.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultInterval.IsNull() || o.LeafQosPolicyShaperDefaultInterval.IsUnknown()) {
		vyosData["interval"] = o.LeafQosPolicyShaperDefaultInterval.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultPriority.IsNull() || o.LeafQosPolicyShaperDefaultPriority.IsUnknown()) {
		vyosData["priority"] = o.LeafQosPolicyShaperDefaultPriority.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultQueueLimit.IsNull() || o.LeafQosPolicyShaperDefaultQueueLimit.IsUnknown()) {
		vyosData["queue-limit"] = o.LeafQosPolicyShaperDefaultQueueLimit.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultQueueType.IsNull() || o.LeafQosPolicyShaperDefaultQueueType.IsUnknown()) {
		vyosData["queue-type"] = o.LeafQosPolicyShaperDefaultQueueType.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultSetDscp.IsNull() || o.LeafQosPolicyShaperDefaultSetDscp.IsUnknown()) {
		vyosData["set-dscp"] = o.LeafQosPolicyShaperDefaultSetDscp.ValueString()
	}
	if !(o.LeafQosPolicyShaperDefaultTarget.IsNull() || o.LeafQosPolicyShaperDefaultTarget.IsUnknown()) {
		vyosData["target"] = o.LeafQosPolicyShaperDefaultTarget.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperDefault) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "default"}})

	// Leafs
	if value, ok := vyosData["bandwidth"]; ok {
		o.LeafQosPolicyShaperDefaultBandwIDth = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultBandwIDth = basetypes.NewStringNull()
	}
	if value, ok := vyosData["burst"]; ok {
		o.LeafQosPolicyShaperDefaultBurst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultBurst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ceiling"]; ok {
		o.LeafQosPolicyShaperDefaultCeiling = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultCeiling = basetypes.NewStringNull()
	}
	if value, ok := vyosData["codel-quantum"]; ok {
		o.LeafQosPolicyShaperDefaultCodelQuantum = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultCodelQuantum = basetypes.NewStringNull()
	}
	if value, ok := vyosData["flows"]; ok {
		o.LeafQosPolicyShaperDefaultFlows = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultFlows = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interval"]; ok {
		o.LeafQosPolicyShaperDefaultInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["priority"]; ok {
		o.LeafQosPolicyShaperDefaultPriority = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultPriority = basetypes.NewStringNull()
	}
	if value, ok := vyosData["queue-limit"]; ok {
		o.LeafQosPolicyShaperDefaultQueueLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultQueueLimit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["queue-type"]; ok {
		o.LeafQosPolicyShaperDefaultQueueType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultQueueType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["set-dscp"]; ok {
		o.LeafQosPolicyShaperDefaultSetDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultSetDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["target"]; ok {
		o.LeafQosPolicyShaperDefaultTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperDefaultTarget = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "default"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperDefault) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"bandwidth":     types.StringType,
		"burst":         types.StringType,
		"ceiling":       types.StringType,
		"codel_quantum": types.StringType,
		"flows":         types.StringType,
		"interval":      types.StringType,
		"priority":      types.StringType,
		"queue_limit":   types.StringType,
		"queue_type":    types.StringType,
		"set_dscp":      types.StringType,
		"target":        types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperDefault) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bits per second  |
|  <number>bit  |  Bits per second  |
|  <number>kbit  |  Kilobits per second  |
|  <number>mbit  |  Megabits per second  |
|  <number>gbit  |  Gigabits per second  |
|  <number>tbit  |  Terabits per second  |
|  <number>%  |  Percentage of interface link speed  |

`,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Bytes  |
|  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"ceiling": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bandwidth limit for this class

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Rate in kbit (kilobit per second)  |
|  <number>%%  |  Percentage of overall rate  |
|  <number>bit  |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
|  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
|  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
|  <number>bps  |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,
		},

		"codel_quantum": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048576  |  Number of bytes used as 'deficit'  |

`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65536  |  Number of flows  |

`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Priority for usage of excess bandwidth

|  Format  |  Description  |
|----------|---------------|
|  u32:0-7  |  Priority order for bandwidth pool  |

`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"queue_limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Queue size in packets  |

`,
		},

		"queue_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

|  Format  |  Description  |
|----------|---------------|
|  drop-tail  |  First-In-First-Out (FIFO)  |
|  fair-queue  |  Stochastic Fair Queue (SFQ)  |
|  fq-codel  |  Fair Queue Codel  |
|  priority  |  Priority queuing  |
|  random-detect  |  Random Early Detection (RED)  |

`,

			// Default:          stringdefault.StaticString(`fq-codel`),
			Computed: true,
		},

		"set_dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Change the Differentiated Services (DiffServ) field in the IP header

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  Priority order for bandwidth pool  |
|  default  |  match DSCP (000000)  |
|  reliability  |  match DSCP (000001)  |
|  throughput  |  match DSCP (000010)  |
|  lowdelay  |  match DSCP (000100)  |
|  priority  |  match DSCP (001000)  |
|  immediate  |  match DSCP (010000)  |
|  flash  |  match DSCP (011000)  |
|  flash-override  |  match DSCP (100000)  |
|  critical  |  match DSCP (101000)  |
|  internet  |  match DSCP (110000)  |
|  network  |  match DSCP (111000)  |
|  AF11  |  High-throughput data  |
|  AF12  |  High-throughput data  |
|  AF13  |  High-throughput data  |
|  AF21  |  Low-latency data  |
|  AF22  |  Low-latency data  |
|  AF23  |  Low-latency data  |
|  AF31  |  Multimedia streaming  |
|  AF32  |  Multimedia streaming  |
|  AF33  |  Multimedia streaming  |
|  AF41  |  Multimedia conferencing  |
|  AF42  |  Multimedia conferencing  |
|  AF43  |  Multimedia conferencing  |
|  CS1  |  Low-priority data  |
|  CS2  |  OAM  |
|  CS3  |  Broadcast video  |
|  CS4  |  Real-time interactive  |
|  CS5  |  Signaling  |
|  CS6  |  Network control  |
|  CS7  |    |
|  EF  |  Expedited Forwarding  |

`,
		},

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Queue delay in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
