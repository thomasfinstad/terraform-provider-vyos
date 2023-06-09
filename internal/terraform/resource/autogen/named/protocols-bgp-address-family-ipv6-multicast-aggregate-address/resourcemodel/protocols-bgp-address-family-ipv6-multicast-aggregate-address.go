// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet       types.String `tfsdk:"as_set" json:"as-set,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap    types.String `tfsdk:"route_map" json:"route-map,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly types.String `tfsdk:"summary_only" json:"summary-only,omitempty"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetVyosPath() []string {
	return []string{
		"protocols",
		"bgp",
		"address-family",
		"ipv6-multicast",
		"aggregate-address",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP aggregate network/prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  BGP aggregate network/prefix  |

`,
		},

		// LeafNodes

		"as_set": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Generate AS-set path information for this aggregate address

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

		"summary_only": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Announce the aggregate summary network only

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet.IsUnknown() {
		jsonData["as-set"] = o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap.ValueString()
	}

	if !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly.IsNull() && !o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly.IsUnknown() {
		jsonData["summary-only"] = o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly.ValueString()
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
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["as-set"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap = basetypes.NewStringNull()
	}

	if value, ok := jsonData["summary-only"]; ok {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
