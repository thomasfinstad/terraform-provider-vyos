// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyShaperClassMatchIPTCP describes the resource data model.
type QosPolicyShaperClassMatchIPTCP struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPTCPAck types.String `tfsdk:"ack" json:"ack,omitempty"`
	LeafQosPolicyShaperClassMatchIPTCPSyn types.String `tfsdk:"syn" json:"syn,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIPTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ack": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP ACK

`,
		},

		"syn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP SYN

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperClassMatchIPTCP) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyShaperClassMatchIPTCPAck.IsNull() && !o.LeafQosPolicyShaperClassMatchIPTCPAck.IsUnknown() {
		jsonData["ack"] = o.LeafQosPolicyShaperClassMatchIPTCPAck.ValueString()
	}

	if !o.LeafQosPolicyShaperClassMatchIPTCPSyn.IsNull() && !o.LeafQosPolicyShaperClassMatchIPTCPSyn.IsUnknown() {
		jsonData["syn"] = o.LeafQosPolicyShaperClassMatchIPTCPSyn.ValueString()
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
func (o *QosPolicyShaperClassMatchIPTCP) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["ack"]; ok {
		o.LeafQosPolicyShaperClassMatchIPTCPAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPTCPAck = basetypes.NewStringNull()
	}

	if value, ok := jsonData["syn"]; ok {
		o.LeafQosPolicyShaperClassMatchIPTCPSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchIPTCPSyn = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
