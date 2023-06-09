// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpParametersBestpathPeerType describes the resource data model.
type VrfNameProtocolsBgpParametersBestpathPeerType struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax types.String `tfsdk:"multipath_relax" json:"multipath-relax,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersBestpathPeerType) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"multipath_relax": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Allow load sharing across routes learned from different peer types

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpParametersBestpathPeerType) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax.IsNull() && !o.LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax.IsUnknown() {
		jsonData["multipath-relax"] = o.LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax.ValueString()
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
func (o *VrfNameProtocolsBgpParametersBestpathPeerType) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["multipath-relax"]; ok {
		o.LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpParametersBestpathPeerTypeMultIPathRelax = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
