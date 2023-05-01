// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyRouteMapRuleMatchCommunity describes the resource data model.
type PolicyRouteMapRuleMatchCommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchCommunityCommunityList types.String `tfsdk:"community_list"`
	LeafPolicyRouteMapRuleMatchCommunityExactMatch    types.String `tfsdk:"exact_match"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchCommunity) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "community"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.IsNull() || o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.IsUnknown()) {
		vyosData["community-list"] = o.LeafPolicyRouteMapRuleMatchCommunityCommunityList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.IsNull() || o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.IsUnknown()) {
		vyosData["exact-match"] = o.LeafPolicyRouteMapRuleMatchCommunityExactMatch.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchCommunity) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "community"}})

	// Leafs
	if value, ok := vyosData["community-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchCommunityCommunityList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchCommunityCommunityList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["exact-match"]; ok {
		o.LeafPolicyRouteMapRuleMatchCommunityExactMatch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchCommunityExactMatch = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "community"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchCommunity) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"community_list": types.StringType,
		"exact_match":    types.StringType,

		// Tags

		// Nodes

	}
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