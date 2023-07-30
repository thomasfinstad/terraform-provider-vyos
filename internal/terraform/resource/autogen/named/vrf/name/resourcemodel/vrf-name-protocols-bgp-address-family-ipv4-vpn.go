// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourVpn struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork bool `tfsdk:"network" vyos:"network,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpn) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpn) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Nodes

	return nil
}
