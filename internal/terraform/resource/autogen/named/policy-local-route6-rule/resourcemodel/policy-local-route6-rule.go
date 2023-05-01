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

// PolicyLocalRoutesixRule describes the resource data model.
type PolicyLocalRoutesixRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPolicyLocalRoutesixRuleFwmark           types.String `tfsdk:"fwmark"`
	LeafPolicyLocalRoutesixRuleSource           types.String `tfsdk:"source"`
	LeafPolicyLocalRoutesixRuleDestination      types.String `tfsdk:"destination"`
	LeafPolicyLocalRoutesixRuleInboundInterface types.String `tfsdk:"inbound_interface"`

	// TagNodes

	// Nodes
	NodePolicyLocalRoutesixRuleSet types.Object `tfsdk:"set"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLocalRoutesixRule) GetVyosPath() []string {
	return []string{
		"policy",
		"local-route6",
		"rule",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyLocalRoutesixRule) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "local-route6", "rule"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyLocalRoutesixRuleFwmark.IsNull() || o.LeafPolicyLocalRoutesixRuleFwmark.IsUnknown()) {
		vyosData["fwmark"] = o.LeafPolicyLocalRoutesixRuleFwmark.ValueString()
	}
	if !(o.LeafPolicyLocalRoutesixRuleSource.IsNull() || o.LeafPolicyLocalRoutesixRuleSource.IsUnknown()) {
		vyosData["source"] = o.LeafPolicyLocalRoutesixRuleSource.ValueString()
	}
	if !(o.LeafPolicyLocalRoutesixRuleDestination.IsNull() || o.LeafPolicyLocalRoutesixRuleDestination.IsUnknown()) {
		vyosData["destination"] = o.LeafPolicyLocalRoutesixRuleDestination.ValueString()
	}
	if !(o.LeafPolicyLocalRoutesixRuleInboundInterface.IsNull() || o.LeafPolicyLocalRoutesixRuleInboundInterface.IsUnknown()) {
		vyosData["inbound-interface"] = o.LeafPolicyLocalRoutesixRuleInboundInterface.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePolicyLocalRoutesixRuleSet.IsNull() || o.NodePolicyLocalRoutesixRuleSet.IsUnknown()) {
		var subModel PolicyLocalRoutesixRuleSet
		diags.Append(o.NodePolicyLocalRoutesixRuleSet.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["set"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyLocalRoutesixRule) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "local-route6", "rule"}})

	// Leafs
	if value, ok := vyosData["fwmark"]; ok {
		o.LeafPolicyLocalRoutesixRuleFwmark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleFwmark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source"]; ok {
		o.LeafPolicyLocalRoutesixRuleSource = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleSource = basetypes.NewStringNull()
	}
	if value, ok := vyosData["destination"]; ok {
		o.LeafPolicyLocalRoutesixRuleDestination = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleDestination = basetypes.NewStringNull()
	}
	if value, ok := vyosData["inbound-interface"]; ok {
		o.LeafPolicyLocalRoutesixRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyLocalRoutesixRuleInboundInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["set"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyLocalRoutesixRuleSet{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyLocalRoutesixRuleSet = data

	} else {
		o.NodePolicyLocalRoutesixRuleSet = basetypes.NewObjectNull(PolicyLocalRoutesixRuleSet{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "local-route6", "rule"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyLocalRoutesixRule) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"fwmark":            types.StringType,
		"source":            types.StringType,
		"destination":       types.StringType,
		"inbound_interface": types.StringType,

		// Tags

		// Nodes
		"set": types.ObjectType{AttrTypes: PolicyLocalRoutesixRuleSet{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 policy local-route rule set number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32765  |  Local-route rule number (1-32765)  |

`,
		},

		// LeafNodes

		"fwmark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Address to match against  |

`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Address to match against  |
|  ipv6net  |  Prefix to match against  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Address to match against  |
|  ipv6net  |  Prefix to match against  |

`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound Interface

`,
		},

		// TagNodes

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},
	}
}