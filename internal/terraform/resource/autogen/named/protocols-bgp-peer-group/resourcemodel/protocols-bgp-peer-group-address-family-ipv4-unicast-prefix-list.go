// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList describes the resource data model.
type ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport types.String `tfsdk:"export" json:"export,omitempty"`
	LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport types.String `tfsdk:"import" json:"import,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport.IsUnknown() {
		jsonData["export"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport.ValueString()
	}

	if !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport.IsNull() && !o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport.IsUnknown() {
		jsonData["import"] = o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport.ValueString()
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
func (o *ProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["export"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListExport = basetypes.NewStringNull()
	}

	if value, ok := jsonData["import"]; ok {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastPrefixListImport = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
