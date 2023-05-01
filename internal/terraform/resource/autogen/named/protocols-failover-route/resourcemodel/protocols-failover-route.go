// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsFailoverRoute describes the resource data model.
type ProtocolsFailoverRoute struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes
	TagProtocolsFailoverRouteNextHop *map[string]ProtocolsFailoverRouteNextHop `tfsdk:"next_hop" json:"next-hop,omitempty"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsFailoverRoute) GetVyosPath() []string {
	return []string{
		"protocols",
		"failover",
		"route",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Failover IPv4 route

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 failover route  |

`,
		},

		// LeafNodes

		// TagNodes

		"next_hop": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsFailoverRouteNextHop{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Next-hop IPv4 router address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Next-hop router address  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsFailoverRoute) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	if !reflect.ValueOf(o.TagProtocolsFailoverRouteNextHop).IsZero() {
		subJSONStr, err := json.Marshal(o.TagProtocolsFailoverRouteNextHop)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["next-hop"] = subData
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
func (o *ProtocolsFailoverRoute) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags
	if value, ok := jsonData["next-hop"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagProtocolsFailoverRouteNextHop = &map[string]ProtocolsFailoverRouteNextHop{}

		err = json.Unmarshal(subJSONStr, o.TagProtocolsFailoverRouteNextHop)
		if err != nil {
			return err
		}
	}

	// Nodes

	return nil
}
