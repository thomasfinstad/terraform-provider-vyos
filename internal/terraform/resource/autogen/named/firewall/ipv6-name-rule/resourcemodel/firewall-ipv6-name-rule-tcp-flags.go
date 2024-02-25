// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixNameRuleTCPFlags describes the resource data model.
type FirewallIPvsixNameRuleTCPFlags struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleTCPFlagsSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsFin types.Bool `tfsdk:"fin" vyos:"fin,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsRst types.Bool `tfsdk:"rst" vyos:"rst,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsUrg types.Bool `tfsdk:"urg" vyos:"urg,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsPsh types.Bool `tfsdk:"psh" vyos:"psh,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsEcn types.Bool `tfsdk:"ecn" vyos:"ecn,omitempty"`
	LeafFirewallIPvsixNameRuleTCPFlagsCwr types.Bool `tfsdk:"cwr" vyos:"cwr,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvsixNameRuleTCPFlagsNot *FirewallIPvsixNameRuleTCPFlagsNot `tfsdk:"not" vyos:"not,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleTCPFlags) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Synchronise flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ack": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acknowledge flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fin": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Finish flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rst": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reset flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"urg": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Urgent flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"psh": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Push flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ecn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Explicit Congestion Notification flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cwr": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"not": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCPFlagsNot{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match flags not set

`,
		},
	}
}