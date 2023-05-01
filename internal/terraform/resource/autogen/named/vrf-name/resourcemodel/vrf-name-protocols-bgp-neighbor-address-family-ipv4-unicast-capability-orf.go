// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["prefix-list"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrf) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["prefix-list"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapabilityOrfPrefixList)
		if err != nil {
			return err
		}
	}

	return nil
}
