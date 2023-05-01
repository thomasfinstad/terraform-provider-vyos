// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMap describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMap struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter outgoing route updates

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to filter incoming route updates

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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMap) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport.IsUnknown() {
		jsonData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMap) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastRouteMapImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
