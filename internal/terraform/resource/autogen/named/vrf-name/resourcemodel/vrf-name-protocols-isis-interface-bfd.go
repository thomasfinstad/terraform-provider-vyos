// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsIsisInterfaceBfd describes the resource data model.
type VrfNameProtocolsIsisInterfaceBfd struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceBfdProfile types.String `tfsdk:"profile" json:"profile,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceBfd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"profile": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BFD profile name  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisInterfaceBfd) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafVrfNameProtocolsIsisInterfaceBfdProfile.IsNull() && !o.LeafVrfNameProtocolsIsisInterfaceBfdProfile.IsUnknown() {
		jsonData["profile"] = o.LeafVrfNameProtocolsIsisInterfaceBfdProfile.ValueString()
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
func (o *VrfNameProtocolsIsisInterfaceBfd) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["profile"]; ok {
		o.LeafVrfNameProtocolsIsisInterfaceBfdProfile = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsIsisInterfaceBfdProfile = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
