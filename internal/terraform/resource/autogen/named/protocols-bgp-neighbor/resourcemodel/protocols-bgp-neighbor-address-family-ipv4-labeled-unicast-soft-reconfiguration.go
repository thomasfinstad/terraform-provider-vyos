// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound types.String `tfsdk:"inbound" json:"inbound,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound.IsNull() && !o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound.IsUnknown() {
		jsonData["inbound"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound.ValueString()
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
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfiguration) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["inbound"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourLabeledUnicastSoftReconfigurationInbound = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
