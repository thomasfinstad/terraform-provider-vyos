// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBabelInterface describes the resource data model.
type ProtocolsBabelInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBabelInterfaceType             types.String `tfsdk:"type" json:"type,omitempty"`
	LeafProtocolsBabelInterfaceSplitHorizon     types.String `tfsdk:"split_horizon" json:"split-horizon,omitempty"`
	LeafProtocolsBabelInterfaceHelloInterval    types.String `tfsdk:"hello_interval" json:"hello-interval,omitempty"`
	LeafProtocolsBabelInterfaceUpdateInterval   types.String `tfsdk:"update_interval" json:"update-interval,omitempty"`
	LeafProtocolsBabelInterfaceRxcost           types.String `tfsdk:"rxcost" json:"rxcost,omitempty"`
	LeafProtocolsBabelInterfaceRttDecay         types.String `tfsdk:"rtt_decay" json:"rtt-decay,omitempty"`
	LeafProtocolsBabelInterfaceRttMin           types.String `tfsdk:"rtt_min" json:"rtt-min,omitempty"`
	LeafProtocolsBabelInterfaceRttMax           types.String `tfsdk:"rtt_max" json:"rtt-max,omitempty"`
	LeafProtocolsBabelInterfaceMaxRttPenalty    types.String `tfsdk:"max_rtt_penalty" json:"max-rtt-penalty,omitempty"`
	LeafProtocolsBabelInterfaceEnableTimestamps types.String `tfsdk:"enable_timestamps" json:"enable-timestamps,omitempty"`
	LeafProtocolsBabelInterfaceChannel          types.String `tfsdk:"channel" json:"channel,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"babel",
		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		// LeafNodes

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface type

|  Format  |  Description  |
|----------|---------------|
|  auto  |  Automatically detect interface type  |
|  wired  |  Wired interface  |
|  wireless  |  Wireless interface  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"split_horizon": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Split horizon parameters

|  Format  |  Description  |
|----------|---------------|
|  default  |  Enable on wired interfaces, and disable on wireless interfaces  |
|  enable  |  Enable split horizon processing  |
|  disable  |  Disable split horizon processing  |

`,

			// Default:          stringdefault.StaticString(`default`),
			Computed: true,
		},

		"hello_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled hellos

|  Format  |  Description  |
|----------|---------------|
|  u32:20-655340  |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`4000`),
			Computed: true,
		},

		"update_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time between scheduled updates

|  Format  |  Description  |
|----------|---------------|
|  u32:20-655340  |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`20000`),
			Computed: true,
		},

		"rxcost": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Base receive cost for this interface

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65534  |  Base receive cost  |

`,
		},

		"rtt_decay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Decay factor for exponential moving average of RTT samples

|  Format  |  Description  |
|----------|---------------|
|  u32:1-256  |  Decay factor, in units of 1/256  |

`,

			// Default:          stringdefault.StaticString(`42`),
			Computed: true,
		},

		"rtt_min": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum RTT

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"rtt_max": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum RTT

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Milliseconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"max_rtt_penalty": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum additional cost due to RTT

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Milliseconds (0 to disable the use of RTT-based cost)  |

`,

			// Default:          stringdefault.StaticString(`150`),
			Computed: true,
		},

		"enable_timestamps": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable timestamps with each Hello and IHU message in order to compute RTT values

`,
		},

		"channel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Channel number for diversity routing

|  Format  |  Description  |
|----------|---------------|
|  u32:1-254  |  Interfaces with a channel number interfere with interfering interfaces and interfaces with the same channel number  |
|  interfering  |  Interfering interfaces are assumed to interfere with all other channels except non-interfering channels  |
|  non-interfering  |  Non-interfering interfaces only interfere with themselves  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBabelInterface) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBabelInterfaceType.IsNull() && !o.LeafProtocolsBabelInterfaceType.IsUnknown() {
		jsonData["type"] = o.LeafProtocolsBabelInterfaceType.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceSplitHorizon.IsNull() && !o.LeafProtocolsBabelInterfaceSplitHorizon.IsUnknown() {
		jsonData["split-horizon"] = o.LeafProtocolsBabelInterfaceSplitHorizon.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceHelloInterval.IsNull() && !o.LeafProtocolsBabelInterfaceHelloInterval.IsUnknown() {
		jsonData["hello-interval"] = o.LeafProtocolsBabelInterfaceHelloInterval.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceUpdateInterval.IsNull() && !o.LeafProtocolsBabelInterfaceUpdateInterval.IsUnknown() {
		jsonData["update-interval"] = o.LeafProtocolsBabelInterfaceUpdateInterval.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceRxcost.IsNull() && !o.LeafProtocolsBabelInterfaceRxcost.IsUnknown() {
		jsonData["rxcost"] = o.LeafProtocolsBabelInterfaceRxcost.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceRttDecay.IsNull() && !o.LeafProtocolsBabelInterfaceRttDecay.IsUnknown() {
		jsonData["rtt-decay"] = o.LeafProtocolsBabelInterfaceRttDecay.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceRttMin.IsNull() && !o.LeafProtocolsBabelInterfaceRttMin.IsUnknown() {
		jsonData["rtt-min"] = o.LeafProtocolsBabelInterfaceRttMin.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceRttMax.IsNull() && !o.LeafProtocolsBabelInterfaceRttMax.IsUnknown() {
		jsonData["rtt-max"] = o.LeafProtocolsBabelInterfaceRttMax.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceMaxRttPenalty.IsNull() && !o.LeafProtocolsBabelInterfaceMaxRttPenalty.IsUnknown() {
		jsonData["max-rtt-penalty"] = o.LeafProtocolsBabelInterfaceMaxRttPenalty.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceEnableTimestamps.IsNull() && !o.LeafProtocolsBabelInterfaceEnableTimestamps.IsUnknown() {
		jsonData["enable-timestamps"] = o.LeafProtocolsBabelInterfaceEnableTimestamps.ValueString()
	}

	if !o.LeafProtocolsBabelInterfaceChannel.IsNull() && !o.LeafProtocolsBabelInterfaceChannel.IsUnknown() {
		jsonData["channel"] = o.LeafProtocolsBabelInterfaceChannel.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBabelInterface) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["type"]; ok {
		o.LeafProtocolsBabelInterfaceType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["split-horizon"]; ok {
		o.LeafProtocolsBabelInterfaceSplitHorizon = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceSplitHorizon = basetypes.NewStringNull()
	}

	if value, ok := jsonData["hello-interval"]; ok {
		o.LeafProtocolsBabelInterfaceHelloInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceHelloInterval = basetypes.NewStringNull()
	}

	if value, ok := jsonData["update-interval"]; ok {
		o.LeafProtocolsBabelInterfaceUpdateInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceUpdateInterval = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rxcost"]; ok {
		o.LeafProtocolsBabelInterfaceRxcost = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceRxcost = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rtt-decay"]; ok {
		o.LeafProtocolsBabelInterfaceRttDecay = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceRttDecay = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rtt-min"]; ok {
		o.LeafProtocolsBabelInterfaceRttMin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceRttMin = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rtt-max"]; ok {
		o.LeafProtocolsBabelInterfaceRttMax = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceRttMax = basetypes.NewStringNull()
	}

	if value, ok := jsonData["max-rtt-penalty"]; ok {
		o.LeafProtocolsBabelInterfaceMaxRttPenalty = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceMaxRttPenalty = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-timestamps"]; ok {
		o.LeafProtocolsBabelInterfaceEnableTimestamps = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceEnableTimestamps = basetypes.NewStringNull()
	}

	if value, ok := jsonData["channel"]; ok {
		o.LeafProtocolsBabelInterfaceChannel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBabelInterfaceChannel = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
