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

// FirewallNameRuleTCPFlagsNot describes the resource data model.
type FirewallNameRuleTCPFlagsNot struct {
	// LeafNodes
	LeafFirewallNameRuleTCPFlagsNotSyn types.String `tfsdk:"syn"`
	LeafFirewallNameRuleTCPFlagsNotAck types.String `tfsdk:"ack"`
	LeafFirewallNameRuleTCPFlagsNotFin types.String `tfsdk:"fin"`
	LeafFirewallNameRuleTCPFlagsNotRst types.String `tfsdk:"rst"`
	LeafFirewallNameRuleTCPFlagsNotUrg types.String `tfsdk:"urg"`
	LeafFirewallNameRuleTCPFlagsNotPsh types.String `tfsdk:"psh"`
	LeafFirewallNameRuleTCPFlagsNotEcn types.String `tfsdk:"ecn"`
	LeafFirewallNameRuleTCPFlagsNotCwr types.String `tfsdk:"cwr"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallNameRuleTCPFlagsNot) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "tcp", "flags", "not"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallNameRuleTCPFlagsNotSyn.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotSyn.IsUnknown()) {
		vyosData["syn"] = o.LeafFirewallNameRuleTCPFlagsNotSyn.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotAck.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotAck.IsUnknown()) {
		vyosData["ack"] = o.LeafFirewallNameRuleTCPFlagsNotAck.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotFin.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotFin.IsUnknown()) {
		vyosData["fin"] = o.LeafFirewallNameRuleTCPFlagsNotFin.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotRst.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotRst.IsUnknown()) {
		vyosData["rst"] = o.LeafFirewallNameRuleTCPFlagsNotRst.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotUrg.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotUrg.IsUnknown()) {
		vyosData["urg"] = o.LeafFirewallNameRuleTCPFlagsNotUrg.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotPsh.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotPsh.IsUnknown()) {
		vyosData["psh"] = o.LeafFirewallNameRuleTCPFlagsNotPsh.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotEcn.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotEcn.IsUnknown()) {
		vyosData["ecn"] = o.LeafFirewallNameRuleTCPFlagsNotEcn.ValueString()
	}
	if !(o.LeafFirewallNameRuleTCPFlagsNotCwr.IsNull() || o.LeafFirewallNameRuleTCPFlagsNotCwr.IsUnknown()) {
		vyosData["cwr"] = o.LeafFirewallNameRuleTCPFlagsNotCwr.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallNameRuleTCPFlagsNot) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "tcp", "flags", "not"}})

	// Leafs
	if value, ok := vyosData["syn"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotSyn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ack"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotAck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["fin"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotFin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotFin = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rst"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotRst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotRst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["urg"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotUrg = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotUrg = basetypes.NewStringNull()
	}
	if value, ok := vyosData["psh"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotPsh = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotPsh = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ecn"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotEcn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotEcn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["cwr"]; ok {
		o.LeafFirewallNameRuleTCPFlagsNotCwr = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTCPFlagsNotCwr = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "tcp", "flags", "not"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallNameRuleTCPFlagsNot) AttributeTypes() map[string]attr.Type {
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
func (o FirewallNameRuleTCPFlagsNot) ResourceSchemaAttributes() map[string]schema.Attribute {
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
