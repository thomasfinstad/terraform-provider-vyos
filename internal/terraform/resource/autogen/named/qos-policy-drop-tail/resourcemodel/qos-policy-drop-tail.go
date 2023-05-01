// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyDropTail describes the resource data model.
type QosPolicyDropTail struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafQosPolicyDropTailDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafQosPolicyDropTailQueueLimit  types.String `tfsdk:"queue_limit" json:"queue-limit,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyDropTail) GetVyosPath() []string {
	return []string{
		"qos",
		"policy",
		"drop-tail",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyDropTail) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Packet limited First In, First Out queue

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"queue_limit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Queue size in packets  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyDropTail) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyDropTailDescrIPtion.IsNull() && !o.LeafQosPolicyDropTailDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafQosPolicyDropTailDescrIPtion.ValueString()
	}

	if !o.LeafQosPolicyDropTailQueueLimit.IsNull() && !o.LeafQosPolicyDropTailQueueLimit.IsUnknown() {
		jsonData["queue-limit"] = o.LeafQosPolicyDropTailQueueLimit.ValueString()
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
func (o *QosPolicyDropTail) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafQosPolicyDropTailDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyDropTailDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["queue-limit"]; ok {
		o.LeafQosPolicyDropTailQueueLimit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyDropTailQueueLimit = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
