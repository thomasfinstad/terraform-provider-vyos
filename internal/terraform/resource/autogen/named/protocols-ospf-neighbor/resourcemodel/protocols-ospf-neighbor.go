// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsOspfNeighbor describes the resource data model.
type ProtocolsOspfNeighbor struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsOspfNeighborPollInterval types.String `tfsdk:"poll_interval" json:"poll-interval,omitempty"`
	LeafProtocolsOspfNeighborPriority     types.String `tfsdk:"priority" json:"priority,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfNeighbor) GetVyosPath() []string {
	return []string{
		"protocols",
		"ospf",
		"neighbor",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfNeighbor) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specify neighbor router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Neighbor IP address  |

`,
		},

		// LeafNodes

		"poll_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dead neighbor polling interval

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Seconds between dead neighbor polling interval  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"priority": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor priority in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Neighbor priority  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsOspfNeighbor) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsOspfNeighborPollInterval.IsNull() && !o.LeafProtocolsOspfNeighborPollInterval.IsUnknown() {
		jsonData["poll-interval"] = o.LeafProtocolsOspfNeighborPollInterval.ValueString()
	}

	if !o.LeafProtocolsOspfNeighborPriority.IsNull() && !o.LeafProtocolsOspfNeighborPriority.IsUnknown() {
		jsonData["priority"] = o.LeafProtocolsOspfNeighborPriority.ValueString()
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
func (o *ProtocolsOspfNeighbor) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["poll-interval"]; ok {
		o.LeafProtocolsOspfNeighborPollInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfNeighborPollInterval = basetypes.NewStringNull()
	}

	if value, ok := jsonData["priority"]; ok {
		o.LeafProtocolsOspfNeighborPriority = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsOspfNeighborPriority = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
