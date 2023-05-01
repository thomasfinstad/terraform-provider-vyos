// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsecIkeGroupDeadPeerDetection describes the resource data model.
type VpnIPsecIkeGroupDeadPeerDetection struct {
	// LeafNodes
	LeafVpnIPsecIkeGroupDeadPeerDetectionAction   types.String `tfsdk:"action" json:"action,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionInterval types.String `tfsdk:"interval" json:"interval,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout  types.String `tfsdk:"timeout" json:"timeout,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecIkeGroupDeadPeerDetection) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive failure action

|  Format  |  Description  |
|----------|---------------|
|  hold  |  Attempt to re-negotiate the connection when matching traffic is seen  |
|  clear  |  Remove the connection immediately  |
|  restart  |  Attempt to re-negotiate the connection immediately  |

`,

			// Default:          stringdefault.StaticString(`clear`),
			Computed: true,
		},

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive interval

|  Format  |  Description  |
|----------|---------------|
|  u32:2-86400  |  Keep-alive interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"timeout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dead Peer Detection keep-alive timeout (IKEv1 only)

|  Format  |  Description  |
|----------|---------------|
|  u32:2-86400  |  Keep-alive timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecIkeGroupDeadPeerDetection) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVpnIPsecIkeGroupDeadPeerDetectionAction.IsNull() && !o.LeafVpnIPsecIkeGroupDeadPeerDetectionAction.IsUnknown() {
		jsonData["action"] = o.LeafVpnIPsecIkeGroupDeadPeerDetectionAction.ValueString()
	}

	if !o.LeafVpnIPsecIkeGroupDeadPeerDetectionInterval.IsNull() && !o.LeafVpnIPsecIkeGroupDeadPeerDetectionInterval.IsUnknown() {
		jsonData["interval"] = o.LeafVpnIPsecIkeGroupDeadPeerDetectionInterval.ValueString()
	}

	if !o.LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout.IsNull() && !o.LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout.IsUnknown() {
		jsonData["timeout"] = o.LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout.ValueString()
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
func (o *VpnIPsecIkeGroupDeadPeerDetection) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["interval"]; ok {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionInterval = basetypes.NewStringNull()
	}

	if value, ok := jsonData["timeout"]; ok {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
