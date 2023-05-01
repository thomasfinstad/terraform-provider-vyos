// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive types.String `tfsdk:"receive" json:"receive,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend    types.String `tfsdk:"send" json:"send,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Capability to receive the ORF

`,
		},

		"send": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Capability to send the ORF

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.IsUnknown() {
		jsonData["receive"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.IsUnknown() {
		jsonData["send"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["receive"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListReceive = basetypes.NewStringNull()
	}

	if value, ok := jsonData["send"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixListSend = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
