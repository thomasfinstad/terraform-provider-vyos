// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisRedistributeIPvfourBgpLevelOne describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfourBgpLevelOne struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric   types.String `tfsdk:"metric" json:"metric,omitempty"`
	LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap types.String `tfsdk:"route_map" json:"route-map,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourBgpLevelOne) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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
func (o *VrfNameProtocolsIsisRedistributeIPvfourBgpLevelOne) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric.IsNull() && !o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric.IsUnknown() {
		jsonData["metric"] = o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric.ValueString()
	}

	if !o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap.IsNull() && !o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap.IsUnknown() {
		jsonData["route-map"] = o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap.ValueString()
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
func (o *VrfNameProtocolsIsisRedistributeIPvfourBgpLevelOne) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["metric"]; ok {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneMetric = basetypes.NewStringNull()
	}

	if value, ok := jsonData["route-map"]; ok {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisRedistributeIPvfourBgpLevelOneRouteMap = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
