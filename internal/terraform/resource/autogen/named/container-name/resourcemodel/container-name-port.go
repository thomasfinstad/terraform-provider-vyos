// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ContainerNamePort describes the resource data model.
type ContainerNamePort struct {
	// LeafNodes
	LeafContainerNamePortSource      types.String `tfsdk:"source" json:"source,omitempty"`
	LeafContainerNamePortDestination types.String `tfsdk:"destination" json:"destination,omitempty"`
	LeafContainerNamePortProtocol    types.String `tfsdk:"protocol" json:"protocol,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNamePort) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Source host port  |
|  start-end  |  Source host port range (e.g. 10025-10030)  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Destination container port  |
|  start-end  |  Destination container port range (e.g. 10025-10030)  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport protocol used for port mapping

|  Format  |  Description  |
|----------|---------------|
|  tcp  |  Use Transmission Control Protocol for given port  |
|  udp  |  Use User Datagram Protocol for given port  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNamePort) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafContainerNamePortSource.IsNull() && !o.LeafContainerNamePortSource.IsUnknown() {
		jsonData["source"] = o.LeafContainerNamePortSource.ValueString()
	}

	if !o.LeafContainerNamePortDestination.IsNull() && !o.LeafContainerNamePortDestination.IsUnknown() {
		jsonData["destination"] = o.LeafContainerNamePortDestination.ValueString()
	}

	if !o.LeafContainerNamePortProtocol.IsNull() && !o.LeafContainerNamePortProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafContainerNamePortProtocol.ValueString()
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
func (o *ContainerNamePort) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["source"]; ok {
		o.LeafContainerNamePortSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNamePortSource = basetypes.NewStringNull()
	}

	if value, ok := jsonData["destination"]; ok {
		o.LeafContainerNamePortDestination = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNamePortDestination = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafContainerNamePortProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafContainerNamePortProtocol = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
