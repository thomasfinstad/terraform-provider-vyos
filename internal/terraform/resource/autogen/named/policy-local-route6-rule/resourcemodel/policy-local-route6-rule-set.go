// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyLocalRoutesixRuleSet describes the resource data model.
type PolicyLocalRoutesixRuleSet struct {
	// LeafNodes
	LeafPolicyLocalRoutesixRuleSetTable types.String `tfsdk:"table" json:"table,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRuleSet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Routing table to forward packet with

|  Format  |  Description  |
|----------|---------------|
|  u32:1-200  |  Table number  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyLocalRoutesixRuleSet) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyLocalRoutesixRuleSetTable.IsNull() && !o.LeafPolicyLocalRoutesixRuleSetTable.IsUnknown() {
		jsonData["table"] = o.LeafPolicyLocalRoutesixRuleSetTable.ValueString()
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
func (o *PolicyLocalRoutesixRuleSet) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["table"]; ok {
		o.LeafPolicyLocalRoutesixRuleSetTable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleSetTable = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
