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

// FirewallIPvsixNameRuleHopLimit describes the resource data model.
type FirewallIPvsixNameRuleHopLimit struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleHopLimitEq types.String `tfsdk:"eq"`
	LeafFirewallIPvsixNameRuleHopLimitGt types.String `tfsdk:"gt"`
	LeafFirewallIPvsixNameRuleHopLimitLt types.String `tfsdk:"lt"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallIPvsixNameRuleHopLimit) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "hop-limit"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallIPvsixNameRuleHopLimitEq.IsNull() || o.LeafFirewallIPvsixNameRuleHopLimitEq.IsUnknown()) {
		vyosData["eq"] = o.LeafFirewallIPvsixNameRuleHopLimitEq.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleHopLimitGt.IsNull() || o.LeafFirewallIPvsixNameRuleHopLimitGt.IsUnknown()) {
		vyosData["gt"] = o.LeafFirewallIPvsixNameRuleHopLimitGt.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleHopLimitLt.IsNull() || o.LeafFirewallIPvsixNameRuleHopLimitLt.IsUnknown()) {
		vyosData["lt"] = o.LeafFirewallIPvsixNameRuleHopLimitLt.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallIPvsixNameRuleHopLimit) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "hop-limit"}})

	// Leafs
	if value, ok := vyosData["eq"]; ok {
		o.LeafFirewallIPvsixNameRuleHopLimitEq = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleHopLimitEq = basetypes.NewStringNull()
	}
	if value, ok := vyosData["gt"]; ok {
		o.LeafFirewallIPvsixNameRuleHopLimitGt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleHopLimitGt = basetypes.NewStringNull()
	}
	if value, ok := vyosData["lt"]; ok {
		o.LeafFirewallIPvsixNameRuleHopLimitLt = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleHopLimitLt = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "hop-limit"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallIPvsixNameRuleHopLimit) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"eq": types.StringType,
		"gt": types.StringType,
		"lt": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleHopLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"eq": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on equal value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Equal to value  |

`,
		},

		"gt": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on greater then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Greater then value  |

`,
		},

		"lt": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on less then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Less then value  |

`,
		},

		// TagNodes

		// Nodes

	}
}
