// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborTimers describes the resource data model.
type VrfNameProtocolsBgpNeighborTimers struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborTimersConnect   types.String `tfsdk:"connect" json:"connect,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersHoldtime  types.String `tfsdk:"holdtime" json:"holdtime,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersKeepalive types.String `tfsdk:"keepalive" json:"keepalive,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborTimers) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP connect timer for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Connect timer in seconds  |
|  0  |  Disable connect timer  |

`,
		},

		"holdtime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP hold timer for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hold timer in seconds  |
|  0  |  Hold timer disabled  |

`,
		},

		"keepalive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Keepalive interval in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborTimers) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborTimersConnect.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborTimersConnect.IsUnknown() {
		jsonData["connect"] = o.LeafVrfNameProtocolsBgpNeighborTimersConnect.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborTimersHoldtime.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborTimersHoldtime.IsUnknown() {
		jsonData["holdtime"] = o.LeafVrfNameProtocolsBgpNeighborTimersHoldtime.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborTimersKeepalive.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborTimersKeepalive.IsUnknown() {
		jsonData["keepalive"] = o.LeafVrfNameProtocolsBgpNeighborTimersKeepalive.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborTimers) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["connect"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborTimersConnect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborTimersConnect = basetypes.NewStringNull()
	}

	if value, ok := jsonData["holdtime"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborTimersHoldtime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborTimersHoldtime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["keepalive"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborTimersKeepalive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborTimersKeepalive = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
