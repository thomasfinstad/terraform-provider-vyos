// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisInterfaceNetwork describes the resource data model.
type VrfNameProtocolsIsisInterfaceNetwork struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint types.String `tfsdk:"point_to_point" json:"point-to-point,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceNetwork) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"point_to_point": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `point-to-point network type

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisInterfaceNetwork) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint.IsNull() && !o.LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint.IsUnknown() {
		jsonData["point-to-point"] = o.LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint.ValueString()
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
func (o *VrfNameProtocolsIsisInterfaceNetwork) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["point-to-point"]; ok {
		o.LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisInterfaceNetworkPointToPoint = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
