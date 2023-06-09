// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyAccessListsixRule describes the resource data model.
type PolicyAccessListsixRule struct {
	// LeafNodes
	LeafPolicyAccessListsixRuleAction      types.String `tfsdk:"action" json:"action,omitempty"`
	LeafPolicyAccessListsixRuleDescrIPtion types.String `tfsdk:"description" json:"description,omitempty"`

	// TagNodes

	// Nodes
	NodePolicyAccessListsixRuleSource *PolicyAccessListsixRuleSource `tfsdk:"source" json:"source,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListsixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		// TagNodes

		// Nodes

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyAccessListsixRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source IPv6 network to match

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyAccessListsixRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyAccessListsixRuleAction.IsNull() && !o.LeafPolicyAccessListsixRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafPolicyAccessListsixRuleAction.ValueString()
	}

	if !o.LeafPolicyAccessListsixRuleDescrIPtion.IsNull() && !o.LeafPolicyAccessListsixRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyAccessListsixRuleDescrIPtion.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyAccessListsixRuleSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyAccessListsixRuleSource)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["source"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyAccessListsixRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafPolicyAccessListsixRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListsixRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyAccessListsixRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAccessListsixRuleDescrIPtion = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyAccessListsixRuleSource = &PolicyAccessListsixRuleSource{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyAccessListsixRuleSource)
		if err != nil {
			return err
		}
	}

	return nil
}
