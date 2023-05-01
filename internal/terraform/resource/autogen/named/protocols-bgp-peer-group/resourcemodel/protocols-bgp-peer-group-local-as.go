// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsBgpPeerGroupLocalAs describes the resource data model.
type ProtocolsBgpPeerGroupLocalAs struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBgpPeerGroupLocalAsNoPrepend *ProtocolsBgpPeerGroupLocalAsNoPrepend `tfsdk:"no_prepend" json:"no-prepend,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupLocalAs) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"no_prepend": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupLocalAsNoPrepend{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable prepending local-as from/to updates for eBGP peers

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupLocalAs) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeProtocolsBgpPeerGroupLocalAsNoPrepend).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeProtocolsBgpPeerGroupLocalAsNoPrepend)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["no-prepend"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupLocalAs) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	// Tags

	// Nodes
	if value, ok := jsonData["no-prepend"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeProtocolsBgpPeerGroupLocalAsNoPrepend = &ProtocolsBgpPeerGroupLocalAsNoPrepend{}

		err = json.Unmarshal(subJSONStr, o.NodeProtocolsBgpPeerGroupLocalAsNoPrepend)
		if err != nil {
			return err
		}
	}

	return nil
}
