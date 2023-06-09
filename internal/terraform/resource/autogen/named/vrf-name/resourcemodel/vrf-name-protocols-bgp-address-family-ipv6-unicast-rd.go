// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn `tfsdk:"vpn" json:"vpn,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Between current address-family and VPN

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["vpn"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRd) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["vpn"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn = &VrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastRdVpn)
		if err != nil {
			return err
		}
	}

	return nil
}
