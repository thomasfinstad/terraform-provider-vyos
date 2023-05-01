// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport.IsUnknown() {
		jsonData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMap) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnRouteMapImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
