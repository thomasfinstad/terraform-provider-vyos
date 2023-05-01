// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PkiKeyPair describes the resource data model.
type PkiKeyPair struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodePkiKeyPairPublic  *PkiKeyPairPublic  `tfsdk:"public" json:"public,omitempty"`
	NodePkiKeyPairPrivate *PkiKeyPairPrivate `tfsdk:"private" json:"private,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiKeyPair) GetVyosPath() []string {
	return []string{
		"pki",
		"key-pair",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPair) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Public and private keys

`,
		},

		// LeafNodes

		// TagNodes

		// Nodes

		"public": schema.SingleNestedAttribute{
			Attributes: PkiKeyPairPublic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Public key

`,
		},

		"private": schema.SingleNestedAttribute{
			Attributes: PkiKeyPairPrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Private key

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiKeyPair) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePkiKeyPairPublic).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePkiKeyPairPublic)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["public"] = subData
	}

	if !reflect.ValueOf(o.NodePkiKeyPairPrivate).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePkiKeyPairPrivate)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["private"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PkiKeyPair) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["public"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePkiKeyPairPublic = &PkiKeyPairPublic{}

		err = json.Unmarshal(subJSONStr, o.NodePkiKeyPairPublic)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["private"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePkiKeyPairPrivate = &PkiKeyPairPrivate{}

		err = json.Unmarshal(subJSONStr, o.NodePkiKeyPairPrivate)
		if err != nil {
			return err
		}
	}

	return nil
}
