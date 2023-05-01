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

// PolicyRouteRuleLimit describes the resource data model.
type PolicyRouteRuleLimit struct {
	// LeafNodes
	LeafPolicyRouteRuleLimitBurst types.String `tfsdk:"burst"`
	LeafPolicyRouteRuleLimitRate  types.String `tfsdk:"rate"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleLimit) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "limit"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleLimitBurst.IsNull() || o.LeafPolicyRouteRuleLimitBurst.IsUnknown()) {
		vyosData["burst"] = o.LeafPolicyRouteRuleLimitBurst.ValueString()
	}
	if !(o.LeafPolicyRouteRuleLimitRate.IsNull() || o.LeafPolicyRouteRuleLimitRate.IsUnknown()) {
		vyosData["rate"] = o.LeafPolicyRouteRuleLimitRate.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleLimit) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "limit"}})

	// Leafs
	if value, ok := vyosData["burst"]; ok {
		o.LeafPolicyRouteRuleLimitBurst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleLimitBurst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rate"]; ok {
		o.LeafPolicyRouteRuleLimitRate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleLimitRate = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "limit"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleLimit) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"burst": types.StringType,
		"rate":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of packets to allow in excess of rate

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Maximum number of packets to allow in excess of rate  |

`,
		},

		"rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum average matching rate

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Maximum average matching rate  |

`,
		},

		// TagNodes

		// Nodes

	}
}
