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

// FirewallNameRuleRecent describes the resource data model.
type FirewallNameRuleRecent struct {
	// LeafNodes
	LeafFirewallNameRuleRecentCount types.String `tfsdk:"count"`
	LeafFirewallNameRuleRecentTime  types.String `tfsdk:"time"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallNameRuleRecent) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "recent"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallNameRuleRecentCount.IsNull() || o.LeafFirewallNameRuleRecentCount.IsUnknown()) {
		vyosData["count"] = o.LeafFirewallNameRuleRecentCount.ValueString()
	}
	if !(o.LeafFirewallNameRuleRecentTime.IsNull() || o.LeafFirewallNameRuleRecentTime.IsUnknown()) {
		vyosData["time"] = o.LeafFirewallNameRuleRecentTime.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallNameRuleRecent) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "recent"}})

	// Leafs
	if value, ok := vyosData["count"]; ok {
		o.LeafFirewallNameRuleRecentCount = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleRecentCount = basetypes.NewStringNull()
	}
	if value, ok := vyosData["time"]; ok {
		o.LeafFirewallNameRuleRecentTime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleRecentTime = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "recent"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallNameRuleRecent) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"count": types.StringType,
		"time":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleRecent) ResourceSchemaAttributes() map[string]schema.Attribute {
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
