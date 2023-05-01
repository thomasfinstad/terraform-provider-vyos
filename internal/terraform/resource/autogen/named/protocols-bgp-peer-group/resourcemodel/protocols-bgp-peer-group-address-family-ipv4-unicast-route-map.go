// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport.ValueString()
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMap) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastRouteMapImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
