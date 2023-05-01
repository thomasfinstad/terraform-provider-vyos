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

// FirewallNameRuleInboundInterface describes the resource data model.
type FirewallNameRuleInboundInterface struct {
	// LeafNodes
	LeafFirewallNameRuleInboundInterfaceInterfaceName  types.String `tfsdk:"interface_name"`
	LeafFirewallNameRuleInboundInterfaceInterfaceGroup types.String `tfsdk:"interface_group"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallNameRuleInboundInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "inbound-interface"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallNameRuleInboundInterfaceInterfaceName.IsNull() || o.LeafFirewallNameRuleInboundInterfaceInterfaceName.IsUnknown()) {
		vyosData["interface-name"] = o.LeafFirewallNameRuleInboundInterfaceInterfaceName.ValueString()
	}
	if !(o.LeafFirewallNameRuleInboundInterfaceInterfaceGroup.IsNull() || o.LeafFirewallNameRuleInboundInterfaceInterfaceGroup.IsUnknown()) {
		vyosData["interface-group"] = o.LeafFirewallNameRuleInboundInterfaceInterfaceGroup.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallNameRuleInboundInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "inbound-interface"}})

	// Leafs
	if value, ok := vyosData["interface-name"]; ok {
		o.LeafFirewallNameRuleInboundInterfaceInterfaceName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleInboundInterfaceInterfaceName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface-group"]; ok {
		o.LeafFirewallNameRuleInboundInterfaceInterfaceGroup = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleInboundInterfaceInterfaceGroup = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "inbound-interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallNameRuleInboundInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"interface_name":  types.StringType,
		"interface_group": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleInboundInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

`,
		},

		"interface_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

`,
		},

		// TagNodes

		// Nodes

	}
}
