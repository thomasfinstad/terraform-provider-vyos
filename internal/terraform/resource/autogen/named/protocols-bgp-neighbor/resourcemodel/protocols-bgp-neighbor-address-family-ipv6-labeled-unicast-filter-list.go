// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastFilterListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
