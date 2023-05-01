// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixUnicastPrefixListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
