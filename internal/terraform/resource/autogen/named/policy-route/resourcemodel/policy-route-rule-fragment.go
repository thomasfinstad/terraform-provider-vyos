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

// PolicyRouteRuleFragment describes the resource data model.
type PolicyRouteRuleFragment struct {
	// LeafNodes
	LeafPolicyRouteRuleFragmentMatchFrag    types.String `tfsdk:"match_frag"`
	LeafPolicyRouteRuleFragmentMatchNonFrag types.String `tfsdk:"match_non_frag"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleFragment) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "fragment"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleFragmentMatchFrag.IsNull() || o.LeafPolicyRouteRuleFragmentMatchFrag.IsUnknown()) {
		vyosData["match-frag"] = o.LeafPolicyRouteRuleFragmentMatchFrag.ValueString()
	}
	if !(o.LeafPolicyRouteRuleFragmentMatchNonFrag.IsNull() || o.LeafPolicyRouteRuleFragmentMatchNonFrag.IsUnknown()) {
		vyosData["match-non-frag"] = o.LeafPolicyRouteRuleFragmentMatchNonFrag.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleFragment) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "fragment"}})

	// Leafs
	if value, ok := vyosData["match-frag"]; ok {
		o.LeafPolicyRouteRuleFragmentMatchFrag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleFragmentMatchFrag = basetypes.NewStringNull()
	}
	if value, ok := vyosData["match-non-frag"]; ok {
		o.LeafPolicyRouteRuleFragmentMatchNonFrag = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleFragmentMatchNonFrag = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "fragment"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleFragment) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"match_frag":     types.StringType,
		"match_non_frag": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleFragment) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_frag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Second and further fragments of fragmented packets

`,
		},

		"match_non_frag": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Head fragments or unfragmented packets

`,
		},

		// TagNodes

		// Nodes

	}
}
