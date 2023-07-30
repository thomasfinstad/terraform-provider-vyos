// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsStaticMulticastRouteNextHop describes the resource data model.
type ProtocolsStaticMulticastRouteNextHop struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDProtocolsStaticMulticastRoute any `tfsdk:"route" vyos:"route,parent-id"`

	// LeafNodes
	LeafProtocolsStaticMulticastRouteNextHopDistance types.String `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticMulticastRouteNextHop) GetVyosPath() []string {
	return []string{
		"protocols",
		"static",
		"multicast",
		"route",
		"next-hop",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticMulticastRouteNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Nexthop IPv4 address

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  Nexthop IPv4 address  |

`,
		},

		// LeafNodes

		"distance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Distance value for this route

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Distance for this route  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticMulticastRouteNextHop) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsStaticMulticastRouteNextHopDistance.IsNull() && !o.LeafProtocolsStaticMulticastRouteNextHopDistance.IsUnknown() {
		jsonData["distance"] = o.LeafProtocolsStaticMulticastRouteNextHopDistance.ValueString()
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
func (o *ProtocolsStaticMulticastRouteNextHop) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["distance"]; ok {
		o.LeafProtocolsStaticMulticastRouteNextHopDistance = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticMulticastRouteNextHopDistance = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
