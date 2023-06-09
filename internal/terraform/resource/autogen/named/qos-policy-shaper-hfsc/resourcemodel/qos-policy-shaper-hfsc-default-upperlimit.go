// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyShaperHfscDefaultUpperlimit describes the resource data model.
type QosPolicyShaperHfscDefaultUpperlimit struct {
	// LeafNodes
	LeafQosPolicyShaperHfscDefaultUpperlimitD    types.String `tfsdk:"d" json:"d,omitempty"`
	LeafQosPolicyShaperHfscDefaultUpperlimitMone types.String `tfsdk:"m1" json:"m1,omitempty"`
	LeafQosPolicyShaperHfscDefaultUpperlimitMtwo types.String `tfsdk:"m2" json:"m2,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscDefaultUpperlimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"d": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service curve delay

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Time in milliseconds  |

`,
		},

		"m1": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m1 parameter for class traffic

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Rate in kbit (kilobit per second)  |
|  <number>%%  |  Percentage of overall rate  |
|  <number>bit  |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
|  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
|  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
|  <number>bps  |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		"m2": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m2 parameter for class traffic

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Rate in kbit (kilobit per second)  |
|  <number>%%  |  Percentage of overall rate  |
|  <number>bit  |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
|  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
|  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
|  <number>bps  |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperHfscDefaultUpperlimit) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyShaperHfscDefaultUpperlimitD.IsNull() && !o.LeafQosPolicyShaperHfscDefaultUpperlimitD.IsUnknown() {
		jsonData["d"] = o.LeafQosPolicyShaperHfscDefaultUpperlimitD.ValueString()
	}

	if !o.LeafQosPolicyShaperHfscDefaultUpperlimitMone.IsNull() && !o.LeafQosPolicyShaperHfscDefaultUpperlimitMone.IsUnknown() {
		jsonData["m1"] = o.LeafQosPolicyShaperHfscDefaultUpperlimitMone.ValueString()
	}

	if !o.LeafQosPolicyShaperHfscDefaultUpperlimitMtwo.IsNull() && !o.LeafQosPolicyShaperHfscDefaultUpperlimitMtwo.IsUnknown() {
		jsonData["m2"] = o.LeafQosPolicyShaperHfscDefaultUpperlimitMtwo.ValueString()
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
func (o *QosPolicyShaperHfscDefaultUpperlimit) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["d"]; ok {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitD = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitD = basetypes.NewStringNull()
	}

	if value, ok := jsonData["m1"]; ok {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitMone = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitMone = basetypes.NewStringNull()
	}

	if value, ok := jsonData["m2"]; ok {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitMtwo = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperHfscDefaultUpperlimitMtwo = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
