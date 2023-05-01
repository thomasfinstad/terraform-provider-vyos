// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// QosPolicyLimiterClassMatchIPDestination describes the resource data model.
type QosPolicyLimiterClassMatchIPDestination struct {
	// LeafNodes
	LeafQosPolicyLimiterClassMatchIPDestinationAddress types.String `tfsdk:"address" json:"address,omitempty"`
	LeafQosPolicyLimiterClassMatchIPDestinationPort    types.String `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClassMatchIPDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *QosPolicyLimiterClassMatchIPDestination) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafQosPolicyLimiterClassMatchIPDestinationAddress.IsNull() && !o.LeafQosPolicyLimiterClassMatchIPDestinationAddress.IsUnknown() {
		jsonData["address"] = o.LeafQosPolicyLimiterClassMatchIPDestinationAddress.ValueString()
	}

	if !o.LeafQosPolicyLimiterClassMatchIPDestinationPort.IsNull() && !o.LeafQosPolicyLimiterClassMatchIPDestinationPort.IsUnknown() {
		jsonData["port"] = o.LeafQosPolicyLimiterClassMatchIPDestinationPort.ValueString()
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
func (o *QosPolicyLimiterClassMatchIPDestination) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafQosPolicyLimiterClassMatchIPDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyLimiterClassMatchIPDestinationAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafQosPolicyLimiterClassMatchIPDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyLimiterClassMatchIPDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
