// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisRedistributeIPvsixBgp describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixBgp struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne *VrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne `tfsdk:"level_1" json:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo `tfsdk:"level_2" json:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisRedistributeIPvsixBgp) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["level-1"] = subData
	}

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["level-2"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisRedistributeIPvsixBgp) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["level-1"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne = &VrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelOne)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["level-2"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo = &VrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisRedistributeIPvsixBgpLevelTwo)
		if err != nil {
			return err
		}
	}

	return nil
}
