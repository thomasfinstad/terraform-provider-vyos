// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyPrefixListRule describes the resource data model.
type PolicyPrefixListRule struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDPolicyPrefixList any `tfsdk:"prefix_list" vyos:"prefix-list,parent-id"`

	// LeafNodes
	LeafPolicyPrefixListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyPrefixListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyPrefixListRuleGe          types.String `tfsdk:"ge" vyos:"ge,omitempty"`
	LeafPolicyPrefixListRuleLe          types.String `tfsdk:"le" vyos:"le,omitempty"`
	LeafPolicyPrefixListRulePrefix      types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyPrefixListRule) GetVyosPath() []string {
	return []string{
		"policy",
		"prefix-list",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyPrefixListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule for this prefix-list

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Prefix-list rule number  |

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

		"ge": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask greater than or equal to it

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-32  |  Netmask greater than length  |

`,
		},

		"le": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length to match a netmask less than or equal to it

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-32  |  Netmask less than length  |

`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix to match

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  Prefix to match against  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyPrefixListRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyPrefixListRuleAction.IsNull() && !o.LeafPolicyPrefixListRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafPolicyPrefixListRuleAction.ValueString()
	}

	if !o.LeafPolicyPrefixListRuleDescrIPtion.IsNull() && !o.LeafPolicyPrefixListRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyPrefixListRuleDescrIPtion.ValueString()
	}

	if !o.LeafPolicyPrefixListRuleGe.IsNull() && !o.LeafPolicyPrefixListRuleGe.IsUnknown() {
		jsonData["ge"] = o.LeafPolicyPrefixListRuleGe.ValueString()
	}

	if !o.LeafPolicyPrefixListRuleLe.IsNull() && !o.LeafPolicyPrefixListRuleLe.IsUnknown() {
		jsonData["le"] = o.LeafPolicyPrefixListRuleLe.ValueString()
	}

	if !o.LeafPolicyPrefixListRulePrefix.IsNull() && !o.LeafPolicyPrefixListRulePrefix.IsUnknown() {
		jsonData["prefix"] = o.LeafPolicyPrefixListRulePrefix.ValueString()
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
func (o *PolicyPrefixListRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafPolicyPrefixListRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyPrefixListRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyPrefixListRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyPrefixListRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ge"]; ok {
		o.LeafPolicyPrefixListRuleGe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyPrefixListRuleGe = basetypes.NewStringNull()
	}

	if value, ok := jsonData["le"]; ok {
		o.LeafPolicyPrefixListRuleLe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyPrefixListRuleLe = basetypes.NewStringNull()
	}

	if value, ok := jsonData["prefix"]; ok {
		o.LeafPolicyPrefixListRulePrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyPrefixListRulePrefix = basetypes.NewStringNull()
	}

	// Nodes

	return nil
}
