// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceIPoeServerInterfaceExternalDhcp describes the resource data model.
type ServiceIPoeServerInterfaceExternalDhcp struct {
	// LeafNodes
	LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay types.String `tfsdk:"dhcp_relay" json:"dhcp-relay,omitempty"`
	LeafServiceIPoeServerInterfaceExternalDhcpGiaddr    types.String `tfsdk:"giaddr" json:"giaddr,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerInterfaceExternalDhcp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dhcp_relay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP Server the request will be redirected to.

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of the DHCP Server  |

`,
		},

		"giaddr": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Relay Agent IPv4 Address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Gateway IP address  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerInterfaceExternalDhcp) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay.IsNull() && !o.LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay.IsUnknown() {
		jsonData["dhcp-relay"] = o.LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay.ValueString()
	}

	if !o.LeafServiceIPoeServerInterfaceExternalDhcpGiaddr.IsNull() && !o.LeafServiceIPoeServerInterfaceExternalDhcpGiaddr.IsUnknown() {
		jsonData["giaddr"] = o.LeafServiceIPoeServerInterfaceExternalDhcpGiaddr.ValueString()
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
func (o *ServiceIPoeServerInterfaceExternalDhcp) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["dhcp-relay"]; ok {
		o.LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerInterfaceExternalDhcpDhcpRelay = basetypes.NewStringNull()
	}

	if value, ok := jsonData["giaddr"]; ok {
		o.LeafServiceIPoeServerInterfaceExternalDhcpGiaddr = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceIPoeServerInterfaceExternalDhcpGiaddr = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
