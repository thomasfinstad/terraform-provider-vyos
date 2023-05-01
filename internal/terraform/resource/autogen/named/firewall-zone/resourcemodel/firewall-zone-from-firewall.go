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

// FirewallZoneFromFirewall describes the resource data model.
type FirewallZoneFromFirewall struct {
	// LeafNodes
	LeafFirewallZoneFromFirewallIPvsixName types.String `tfsdk:"ipv6_name"`
	LeafFirewallZoneFromFirewallName       types.String `tfsdk:"name"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallZoneFromFirewall) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "zone", "from", "firewall"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallZoneFromFirewallIPvsixName.IsNull() || o.LeafFirewallZoneFromFirewallIPvsixName.IsUnknown()) {
		vyosData["ipv6-name"] = o.LeafFirewallZoneFromFirewallIPvsixName.ValueString()
	}
	if !(o.LeafFirewallZoneFromFirewallName.IsNull() || o.LeafFirewallZoneFromFirewallName.IsUnknown()) {
		vyosData["name"] = o.LeafFirewallZoneFromFirewallName.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallZoneFromFirewall) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "zone", "from", "firewall"}})

	// Leafs
	if value, ok := vyosData["ipv6-name"]; ok {
		o.LeafFirewallZoneFromFirewallIPvsixName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallZoneFromFirewallIPvsixName = basetypes.NewStringNull()
	}
	if value, ok := vyosData["name"]; ok {
		o.LeafFirewallZoneFromFirewallName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallZoneFromFirewallName = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "zone", "from", "firewall"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallZoneFromFirewall) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"ipv6_name": types.StringType,
		"name":      types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneFromFirewall) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv6_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 firewall ruleset

`,
		},

		"name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 firewall ruleset

`,
		},

		// TagNodes

		// Nodes

	}
}
