// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PkiKeyPairPublic describes the resource data model.
type PkiKeyPairPublic struct {
	// LeafNodes
	LeafPkiKeyPairPublicKey types.String `tfsdk:"key" json:"key,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiKeyPairPublic) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Public key in PEM format

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiKeyPairPublic) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPkiKeyPairPublicKey.IsNull() && !o.LeafPkiKeyPairPublicKey.IsUnknown() {
		jsonData["key"] = o.LeafPkiKeyPairPublicKey.ValueString()
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
func (o *PkiKeyPairPublic) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["key"]; ok {
		o.LeafPkiKeyPairPublicKey = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPkiKeyPairPublicKey = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
