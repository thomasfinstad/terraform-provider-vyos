// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteMapRuleMatchCommunity describes the resource data model.
type PolicyRouteMapRuleMatchCommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchCommunityCommunityList types.String `tfsdk:"community_list" json:"community-list,omitempty"`
	LeafPolicyRouteMapRuleMatchCommunityExactMatch    types.String `tfsdk:"exact_match" json:"exact-match,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"community_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP community-list to match

`,
		},

		"exact_match": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Community-list to exactly match

`,
		},

		// TagNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteMapRuleMatchCommunity) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.IsNull() && !o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.IsUnknown() {
		jsonData["community-list"] = o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.ValueString()
	}

	if !o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.IsNull() && !o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.IsUnknown() {
		jsonData["exact-match"] = o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.ValueString()
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
func (o *PolicyRouteMapRuleMatchCommunity) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["community-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchCommunityCommunityList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchCommunityCommunityList = basetypes.NewStringNull()
	}

	if value, ok := jsonData["exact-match"]; ok {
		o.LeafPolicyRouteMapRuleMatchCommunityExactMatch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchCommunityExactMatch = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
