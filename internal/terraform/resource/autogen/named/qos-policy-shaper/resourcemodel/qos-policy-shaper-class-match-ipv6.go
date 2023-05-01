// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyShaperClassMatchIPvsix describes the resource data model.
type QosPolicyShaperClassMatchIPvsix struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPvsixDscp      types.String `tfsdk:"dscp" json:"dscp,omitempty"`
	LeafQosPolicyShaperClassMatchIPvsixMaxLength types.String `tfsdk:"max_length" json:"max-length,omitempty"`
	LeafQosPolicyShaperClassMatchIPvsixProtocol  types.String `tfsdk:"protocol" json:"protocol,omitempty"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperClassMatchIPvsixDestination *QosPolicyShaperClassMatchIPvsixDestination `tfsdk:"destination" json:"destination,omitempty"`
	NodeQosPolicyShaperClassMatchIPvsixSource      *QosPolicyShaperClassMatchIPvsixSource      `tfsdk:"source" json:"source,omitempty"`
	NodeQosPolicyShaperClassMatchIPvsixTCP         *QosPolicyShaperClassMatchIPvsixTCP         `tfsdk:"tcp" json:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  Differentiated Services Codepoint (DSCP) value   |
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

		"max_length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet length

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Maximum packet/payload length  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Protocol name  |

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperClassMatchIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyShaperClassMatchIPvsixDscp.IsNull() && !o.LeafQosPolicyShaperClassMatchIPvsixDscp.IsUnknown() {
		jsonData["dscp"] = o.LeafQosPolicyShaperClassMatchIPvsixDscp.ValueString()
	}

	if !o.LeafQosPolicyShaperClassMatchIPvsixMaxLength.IsNull() && !o.LeafQosPolicyShaperClassMatchIPvsixMaxLength.IsUnknown() {
		jsonData["max-length"] = o.LeafQosPolicyShaperClassMatchIPvsixMaxLength.ValueString()
	}

	if !o.LeafQosPolicyShaperClassMatchIPvsixProtocol.IsNull() && !o.LeafQosPolicyShaperClassMatchIPvsixProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafQosPolicyShaperClassMatchIPvsixProtocol.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeQosPolicyShaperClassMatchIPvsixDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperClassMatchIPvsixDestination)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["destination"] = subData
	}

	if !reflect.ValueOf(o.NodeQosPolicyShaperClassMatchIPvsixSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperClassMatchIPvsixSource)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["source"] = subData
	}

	if !reflect.ValueOf(o.NodeQosPolicyShaperClassMatchIPvsixTCP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeQosPolicyShaperClassMatchIPvsixTCP)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["tcp"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperClassMatchIPvsix) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["dscp"]; ok {
		o.LeafQosPolicyShaperClassMatchIPvsixDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPvsixDscp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["max-length"]; ok {
		o.LeafQosPolicyShaperClassMatchIPvsixMaxLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPvsixMaxLength = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafQosPolicyShaperClassMatchIPvsixProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPvsixProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperClassMatchIPvsixDestination = &QosPolicyShaperClassMatchIPvsixDestination{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperClassMatchIPvsixDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperClassMatchIPvsixSource = &QosPolicyShaperClassMatchIPvsixSource{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperClassMatchIPvsixSource)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["tcp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeQosPolicyShaperClassMatchIPvsixTCP = &QosPolicyShaperClassMatchIPvsixTCP{}

		err = json.Unmarshal(subJSONStr, o.NodeQosPolicyShaperClassMatchIPvsixTCP)
		if err != nil {
			return err
		}
	}

	return nil
}
