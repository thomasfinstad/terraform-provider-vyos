// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteMapRuleMatchIPvsixNexthop describes the resource data model.
type PolicyRouteMapRuleMatchIPvsixNexthop struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress    types.String `tfsdk:"address" json:"address,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList types.String `tfsdk:"access_list" json:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList types.String `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPvsixNexthopType       types.String `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPvsixNexthop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 address of next-hop

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Nexthop IPv6 address  |

`,
		},

		"access_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 access-list to match

|  Format  |  Description  |
|----------|---------------|
|  txt  |  IPV6 access list name  |

`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 prefix-list to match

`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match type

|  Format  |  Description  |
|----------|---------------|
|  blackhole  |  Blackhole  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchIPvsixNexthop) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress.IsUnknown() {
		jsonData["address"] = o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList.IsUnknown() {
		jsonData["access-list"] = o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList.IsUnknown() {
		jsonData["prefix-list"] = o.LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopType.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPvsixNexthopType.IsUnknown() {
		jsonData["type"] = o.LeafPolicyRouteMapRuleMatchIPvsixNexthopType.ValueString()
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
func (o *PolicyRouteMapRuleMatchIPvsixNexthop) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["address"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAddress = basetypes.NewStringNull()
	}

	if value, ok := jsonData["access-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopAccessList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["prefix-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopPrefixList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["type"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPvsixNexthopType = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
