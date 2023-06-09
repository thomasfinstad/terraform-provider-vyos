// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemFlowAccountingNetflowServer describes the resource data model.
type SystemFlowAccountingNetflowServer struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafSystemFlowAccountingNetflowServerPort types.String `tfsdk:"port" json:"port,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFlowAccountingNetflowServer) GetVyosPath() []string {
	return []string{
		"system",
		"flow-accounting",
		"netflow",
		"server",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFlowAccountingNetflowServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `NetFlow destination server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 server to export NetFlow  |
|  ipv6  |  IPv6 server to export NetFlow  |

`,
		},

		// LeafNodes

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NetFlow port number

|  Format  |  Description  |
|----------|---------------|
|  u32:1025-65535  |  NetFlow port number  |

`,

			// Default:          stringdefault.StaticString(`2055`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemFlowAccountingNetflowServer) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafSystemFlowAccountingNetflowServerPort.IsNull() && !o.LeafSystemFlowAccountingNetflowServerPort.IsUnknown() {
		jsonData["port"] = o.LeafSystemFlowAccountingNetflowServerPort.ValueString()
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
func (o *SystemFlowAccountingNetflowServer) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["port"]; ok {
		o.LeafSystemFlowAccountingNetflowServerPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafSystemFlowAccountingNetflowServerPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
