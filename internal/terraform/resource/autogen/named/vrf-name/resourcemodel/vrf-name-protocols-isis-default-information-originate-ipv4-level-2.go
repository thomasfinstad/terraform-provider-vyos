// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways   types.String `tfsdk:"always" json:"always,omitempty"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric   types.String `tfsdk:"metric" json:"metric,omitempty"`
	LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap types.String `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"always": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Always advertise default route

`,
		},

		"metric": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Default metric value  |

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
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways.IsNull() && !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways.IsUnknown() {
		jsonData["always"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways.ValueString()
	}

	if !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric.IsNull() && !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric.IsUnknown() {
		jsonData["metric"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric.ValueString()
	}

	if !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap.IsNull() && !o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap.ValueString()
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
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["always"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoAlways = basetypes.NewStringNull()
	}

	if value, ok := jsonData["metric"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoMetric = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisDefaultInformationOriginateIPvfourLevelTwoRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
