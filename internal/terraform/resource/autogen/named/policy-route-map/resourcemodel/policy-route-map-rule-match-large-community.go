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

// PolicyRouteMapRuleMatchLargeCommunity describes the resource data model.
type PolicyRouteMapRuleMatchLargeCommunity struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList types.String `tfsdk:"large_community_list"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchLargeCommunity) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "large-community"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList.IsNull() || o.LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList.IsUnknown()) {
		vyosData["large-community-list"] = o.LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchLargeCommunity) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "large-community"}})

	// Leafs
	if value, ok := vyosData["large-community-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchLargeCommunityLargeCommunityList = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "large-community"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchLargeCommunity) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"large_community_list": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchLargeCommunity) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"large_community_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP large-community-list to match

`,
		},

		// TagNodes

		// Nodes

	}
}