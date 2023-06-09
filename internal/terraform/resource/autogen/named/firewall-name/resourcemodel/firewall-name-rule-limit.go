// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallNameRuleLimit describes the resource data model.
type FirewallNameRuleLimit struct {
	// LeafNodes
	LeafFirewallNameRuleLimitBurst types.String `tfsdk:"burst" json:"burst,omitempty"`
	LeafFirewallNameRuleLimitRate  types.String `tfsdk:"rate" json:"rate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
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
|  txt  |  integer/unit (Example: 5/minute)  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallNameRuleLimit) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallNameRuleLimitBurst.IsNull() && !o.LeafFirewallNameRuleLimitBurst.IsUnknown() {
		jsonData["burst"] = o.LeafFirewallNameRuleLimitBurst.ValueString()
	}

	if !o.LeafFirewallNameRuleLimitRate.IsNull() && !o.LeafFirewallNameRuleLimitRate.IsUnknown() {
		jsonData["rate"] = o.LeafFirewallNameRuleLimitRate.ValueString()
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
func (o *FirewallNameRuleLimit) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["burst"]; ok {
		o.LeafFirewallNameRuleLimitBurst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleLimitBurst = basetypes.NewStringNull()
	}

	if value, ok := jsonData["rate"]; ok {
		o.LeafFirewallNameRuleLimitRate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleLimitRate = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
