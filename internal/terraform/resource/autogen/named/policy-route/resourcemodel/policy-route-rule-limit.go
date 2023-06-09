// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRuleLimit describes the resource data model.
type PolicyRouteRuleLimit struct {
	// LeafNodes
	LeafPolicyRouteRuleLimitBurst types.String `tfsdk:"burst" json:"burst,omitempty"`
	LeafPolicyRouteRuleLimitRate  types.String `tfsdk:"rate" json:"rate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of packets to allow in excess of rate

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Maximum number of packets to allow in excess of rate  |

`,
		},

		"rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum average matching rate

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Maximum average matching rate  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRuleLimit) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleLimitBurst.IsNull() && !o.LeafPolicyRouteRuleLimitBurst.IsUnknown() {
		jsonData["burst"] = o.LeafPolicyRouteRuleLimitBurst.ValueString()
	}

	if !o.LeafPolicyRouteRuleLimitRate.IsNull() && !o.LeafPolicyRouteRuleLimitRate.IsUnknown() {
		jsonData["rate"] = o.LeafPolicyRouteRuleLimitRate.ValueString()
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
func (o *PolicyRouteRuleLimit) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["burst"]; ok {
		o.LeafPolicyRouteRuleLimitBurst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleLimitBurst = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rate"]; ok {
		o.LeafPolicyRouteRuleLimitRate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleLimitRate = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
