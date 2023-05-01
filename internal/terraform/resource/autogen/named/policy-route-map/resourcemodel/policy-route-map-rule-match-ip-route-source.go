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

// PolicyRouteMapRuleMatchIPRouteSource describes the resource data model.
type PolicyRouteMapRuleMatchIPRouteSource struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList types.String `tfsdk:"access_list"`
	LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList types.String `tfsdk:"prefix_list"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchIPRouteSource) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "route-source"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList.IsUnknown()) {
		vyosData["access-list"] = o.LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList.IsUnknown()) {
		vyosData["prefix-list"] = o.LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchIPRouteSource) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "route-source"}})

	// Leafs
	if value, ok := vyosData["access-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPRouteSourceAccessList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPRouteSourcePrefixList = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "route-source"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchIPRouteSource) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"access_list": types.StringType,
		"prefix_list": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPRouteSource) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		// TagNodes

		// Nodes

	}
}
