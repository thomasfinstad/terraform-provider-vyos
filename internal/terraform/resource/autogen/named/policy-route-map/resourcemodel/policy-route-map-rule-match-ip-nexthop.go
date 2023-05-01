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

// PolicyRouteMapRuleMatchIPNexthop describes the resource data model.
type PolicyRouteMapRuleMatchIPNexthop struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPNexthopAddress    types.String `tfsdk:"address"`
	LeafPolicyRouteMapRuleMatchIPNexthopAccessList types.String `tfsdk:"access_list"`
	LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen  types.String `tfsdk:"prefix_len"`
	LeafPolicyRouteMapRuleMatchIPNexthopPrefixList types.String `tfsdk:"prefix_list"`
	LeafPolicyRouteMapRuleMatchIPNexthopType       types.String `tfsdk:"type"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchIPNexthop) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "nexthop"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchIPNexthopAddress.IsNull() || o.LeafPolicyRouteMapRuleMatchIPNexthopAddress.IsUnknown()) {
		vyosData["address"] = o.LeafPolicyRouteMapRuleMatchIPNexthopAddress.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPNexthopAccessList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPNexthopAccessList.IsUnknown()) {
		vyosData["access-list"] = o.LeafPolicyRouteMapRuleMatchIPNexthopAccessList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen.IsNull() || o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen.IsUnknown()) {
		vyosData["prefix-len"] = o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixList.IsUnknown()) {
		vyosData["prefix-list"] = o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPNexthopType.IsNull() || o.LeafPolicyRouteMapRuleMatchIPNexthopType.IsUnknown()) {
		vyosData["type"] = o.LeafPolicyRouteMapRuleMatchIPNexthopType.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchIPNexthop) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "nexthop"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPNexthopAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPNexthopAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["access-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPNexthopAccessList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPNexthopAccessList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix-len"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixLen = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPNexthopPrefixList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["type"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPNexthopType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPNexthopType = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "nexthop"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchIPNexthop) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address":     types.StringType,
		"access_list": types.StringType,
		"prefix_len":  types.StringType,
		"prefix_list": types.StringType,
		"type":        types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchIPNexthop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address to match

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Nexthop IP address  |

`,
		},

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

		"prefix_len": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-length to match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  Prefix length  |

`,
		},

		"prefix_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP prefix-list to match

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
