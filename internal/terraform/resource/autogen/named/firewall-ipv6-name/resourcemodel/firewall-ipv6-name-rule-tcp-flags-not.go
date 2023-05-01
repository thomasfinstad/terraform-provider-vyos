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

// FirewallIPvsixNameRuleTCPFlagsNot describes the resource data model.
type FirewallIPvsixNameRuleTCPFlagsNot struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleTCPFlagsNotSyn types.String `tfsdk:"syn"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotAck types.String `tfsdk:"ack"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotFin types.String `tfsdk:"fin"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotRst types.String `tfsdk:"rst"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotUrg types.String `tfsdk:"urg"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotPsh types.String `tfsdk:"psh"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotEcn types.String `tfsdk:"ecn"`
	LeafFirewallIPvsixNameRuleTCPFlagsNotCwr types.String `tfsdk:"cwr"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallIPvsixNameRuleTCPFlagsNot) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "tcp", "flags", "not"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotSyn.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotSyn.IsUnknown()) {
		vyosData["syn"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotSyn.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotAck.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotAck.IsUnknown()) {
		vyosData["ack"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotAck.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotFin.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotFin.IsUnknown()) {
		vyosData["fin"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotFin.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotRst.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotRst.IsUnknown()) {
		vyosData["rst"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotRst.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotUrg.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotUrg.IsUnknown()) {
		vyosData["urg"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotUrg.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotPsh.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotPsh.IsUnknown()) {
		vyosData["psh"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotPsh.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotEcn.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotEcn.IsUnknown()) {
		vyosData["ecn"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotEcn.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleTCPFlagsNotCwr.IsNull() || o.LeafFirewallIPvsixNameRuleTCPFlagsNotCwr.IsUnknown()) {
		vyosData["cwr"] = o.LeafFirewallIPvsixNameRuleTCPFlagsNotCwr.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallIPvsixNameRuleTCPFlagsNot) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "tcp", "flags", "not"}})

	// Leafs
	if value, ok := vyosData["syn"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotSyn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ack"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotAck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["fin"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotFin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotFin = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rst"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotRst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotRst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["urg"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotUrg = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotUrg = basetypes.NewStringNull()
	}
	if value, ok := vyosData["psh"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotPsh = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotPsh = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ecn"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotEcn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotEcn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["cwr"]; ok {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotCwr = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleTCPFlagsNotCwr = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule", "tcp", "flags", "not"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallIPvsixNameRuleTCPFlagsNot) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"syn": types.StringType,
		"ack": types.StringType,
		"fin": types.StringType,
		"rst": types.StringType,
		"urg": types.StringType,
		"psh": types.StringType,
		"ecn": types.StringType,
		"cwr": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleTCPFlagsNot) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Synchronise flag

`,
		},

		"ack": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Acknowledge flag

`,
		},

		"fin": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Finish flag

`,
		},

		"rst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Reset flag

`,
		},

		"urg": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Urgent flag

`,
		},

		"psh": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Push flag

`,
		},

		"ecn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Explicit Congestion Notification flag

`,
		},

		"cwr": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
		},

		// TagNodes

		// Nodes

	}
}
