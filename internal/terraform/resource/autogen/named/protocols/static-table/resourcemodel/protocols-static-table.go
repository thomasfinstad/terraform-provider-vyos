// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsStaticTable describes the resource data model.
type ProtocolsStaticTable struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsStaticTableDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsStaticTableRoute    bool `tfsdk:"route" vyos:"route,child"`
	ExistsTagProtocolsStaticTableRoutesix bool `tfsdk:"route6" vyos:"route6,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticTable) GetVyosPath() []string {
	return []string{
		"protocols",
		"static",
		"table",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticTable) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Policy route table number

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-200  |  Policy route table number  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Description  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsStaticTable) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafProtocolsStaticTableDescrIPtion.IsNull() && !o.LeafProtocolsStaticTableDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafProtocolsStaticTableDescrIPtion.ValueString()
	}

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsStaticTable) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafProtocolsStaticTableDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsStaticTableDescrIPtion = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
