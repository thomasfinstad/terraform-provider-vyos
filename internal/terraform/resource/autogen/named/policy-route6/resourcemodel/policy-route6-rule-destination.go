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

// PolicyRoutesixRuleDestination describes the resource data model.
type PolicyRoutesixRuleDestination struct {
	// LeafNodes
	LeafPolicyRoutesixRuleDestinationAddress types.String `tfsdk:"address"`
	LeafPolicyRoutesixRuleDestinationPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
	NodePolicyRoutesixRuleDestinationGroup types.Object `tfsdk:"group"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRoutesixRuleDestination) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "destination"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRoutesixRuleDestinationAddress.IsNull() || o.LeafPolicyRoutesixRuleDestinationAddress.IsUnknown()) {
		vyosData["address"] = o.LeafPolicyRoutesixRuleDestinationAddress.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleDestinationPort.IsNull() || o.LeafPolicyRoutesixRuleDestinationPort.IsUnknown()) {
		vyosData["port"] = o.LeafPolicyRoutesixRuleDestinationPort.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePolicyRoutesixRuleDestinationGroup.IsNull() || o.NodePolicyRoutesixRuleDestinationGroup.IsUnknown()) {
		var subModel PolicyRoutesixRuleDestinationGroup
		diags.Append(o.NodePolicyRoutesixRuleDestinationGroup.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["group"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRoutesixRuleDestination) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "destination"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafPolicyRoutesixRuleDestinationAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafPolicyRoutesixRuleDestinationPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDestinationPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["group"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleDestinationGroup{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleDestinationGroup = data

	} else {
		o.NodePolicyRoutesixRuleDestinationGroup = basetypes.NewObjectNull(PolicyRoutesixRuleDestinationGroup{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route6", "rule", "destination"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRoutesixRuleDestination) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes
		"group": types.ObjectType{AttrTypes: PolicyRoutesixRuleDestinationGroup{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP address to match  |
|  ipv6net  |  Subnet to match  |
|  ipv6range  |  IP range to match  |
|  !ipv6  |  Match everything except the specified address  |
|  !ipv6net  |  Match everything except the specified prefix  |
|  !ipv6range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleDestinationGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
