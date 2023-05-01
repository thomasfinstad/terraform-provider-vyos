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

// PolicyRouteRuleTCPFlags describes the resource data model.
type PolicyRouteRuleTCPFlags struct {
	// LeafNodes
	LeafPolicyRouteRuleTCPFlagsSyn types.String `tfsdk:"syn"`
	LeafPolicyRouteRuleTCPFlagsAck types.String `tfsdk:"ack"`
	LeafPolicyRouteRuleTCPFlagsFin types.String `tfsdk:"fin"`
	LeafPolicyRouteRuleTCPFlagsRst types.String `tfsdk:"rst"`
	LeafPolicyRouteRuleTCPFlagsUrg types.String `tfsdk:"urg"`
	LeafPolicyRouteRuleTCPFlagsPsh types.String `tfsdk:"psh"`
	LeafPolicyRouteRuleTCPFlagsEcn types.String `tfsdk:"ecn"`
	LeafPolicyRouteRuleTCPFlagsCwr types.String `tfsdk:"cwr"`

	// TagNodes

	// Nodes
	NodePolicyRouteRuleTCPFlagsNot types.Object `tfsdk:"not"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleTCPFlags) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "tcp", "flags"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleTCPFlagsSyn.IsNull() || o.LeafPolicyRouteRuleTCPFlagsSyn.IsUnknown()) {
		vyosData["syn"] = o.LeafPolicyRouteRuleTCPFlagsSyn.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsAck.IsNull() || o.LeafPolicyRouteRuleTCPFlagsAck.IsUnknown()) {
		vyosData["ack"] = o.LeafPolicyRouteRuleTCPFlagsAck.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsFin.IsNull() || o.LeafPolicyRouteRuleTCPFlagsFin.IsUnknown()) {
		vyosData["fin"] = o.LeafPolicyRouteRuleTCPFlagsFin.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsRst.IsNull() || o.LeafPolicyRouteRuleTCPFlagsRst.IsUnknown()) {
		vyosData["rst"] = o.LeafPolicyRouteRuleTCPFlagsRst.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsUrg.IsNull() || o.LeafPolicyRouteRuleTCPFlagsUrg.IsUnknown()) {
		vyosData["urg"] = o.LeafPolicyRouteRuleTCPFlagsUrg.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsPsh.IsNull() || o.LeafPolicyRouteRuleTCPFlagsPsh.IsUnknown()) {
		vyosData["psh"] = o.LeafPolicyRouteRuleTCPFlagsPsh.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsEcn.IsNull() || o.LeafPolicyRouteRuleTCPFlagsEcn.IsUnknown()) {
		vyosData["ecn"] = o.LeafPolicyRouteRuleTCPFlagsEcn.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTCPFlagsCwr.IsNull() || o.LeafPolicyRouteRuleTCPFlagsCwr.IsUnknown()) {
		vyosData["cwr"] = o.LeafPolicyRouteRuleTCPFlagsCwr.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePolicyRouteRuleTCPFlagsNot.IsNull() || o.NodePolicyRouteRuleTCPFlagsNot.IsUnknown()) {
		var subModel PolicyRouteRuleTCPFlagsNot
		diags.Append(o.NodePolicyRouteRuleTCPFlagsNot.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["not"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleTCPFlags) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "tcp", "flags"}})

	// Leafs
	if value, ok := vyosData["syn"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsSyn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsSyn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ack"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsAck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsAck = basetypes.NewStringNull()
	}
	if value, ok := vyosData["fin"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsFin = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsFin = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rst"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsRst = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsRst = basetypes.NewStringNull()
	}
	if value, ok := vyosData["urg"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsUrg = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsUrg = basetypes.NewStringNull()
	}
	if value, ok := vyosData["psh"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsPsh = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsPsh = basetypes.NewStringNull()
	}
	if value, ok := vyosData["ecn"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsEcn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsEcn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["cwr"]; ok {
		o.LeafPolicyRouteRuleTCPFlagsCwr = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTCPFlagsCwr = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["not"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRouteRuleTCPFlagsNot{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRouteRuleTCPFlagsNot = data

	} else {
		o.NodePolicyRouteRuleTCPFlagsNot = basetypes.NewObjectNull(PolicyRouteRuleTCPFlagsNot{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "tcp", "flags"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleTCPFlags) AttributeTypes() map[string]attr.Type {
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
		"not": types.ObjectType{AttrTypes: PolicyRouteRuleTCPFlagsNot{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleTCPFlags) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"not": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTCPFlagsNot{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match flags not set

`,
		},
	}
}