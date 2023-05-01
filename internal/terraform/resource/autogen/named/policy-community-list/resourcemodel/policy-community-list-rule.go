// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyCommunityListRule describes the resource data model.
type PolicyCommunityListRule struct {
	// LeafNodes
	LeafPolicyCommunityListRuleAction      types.String `tfsdk:"action" json:"action,omitempty"`
	LeafPolicyCommunityListRuleDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`
	LeafPolicyCommunityListRuleRegex       types.String `tfsdk:"regex" json:"regex,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyCommunityListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

|  Format  |  Description  |
|----------|---------------|
|  permit  |  Permit matching entries  |
|  deny  |  Deny matching entries  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"regex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against a community-list

|  Format  |  Description  |
|----------|---------------|
|  <aa:nn>  |  Community number in AA:NN format  |
|  local-AS  |  Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03  |
|  no-advertise  |  Well-known communities value NO_ADVERTISE 0xFFFFFF02  |
|  no-export  |  Well-known communities value NO_EXPORT 0xFFFFFF01  |
|  internet  |  Well-known communities value 0  |
|  additive  |  New value is appended to the existing value  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyCommunityListRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyCommunityListRuleAction.IsNull() && !o.LeafPolicyCommunityListRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafPolicyCommunityListRuleAction.ValueString()
	}

	if !o.LeafPolicyCommunityListRuleDescrIPtion.IsNull() && !o.LeafPolicyCommunityListRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyCommunityListRuleDescrIPtion.ValueString()
	}

	if !o.LeafPolicyCommunityListRuleRegex.IsNull() && !o.LeafPolicyCommunityListRuleRegex.IsUnknown() {
		jsonData["regex"] = o.LeafPolicyCommunityListRuleRegex.ValueString()
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
func (o *PolicyCommunityListRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafPolicyCommunityListRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyCommunityListRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyCommunityListRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyCommunityListRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["regex"]; ok {
		o.LeafPolicyCommunityListRuleRegex = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyCommunityListRuleRegex = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
