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

// PolicyRoutesixRuleTCP describes the resource data model.
type PolicyRoutesixRuleTCP struct {
	// LeafNodes
	LeafPolicyRoutesixRuleTCPMss types.String `tfsdk:"mss"`

	// TagNodes

	// Nodes
	NodePolicyRoutesixRuleTCPFlags types.Object `tfsdk:"flags"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRoutesixRuleTCP) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "tcp"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRoutesixRuleTCPMss.IsNull() || o.LeafPolicyRoutesixRuleTCPMss.IsUnknown()) {
		vyosData["mss"] = o.LeafPolicyRoutesixRuleTCPMss.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePolicyRoutesixRuleTCPFlags.IsNull() || o.NodePolicyRoutesixRuleTCPFlags.IsUnknown()) {
		var subModel PolicyRoutesixRuleTCPFlags
		diags.Append(o.NodePolicyRoutesixRuleTCPFlags.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["flags"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRoutesixRuleTCP) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "tcp"}})

	// Leafs
	if value, ok := vyosData["mss"]; ok {
		o.LeafPolicyRoutesixRuleTCPMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTCPMss = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["flags"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleTCPFlags{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleTCPFlags = data

	} else {
		o.NodePolicyRoutesixRuleTCPFlags = basetypes.NewObjectNull(PolicyRoutesixRuleTCPFlags{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "tcp"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRoutesixRuleTCP) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"mss": types.StringType,

		// Tags

		// Nodes
		"flags": types.ObjectType{AttrTypes: PolicyRoutesixRuleTCPFlags{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16384  |  Maximum segment size  |
|  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |

`,
		},

		// TagNodes

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTCPFlags{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},
	}
}