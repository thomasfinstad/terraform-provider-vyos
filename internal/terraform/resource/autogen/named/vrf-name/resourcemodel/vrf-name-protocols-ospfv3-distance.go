// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsOspfvthreeDistance describes the resource data model.
type VrfNameProtocolsOspfvthreeDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeDistanceGlobal types.String `tfsdk:"global" json:"global,omitempty"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsOspfvthreeDistanceOspfvthree *VrfNameProtocolsOspfvthreeDistanceOspfvthree `tfsdk:"ospfv3" json:"ospfv3,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"global": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Administrative distance  |

`,
		},

		// TagNodes

		// Nodes

		"ospfv3": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDistanceOspfvthree{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `OSPFv3 administrative distance

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfvthreeDistance) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsOspfvthreeDistanceGlobal.IsNull() && !o.LeafVrfNameProtocolsOspfvthreeDistanceGlobal.IsUnknown() {
		jsonData["global"] = o.LeafVrfNameProtocolsOspfvthreeDistanceGlobal.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeVrfNameProtocolsOspfvthreeDistanceOspfvthree).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeVrfNameProtocolsOspfvthreeDistanceOspfvthree)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ospfv3"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfvthreeDistance) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["global"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeDistanceGlobal = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeDistanceGlobal = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["ospfv3"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeVrfNameProtocolsOspfvthreeDistanceOspfvthree = &VrfNameProtocolsOspfvthreeDistanceOspfvthree{}

		err = json.Unmarshal(subJSONStr, o.NodeVrfNameProtocolsOspfvthreeDistanceOspfvthree)
		if err != nil {
			return err
		}
	}

	return nil
}
