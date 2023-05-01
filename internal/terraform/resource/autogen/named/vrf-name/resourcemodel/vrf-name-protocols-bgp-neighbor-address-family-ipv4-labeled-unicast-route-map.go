// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport.IsUnknown() {
		jsonData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMap) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastRouteMapImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
