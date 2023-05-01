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

// FirewallIPvsixNameRuleConnectionStatus describes the resource data model.
type FirewallIPvsixNameRuleConnectionStatus struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleConnectionStatusNat types.String `tfsdk:"nat"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallIPvsixNameRuleConnectionStatus) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "connection-status"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallIPvsixNameRuleConnectionStatusNat.IsNull() || o.LeafFirewallIPvsixNameRuleConnectionStatusNat.IsUnknown()) {
		vyosData["nat"] = o.LeafFirewallIPvsixNameRuleConnectionStatusNat.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallIPvsixNameRuleConnectionStatus) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "connection-status"}})

	// Leafs
	if value, ok := vyosData["nat"]; ok {
		o.LeafFirewallIPvsixNameRuleConnectionStatusNat = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleConnectionStatusNat = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "connection-status"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallIPvsixNameRuleConnectionStatus) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"nat": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleConnectionStatus) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"nat": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAT connection status

|  Format  |  Description  |
|----------|---------------|
|  destination  |  Match connections that are subject to destination NAT  |
|  source  |  Match connections that are subject to source NAT  |

`,
		},

		// TagNodes

		// Nodes

	}
}