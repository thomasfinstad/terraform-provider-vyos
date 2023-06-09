// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteMapRuleMatchIPAddress describes the resource data model.
type PolicyRouteMapRuleMatchIPAddress struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPAddressAccessList types.String `tfsdk:"access_list" json:"access-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixList types.String `tfsdk:"prefix_list" json:"prefix-list,omitempty"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixLen  types.String `tfsdk:"prefix_len" json:"prefix-len,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"access_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP access-list to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-99  |  IP standard access list  |
|  u32:100-199  |  IP extended access list  |
|  u32:1300-1999  |  IP standard access list (expanded range)  |
|  u32:2000-2699  |  IP extended access list (expanded range)  |

`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-list to match

`,
		},

		"prefix_len": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-length to match (can be used for kernel routes only)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  Prefix length  |

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchIPAddress) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.IsUnknown() {
		jsonData["access-list"] = o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.IsUnknown() {
		jsonData["prefix-list"] = o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.IsNull() && !o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.IsUnknown() {
		jsonData["prefix-len"] = o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.ValueString()
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
func (o *PolicyRouteMapRuleMatchIPAddress) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["access-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressAccessList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressAccessList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["prefix-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["prefix-len"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
