// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap types.String `tfsdk:"advertise_map" json:"advertise-map,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap     types.String `tfsdk:"exist_map" json:"exist-map,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap  types.String `tfsdk:"non_exist_map" json:"non-exist-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to conditionally advertise routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"non_exist_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap.IsUnknown() {
		jsonData["advertise-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap.IsUnknown() {
		jsonData["exist-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap.IsUnknown() {
		jsonData["non-exist-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["advertise-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseAdvertiseMap = basetypes.NewStringNull()
	}

	if value, ok := jsonData["exist-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseExistMap = basetypes.NewStringNull()
	}

	if value, ok := jsonData["non-exist-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertiseNonExistMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
