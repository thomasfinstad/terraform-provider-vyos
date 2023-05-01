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

// PolicyRouteMapRuleMatchIPAddress describes the resource data model.
type PolicyRouteMapRuleMatchIPAddress struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchIPAddressAccessList types.String `tfsdk:"access_list"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixList types.String `tfsdk:"prefix_list"`
	LeafPolicyRouteMapRuleMatchIPAddressPrefixLen  types.String `tfsdk:"prefix_len"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteMapRuleMatchIPAddress) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "address"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.IsUnknown()) {
		vyosData["access-list"] = o.LeafPolicyRouteMapRuleMatchIPAddressAccessList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.IsNull() || o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.IsUnknown()) {
		vyosData["prefix-list"] = o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList.ValueString()
	}
	if !(o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.IsNull() || o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.IsUnknown()) {
		vyosData["prefix-len"] = o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteMapRuleMatchIPAddress) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "address"}})

	// Leafs
	if value, ok := vyosData["access-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressAccessList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressAccessList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix-list"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["prefix-len"]; ok {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteMapRuleMatchIPAddressPrefixLen = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route-map", "rule", "match", "ip", "address"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteMapRuleMatchIPAddress) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"access_list": types.StringType,
		"prefix_list": types.StringType,
		"prefix_len":  types.StringType,

		// Tags

		// Nodes

	}
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