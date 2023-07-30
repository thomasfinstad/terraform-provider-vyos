// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyExtcommunityListRule describes the resource data model.
type PolicyExtcommunityListRule struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDPolicyExtcommunityList any `tfsdk:"extcommunity_list" vyos:"extcommunity-list,parent-id"`

	// LeafNodes
	LeafPolicyExtcommunityListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyExtcommunityListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyExtcommunityListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyExtcommunityListRule) GetVyosPath() []string {
	return []string{
		"policy",
		"extcommunity-list",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyExtcommunityListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule for this BGP extended community list

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Extended community-list rule number  |

`,
		},

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
			MarkdownDescription: `Regular expression to match against an extended community list

    |  Format  |  Description  |
    |----------|---------------|
    |  <aa:nn:nn>  |  Extended community list regular expression  |
    |  <rt aa:nn:nn>  |  Route Target regular expression  |
    |  <soo aa:nn:nn>  |  Site of Origin regular expression  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyExtcommunityListRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyExtcommunityListRuleAction.IsNull() && !o.LeafPolicyExtcommunityListRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafPolicyExtcommunityListRuleAction.ValueString()
	}

	if !o.LeafPolicyExtcommunityListRuleDescrIPtion.IsNull() && !o.LeafPolicyExtcommunityListRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyExtcommunityListRuleDescrIPtion.ValueString()
	}

	if !o.LeafPolicyExtcommunityListRuleRegex.IsNull() && !o.LeafPolicyExtcommunityListRuleRegex.IsUnknown() {
		jsonData["regex"] = o.LeafPolicyExtcommunityListRuleRegex.ValueString()
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
func (o *PolicyExtcommunityListRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafPolicyExtcommunityListRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyExtcommunityListRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyExtcommunityListRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyExtcommunityListRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["regex"]; ok {
		o.LeafPolicyExtcommunityListRuleRegex = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyExtcommunityListRuleRegex = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
