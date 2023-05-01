// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallIPvsixNameRuleRecent describes the resource data model.
type FirewallIPvsixNameRuleRecent struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleRecentCount types.String `tfsdk:"count" json:"count,omitempty"`
	LeafFirewallIPvsixNameRuleRecentTime  types.String `tfsdk:"time" json:"time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleRecent) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"count": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen more than N times

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Source addresses seen more than N times  |

`,
		},

		"time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source addresses seen in the last second/minute/hour

|  Format  |  Description  |
|----------|---------------|
|  second  |  Source addresses seen COUNT times in the last second  |
|  minute  |  Source addresses seen COUNT times in the last minute  |
|  hour  |  Source addresses seen COUNT times in the last hour  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallIPvsixNameRuleRecent) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallIPvsixNameRuleRecentCount.IsNull() && !o.LeafFirewallIPvsixNameRuleRecentCount.IsUnknown() {
		jsonData["count"] = o.LeafFirewallIPvsixNameRuleRecentCount.ValueString()
	}

	if !o.LeafFirewallIPvsixNameRuleRecentTime.IsNull() && !o.LeafFirewallIPvsixNameRuleRecentTime.IsUnknown() {
		jsonData["time"] = o.LeafFirewallIPvsixNameRuleRecentTime.ValueString()
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
func (o *FirewallIPvsixNameRuleRecent) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["count"]; ok {
		o.LeafFirewallIPvsixNameRuleRecentCount = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleRecentCount = basetypes.NewStringNull()
	}

	if value, ok := jsonData["time"]; ok {
		o.LeafFirewallIPvsixNameRuleRecentTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleRecentTime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
