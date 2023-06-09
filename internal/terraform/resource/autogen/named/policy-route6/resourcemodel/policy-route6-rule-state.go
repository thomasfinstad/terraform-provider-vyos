// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRoutesixRuleState describes the resource data model.
type PolicyRoutesixRuleState struct {
	// LeafNodes
	LeafPolicyRoutesixRuleStateEstablished types.String `tfsdk:"established" json:"established,omitempty"`
	LeafPolicyRoutesixRuleStateInvalID     types.String `tfsdk:"invalid" json:"invalid,omitempty"`
	LeafPolicyRoutesixRuleStateNew         types.String `tfsdk:"new" json:"new,omitempty"`
	LeafPolicyRoutesixRuleStateRelated     types.String `tfsdk:"related" json:"related,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleState) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"established": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Established state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"invalid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Invalid state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"new": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `New state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"related": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Related state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleState) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRoutesixRuleStateEstablished.IsNull() && !o.LeafPolicyRoutesixRuleStateEstablished.IsUnknown() {
		jsonData["established"] = o.LeafPolicyRoutesixRuleStateEstablished.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleStateInvalID.IsNull() && !o.LeafPolicyRoutesixRuleStateInvalID.IsUnknown() {
		jsonData["invalid"] = o.LeafPolicyRoutesixRuleStateInvalID.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleStateNew.IsNull() && !o.LeafPolicyRoutesixRuleStateNew.IsUnknown() {
		jsonData["new"] = o.LeafPolicyRoutesixRuleStateNew.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleStateRelated.IsNull() && !o.LeafPolicyRoutesixRuleStateRelated.IsUnknown() {
		jsonData["related"] = o.LeafPolicyRoutesixRuleStateRelated.ValueString()
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
func (o *PolicyRoutesixRuleState) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["established"]; ok {
		o.LeafPolicyRoutesixRuleStateEstablished = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleStateEstablished = basetypes.NewStringNull()
	}

	if value, ok := jsonData["invalid"]; ok {
		o.LeafPolicyRoutesixRuleStateInvalID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleStateInvalID = basetypes.NewStringNull()
	}

	if value, ok := jsonData["new"]; ok {
		o.LeafPolicyRoutesixRuleStateNew = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleStateNew = basetypes.NewStringNull()
	}

	if value, ok := jsonData["related"]; ok {
		o.LeafPolicyRoutesixRuleStateRelated = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleStateRelated = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
