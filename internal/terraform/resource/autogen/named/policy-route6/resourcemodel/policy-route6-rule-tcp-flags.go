// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyRoutesixRuleTCPFlags describes the resource data model.
type PolicyRoutesixRuleTCPFlags struct {
	// LeafNodes
	PolicyRoutesixRuleTCPFlagsSyn customtypes.CustomStringValue `tfsdk:"syn" json:"syn,omitempty"`
	PolicyRoutesixRuleTCPFlagsAck customtypes.CustomStringValue `tfsdk:"ack" json:"ack,omitempty"`
	PolicyRoutesixRuleTCPFlagsFin customtypes.CustomStringValue `tfsdk:"fin" json:"fin,omitempty"`
	PolicyRoutesixRuleTCPFlagsRst customtypes.CustomStringValue `tfsdk:"rst" json:"rst,omitempty"`
	PolicyRoutesixRuleTCPFlagsUrg customtypes.CustomStringValue `tfsdk:"urg" json:"urg,omitempty"`
	PolicyRoutesixRuleTCPFlagsPsh customtypes.CustomStringValue `tfsdk:"psh" json:"psh,omitempty"`
	PolicyRoutesixRuleTCPFlagsEcn customtypes.CustomStringValue `tfsdk:"ecn" json:"ecn,omitempty"`
	PolicyRoutesixRuleTCPFlagsCwr customtypes.CustomStringValue `tfsdk:"cwr" json:"cwr,omitempty"`

	// TagNodes

	// Nodes
	PolicyRoutesixRuleTCPFlagsNot types.Object `tfsdk:"not" json:"not,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyRoutesixRuleTCPFlags) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"syn": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Synchronise flag

`,
		},

		"ack": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Acknowledge flag

`,
		},

		"fin": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Finish flag

`,
		},

		"rst": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Reset flag

`,
		},

		"urg": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Urgent flag

`,
		},

		"psh": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Push flag

`,
		},

		"ecn": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Explicit Congestion Notification flag

`,
		},

		"cwr": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Congestion Window Reduced flag

`,
		},

		// TagNodes

		// Nodes

		"not": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTCPFlagsNot{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match flags not set

`,
		},
	}
}
