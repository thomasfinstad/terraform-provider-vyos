// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpParametersGracefulRestart describes the resource data model.
type VrfNameProtocolsBgpParametersGracefulRestart struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime types.String `tfsdk:"stalepath_time" json:"stalepath-time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersGracefulRestart) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"stalepath_time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum time to hold onto restarting neighbors stale paths

|  Format  |  Description  |
|----------|---------------|
|  u32:1-3600  |  Hold time in seconds  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpParametersGracefulRestart) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime.IsNull() && !o.LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime.IsUnknown() {
		jsonData["stalepath-time"] = o.LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime.ValueString()
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
func (o *VrfNameProtocolsBgpParametersGracefulRestart) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["stalepath-time"]; ok {
		o.LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpParametersGracefulRestartStalepathTime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
