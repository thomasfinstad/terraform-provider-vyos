// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport.IsUnknown() {
		jsonData["export"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport.ValueString()
	}

	if !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport.IsNull() && !o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport.IsUnknown() {
		jsonData["import"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport.ValueString()
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
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourMulticastPrefixListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
