// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyPriorityQueueClassMatchIPSource describes the resource data model.
type QosPolicyPriorityQueueClassMatchIPSource struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchIPSourceAddress types.String `tfsdk:"address" json:"address,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchIPSourcePort    types.String `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatchIPSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
|  ipv4net  |  IPv4 prefix  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyPriorityQueueClassMatchIPSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyPriorityQueueClassMatchIPSourceAddress.IsNull() && !o.LeafQosPolicyPriorityQueueClassMatchIPSourceAddress.IsUnknown() {
		jsonData["address"] = o.LeafQosPolicyPriorityQueueClassMatchIPSourceAddress.ValueString()
	}

	if !o.LeafQosPolicyPriorityQueueClassMatchIPSourcePort.IsNull() && !o.LeafQosPolicyPriorityQueueClassMatchIPSourcePort.IsUnknown() {
		jsonData["port"] = o.LeafQosPolicyPriorityQueueClassMatchIPSourcePort.ValueString()
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
func (o *QosPolicyPriorityQueueClassMatchIPSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPSourceAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPSourcePort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
