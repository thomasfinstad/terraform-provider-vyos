// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList)
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
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf) UnmarshalJSON(jsonStr []byte) error {
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

		o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrfPrefixList)
		if err != nil {
			return err
		}
	}

	return nil
}
