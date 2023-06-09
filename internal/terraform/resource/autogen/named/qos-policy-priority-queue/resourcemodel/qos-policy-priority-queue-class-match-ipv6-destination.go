// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyPriorityQueueClassMatchIPvsixDestination describes the resource data model.
type QosPolicyPriorityQueueClassMatchIPvsixDestination struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress types.String `tfsdk:"address" json:"address,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort    types.String `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatchIPvsixDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

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
func (o *QosPolicyPriorityQueueClassMatchIPvsixDestination) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress.IsNull() && !o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress.IsUnknown() {
		jsonData["address"] = o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress.ValueString()
	}

	if !o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort.IsNull() && !o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort.IsUnknown() {
		jsonData["port"] = o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort.ValueString()
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
func (o *QosPolicyPriorityQueueClassMatchIPvsixDestination) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchIPvsixDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
