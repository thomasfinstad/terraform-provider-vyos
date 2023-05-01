// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv4 prefix-list  |

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv4 prefix-list  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastPrefixListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
