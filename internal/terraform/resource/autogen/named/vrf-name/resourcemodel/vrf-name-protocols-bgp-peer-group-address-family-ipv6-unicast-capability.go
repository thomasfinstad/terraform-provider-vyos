// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf `tfsdk:"orf" json:"orf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"orf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise ORF capability to this peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["orf"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapability) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["orf"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixUnicastCapabilityOrf)
		if err != nil {
			return err
		}
	}

	return nil
}
