// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisDefaultInformationOriginateIPvsix describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginateIPvsix struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne `tfsdk:"level_1" json:"level-1,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo *VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo `tfsdk:"level_2" json:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginateIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route into level-2

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvsix) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne)
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

	if !reflect.ValueOf(o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo)
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
func (o *VrfNameProtocolsIsisDefaultInformationOriginateIPvsix) UnmarshalJSON(jsonStr []byte) error {
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

		o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne = &VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelOne)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["level-2"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo = &VrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsixLevelTwo)
		if err != nil {
			return err
		}
	}

	return nil
}
