// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBfdPeerSource describes the resource data model.
type ProtocolsBfdPeerSource struct {
	// LeafNodes
	LeafProtocolsBfdPeerSourceInterface types.String `tfsdk:"interface" json:"interface,omitempty"`
	LeafProtocolsBfdPeerSourceAddress   types.String `tfsdk:"address" json:"address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBfdPeerSource) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local address to bind our peer listener to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Local IPv4 address used to connect to the peer  |
|  ipv6  |  Local IPv6 address used to connect to the peer  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBfdPeerSource) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBfdPeerSourceInterface.IsNull() && !o.LeafProtocolsBfdPeerSourceInterface.IsUnknown() {
		jsonData["interface"] = o.LeafProtocolsBfdPeerSourceInterface.ValueString()
	}

	if !o.LeafProtocolsBfdPeerSourceAddress.IsNull() && !o.LeafProtocolsBfdPeerSourceAddress.IsUnknown() {
		jsonData["address"] = o.LeafProtocolsBfdPeerSourceAddress.ValueString()
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
func (o *ProtocolsBfdPeerSource) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["interface"]; ok {
		o.LeafProtocolsBfdPeerSourceInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerSourceInterface = basetypes.NewStringNull()
	}

	if value, ok := jsonData["address"]; ok {
		o.LeafProtocolsBfdPeerSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBfdPeerSourceAddress = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
