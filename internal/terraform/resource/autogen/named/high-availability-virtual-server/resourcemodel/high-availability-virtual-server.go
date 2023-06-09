// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// HighAvailabilityVirtualServer describes the resource data model.
type HighAvailabilityVirtualServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerAlgorithm          types.String `tfsdk:"algorithm" json:"algorithm,omitempty"`
	LeafHighAvailabilityVirtualServerDelayLoop          types.String `tfsdk:"delay_loop" json:"delay-loop,omitempty"`
	LeafHighAvailabilityVirtualServerForwardMethod      types.String `tfsdk:"forward_method" json:"forward-method,omitempty"`
	LeafHighAvailabilityVirtualServerFwmark             types.String `tfsdk:"fwmark" json:"fwmark,omitempty"`
	LeafHighAvailabilityVirtualServerPort               types.String `tfsdk:"port" json:"port,omitempty"`
	LeafHighAvailabilityVirtualServerPersistenceTimeout types.String `tfsdk:"persistence_timeout" json:"persistence-timeout,omitempty"`
	LeafHighAvailabilityVirtualServerProtocol           types.String `tfsdk:"protocol" json:"protocol,omitempty"`

	// TagNodes
	TagHighAvailabilityVirtualServerRealServer *map[string]HighAvailabilityVirtualServerRealServer `tfsdk:"real_server" json:"real-server,omitempty"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServer) GetVyosPath() []string {
	return []string{
		"high-availability",
		"virtual-server",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Load-balancing virtual server address

`,
		},

		// LeafNodes

		"algorithm": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Schedule algorithm (default - least-connection)

|  Format  |  Description  |
|----------|---------------|
|  round-robin  |  Round robin  |
|  weighted-round-robin  |  Weighted round robin  |
|  least-connection  |  Least connection  |
|  weighted-least-connection  |  Weighted least connection  |
|  source-hashing  |  Source hashing  |
|  destination-hashing  |  Destination hashing  |
|  locality-based-least-connection  |  Locality-Based least connection  |

`,

			// Default:          stringdefault.StaticString(`least-connection`),
			Computed: true,
		},

		"delay_loop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between health-checks (in seconds)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-600  |  Interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"forward_method": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding method

|  Format  |  Description  |
|----------|---------------|
|  direct  |  Direct routing  |
|  nat  |  NAT  |
|  tunnel  |  Tunneling  |

`,

			// Default:          stringdefault.StaticString(`nat`),
			Computed: true,
		},

		"fwmark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Match firewall mark value  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Numeric IP port  |

`,
		},

		"persistence_timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for persistent connections

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  Timeout for persistent connections  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol for port checks

|  Format  |  Description  |
|----------|---------------|
|  tcp  |  TCP  |
|  udp  |  UDP  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// TagNodes

		"real_server": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: HighAvailabilityVirtualServerRealServer{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Real server address

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *HighAvailabilityVirtualServer) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafHighAvailabilityVirtualServerAlgorithm.IsNull() && !o.LeafHighAvailabilityVirtualServerAlgorithm.IsUnknown() {
		jsonData["algorithm"] = o.LeafHighAvailabilityVirtualServerAlgorithm.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerDelayLoop.IsNull() && !o.LeafHighAvailabilityVirtualServerDelayLoop.IsUnknown() {
		jsonData["delay-loop"] = o.LeafHighAvailabilityVirtualServerDelayLoop.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerForwardMethod.IsNull() && !o.LeafHighAvailabilityVirtualServerForwardMethod.IsUnknown() {
		jsonData["forward-method"] = o.LeafHighAvailabilityVirtualServerForwardMethod.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerFwmark.IsNull() && !o.LeafHighAvailabilityVirtualServerFwmark.IsUnknown() {
		jsonData["fwmark"] = o.LeafHighAvailabilityVirtualServerFwmark.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerPort.IsNull() && !o.LeafHighAvailabilityVirtualServerPort.IsUnknown() {
		jsonData["port"] = o.LeafHighAvailabilityVirtualServerPort.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerPersistenceTimeout.IsNull() && !o.LeafHighAvailabilityVirtualServerPersistenceTimeout.IsUnknown() {
		jsonData["persistence-timeout"] = o.LeafHighAvailabilityVirtualServerPersistenceTimeout.ValueString()
	}

	if !o.LeafHighAvailabilityVirtualServerProtocol.IsNull() && !o.LeafHighAvailabilityVirtualServerProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafHighAvailabilityVirtualServerProtocol.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagHighAvailabilityVirtualServerRealServer).IsZero() {
		subJSONStr, err := json.Marshal(o.TagHighAvailabilityVirtualServerRealServer)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["real-server"] = subData
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *HighAvailabilityVirtualServer) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["algorithm"]; ok {
		o.LeafHighAvailabilityVirtualServerAlgorithm = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerAlgorithm = basetypes.NewStringNull()
	}

	if value, ok := jsonData["delay-loop"]; ok {
		o.LeafHighAvailabilityVirtualServerDelayLoop = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerDelayLoop = basetypes.NewStringNull()
	}

	if value, ok := jsonData["forward-method"]; ok {
		o.LeafHighAvailabilityVirtualServerForwardMethod = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerForwardMethod = basetypes.NewStringNull()
	}

	if value, ok := jsonData["fwmark"]; ok {
		o.LeafHighAvailabilityVirtualServerFwmark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerFwmark = basetypes.NewStringNull()
	}

	if value, ok := jsonData["port"]; ok {
		o.LeafHighAvailabilityVirtualServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerPort = basetypes.NewStringNull()
	}

	if value, ok := jsonData["persistence-timeout"]; ok {
		o.LeafHighAvailabilityVirtualServerPersistenceTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerPersistenceTimeout = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafHighAvailabilityVirtualServerProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafHighAvailabilityVirtualServerProtocol = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["real-server"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagHighAvailabilityVirtualServerRealServer = &map[string]HighAvailabilityVirtualServerRealServer{}

		err = json.Unmarshal(subJSONStr, o.TagHighAvailabilityVirtualServerRealServer)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
