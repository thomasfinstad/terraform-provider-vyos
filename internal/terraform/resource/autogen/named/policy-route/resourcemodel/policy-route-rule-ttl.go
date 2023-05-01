// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRuleTTL describes the resource data model.
type PolicyRouteRuleTTL struct {
	// LeafNodes
	LeafPolicyRouteRuleTTLEq types.String `tfsdk:"eq" json:"eq,omitempty"`
	LeafPolicyRouteRuleTTLGt types.String `tfsdk:"gt" json:"gt,omitempty"`
	LeafPolicyRouteRuleTTLLt types.String `tfsdk:"lt" json:"lt,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleTTL) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on equal value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Equal to value  |

`,
		},

		"gt": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on greater then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Greater then value  |

`,
		},

		"lt": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on less then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Less then value  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRuleTTL) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleTTLEq.IsNull() && !o.LeafPolicyRouteRuleTTLEq.IsUnknown() {
		jsonData["eq"] = o.LeafPolicyRouteRuleTTLEq.ValueString()
	}

	if !o.LeafPolicyRouteRuleTTLGt.IsNull() && !o.LeafPolicyRouteRuleTTLGt.IsUnknown() {
		jsonData["gt"] = o.LeafPolicyRouteRuleTTLGt.ValueString()
	}

	if !o.LeafPolicyRouteRuleTTLLt.IsNull() && !o.LeafPolicyRouteRuleTTLLt.IsUnknown() {
		jsonData["lt"] = o.LeafPolicyRouteRuleTTLLt.ValueString()
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
func (o *PolicyRouteRuleTTL) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["eq"]; ok {
		o.LeafPolicyRouteRuleTTLEq = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTTLEq = basetypes.NewStringNull()
	}

	if value, ok := jsonData["gt"]; ok {
		o.LeafPolicyRouteRuleTTLGt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTTLGt = basetypes.NewStringNull()
	}

	if value, ok := jsonData["lt"]; ok {
		o.LeafPolicyRouteRuleTTLLt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTTLLt = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
