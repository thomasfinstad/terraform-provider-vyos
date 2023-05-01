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

// PolicyRoutesixRuleIPsec describes the resource data model.
type PolicyRoutesixRuleIPsec struct {
	// LeafNodes
	LeafPolicyRoutesixRuleIPsecMatchIPsec types.String `tfsdk:"match_ipsec"`
	LeafPolicyRoutesixRuleIPsecMatchNone  types.String `tfsdk:"match_none"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRoutesixRuleIPsec) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "ipsec"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRoutesixRuleIPsecMatchIPsec.IsNull() || o.LeafPolicyRoutesixRuleIPsecMatchIPsec.IsUnknown()) {
		vyosData["match-ipsec"] = o.LeafPolicyRoutesixRuleIPsecMatchIPsec.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleIPsecMatchNone.IsNull() || o.LeafPolicyRoutesixRuleIPsecMatchNone.IsUnknown()) {
		vyosData["match-none"] = o.LeafPolicyRoutesixRuleIPsecMatchNone.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRoutesixRuleIPsec) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "ipsec"}})

	// Leafs
	if value, ok := vyosData["match-ipsec"]; ok {
		o.LeafPolicyRoutesixRuleIPsecMatchIPsec = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleIPsecMatchIPsec = basetypes.NewStringNull()
	}
	if value, ok := vyosData["match-none"]; ok {
		o.LeafPolicyRoutesixRuleIPsecMatchNone = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleIPsecMatchNone = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "ipsec"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRoutesixRuleIPsec) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"match_ipsec": types.StringType,
		"match_none":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleIPsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"match_ipsec": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"match_none": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound non-IPsec packets

`,
		},

		// TagNodes

		// Nodes

	}
}