// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor types.String `tfsdk:"backdoor" json:"backdoor,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap types.String `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv6-labeled-unicast",
		"network",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Labeled Unicast IPv6 BGP network/prefix  |

`,
		},

		// LeafNodes

		"backdoor": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use BGP network/prefix as a backdoor route

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.IsUnknown() {
		jsonData["backdoor"] = o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap.ValueString()
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
func (o *ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["backdoor"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkBackdoor = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetworkRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
