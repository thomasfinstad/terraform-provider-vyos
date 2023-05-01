// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter outgoing route updates to this peer-group  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list to filter incoming route updates from this peer-group  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
