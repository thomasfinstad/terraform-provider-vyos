// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBfdPeerInterval describes the resource data model.
type ProtocolsBfdPeerInterval struct {
	// LeafNodes
	LeafProtocolsBfdPeerIntervalReceive      types.String `tfsdk:"receive" json:"receive,omitempty"`
	LeafProtocolsBfdPeerIntervalTransmit     types.String `tfsdk:"transmit" json:"transmit,omitempty"`
	LeafProtocolsBfdPeerIntervalMultIPlier   types.String `tfsdk:"multiplier" json:"multiplier,omitempty"`
	LeafProtocolsBfdPeerIntervalEchoInterval types.String `tfsdk:"echo_interval" json:"echo-interval,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBfdPeerInterval) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval of receiving control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"transmit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval of transmitting control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"multiplier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Multiplier to determine packet loss

|  Format  |  Description  |
|----------|---------------|
|  u32:2-255  |  Remote transmission interval will be multiplied by this value  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"echo_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Echo receive transmission interval

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  The minimal echo receive transmission interval that this system is capable of handling  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBfdPeerInterval) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBfdPeerIntervalReceive.IsNull() && !o.LeafProtocolsBfdPeerIntervalReceive.IsUnknown() {
		jsonData["receive"] = o.LeafProtocolsBfdPeerIntervalReceive.ValueString()
	}

	if !o.LeafProtocolsBfdPeerIntervalTransmit.IsNull() && !o.LeafProtocolsBfdPeerIntervalTransmit.IsUnknown() {
		jsonData["transmit"] = o.LeafProtocolsBfdPeerIntervalTransmit.ValueString()
	}

	if !o.LeafProtocolsBfdPeerIntervalMultIPlier.IsNull() && !o.LeafProtocolsBfdPeerIntervalMultIPlier.IsUnknown() {
		jsonData["multiplier"] = o.LeafProtocolsBfdPeerIntervalMultIPlier.ValueString()
	}

	if !o.LeafProtocolsBfdPeerIntervalEchoInterval.IsNull() && !o.LeafProtocolsBfdPeerIntervalEchoInterval.IsUnknown() {
		jsonData["echo-interval"] = o.LeafProtocolsBfdPeerIntervalEchoInterval.ValueString()
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
func (o *ProtocolsBfdPeerInterval) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["receive"]; ok {
		o.LeafProtocolsBfdPeerIntervalReceive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerIntervalReceive = basetypes.NewStringNull()
	}

	if value, ok := jsonData["transmit"]; ok {
		o.LeafProtocolsBfdPeerIntervalTransmit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerIntervalTransmit = basetypes.NewStringNull()
	}

	if value, ok := jsonData["multiplier"]; ok {
		o.LeafProtocolsBfdPeerIntervalMultIPlier = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerIntervalMultIPlier = basetypes.NewStringNull()
	}

	if value, ok := jsonData["echo-interval"]; ok {
		o.LeafProtocolsBfdPeerIntervalEchoInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerIntervalEchoInterval = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
