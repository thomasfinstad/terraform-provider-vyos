// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsFailoverRouteNextHopCheck describes the resource data model.
type ProtocolsFailoverRouteNextHopCheck struct {
	// LeafNodes
	LeafProtocolsFailoverRouteNextHopCheckPort    types.String `tfsdk:"port" json:"port,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckTarget  types.String `tfsdk:"target" json:"target,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckTimeout types.String `tfsdk:"timeout" json:"timeout,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckType    types.String `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHopCheck) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,
		},

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Check target address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to check  |

`,
		},

		"timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Timeout between checks

|  Format  |  Description  |
|----------|---------------|
|  u32:1-300  |  Timeout in seconds between checks  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Check type

|  Format  |  Description  |
|----------|---------------|
|  arp  |  Check target by ARP  |
|  icmp  |  Check target by ICMP  |
|  tcp  |  Check target by TCP  |

`,

			// Default:          stringdefault.StaticString(`icmp`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsFailoverRouteNextHopCheck) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsFailoverRouteNextHopCheckPort.IsNull() && !o.LeafProtocolsFailoverRouteNextHopCheckPort.IsUnknown() {
		jsonData["port"] = o.LeafProtocolsFailoverRouteNextHopCheckPort.ValueString()
	}

	if !o.LeafProtocolsFailoverRouteNextHopCheckTarget.IsNull() && !o.LeafProtocolsFailoverRouteNextHopCheckTarget.IsUnknown() {
		jsonData["target"] = o.LeafProtocolsFailoverRouteNextHopCheckTarget.ValueString()
	}

	if !o.LeafProtocolsFailoverRouteNextHopCheckTimeout.IsNull() && !o.LeafProtocolsFailoverRouteNextHopCheckTimeout.IsUnknown() {
		jsonData["timeout"] = o.LeafProtocolsFailoverRouteNextHopCheckTimeout.ValueString()
	}

	if !o.LeafProtocolsFailoverRouteNextHopCheckType.IsNull() && !o.LeafProtocolsFailoverRouteNextHopCheckType.IsUnknown() {
		jsonData["type"] = o.LeafProtocolsFailoverRouteNextHopCheckType.ValueString()
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
func (o *ProtocolsFailoverRouteNextHopCheck) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["port"]; ok {
		o.LeafProtocolsFailoverRouteNextHopCheckPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopCheckPort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["target"]; ok {
		o.LeafProtocolsFailoverRouteNextHopCheckTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopCheckTarget = basetypes.NewStringNull()
	}

	if value, ok := jsonData["timeout"]; ok {
		o.LeafProtocolsFailoverRouteNextHopCheckTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopCheckTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["type"]; ok {
		o.LeafProtocolsFailoverRouteNextHopCheckType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsFailoverRouteNextHopCheckType = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
