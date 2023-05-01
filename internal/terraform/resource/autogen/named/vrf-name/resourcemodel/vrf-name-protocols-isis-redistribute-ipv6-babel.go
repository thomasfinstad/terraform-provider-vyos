// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisRedistributeIPvsixBabel describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsixBabel struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne *VrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne `tfsdk:"level_1" json:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo *VrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo `tfsdk:"level_2" json:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsixBabel) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisRedistributeIPvsixBabel) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne)
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

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo)
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
func (o *VrfNameProtocolsIsisRedistributeIPvsixBabel) UnmarshalJSON(jsonStr []byte) error {
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

		o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne = &VrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelOne)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["level-2"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo = &VrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisRedistributeIPvsixBabelLevelTwo)
		if err != nil {
			return err
		}
	}

	return nil
}
