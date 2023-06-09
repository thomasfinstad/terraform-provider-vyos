// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive types.String `tfsdk:"receive" json:"receive,omitempty"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend    types.String `tfsdk:"send" json:"send,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive.IsUnknown() {
		jsonData["receive"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive.ValueString()
	}

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend.IsUnknown() {
		jsonData["send"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixList) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["receive"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListReceive = basetypes.NewStringNull()
	}

	if value, ok := jsonData["send"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourMulticastCapabilityOrfPrefixListSend = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
