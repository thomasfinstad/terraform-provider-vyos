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

// PolicyRouteRuleTime describes the resource data model.
type PolicyRouteRuleTime struct {
	// LeafNodes
	LeafPolicyRouteRuleTimeMonthdays types.String `tfsdk:"monthdays"`
	LeafPolicyRouteRuleTimeStartdate types.String `tfsdk:"startdate"`
	LeafPolicyRouteRuleTimeStarttime types.String `tfsdk:"starttime"`
	LeafPolicyRouteRuleTimeStopdate  types.String `tfsdk:"stopdate"`
	LeafPolicyRouteRuleTimeStoptime  types.String `tfsdk:"stoptime"`
	LeafPolicyRouteRuleTimeUtc       types.String `tfsdk:"utc"`
	LeafPolicyRouteRuleTimeWeekdays  types.String `tfsdk:"weekdays"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRouteRuleTime) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route", "rule", "time"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRouteRuleTimeMonthdays.IsNull() || o.LeafPolicyRouteRuleTimeMonthdays.IsUnknown()) {
		vyosData["monthdays"] = o.LeafPolicyRouteRuleTimeMonthdays.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeStartdate.IsNull() || o.LeafPolicyRouteRuleTimeStartdate.IsUnknown()) {
		vyosData["startdate"] = o.LeafPolicyRouteRuleTimeStartdate.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeStarttime.IsNull() || o.LeafPolicyRouteRuleTimeStarttime.IsUnknown()) {
		vyosData["starttime"] = o.LeafPolicyRouteRuleTimeStarttime.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeStopdate.IsNull() || o.LeafPolicyRouteRuleTimeStopdate.IsUnknown()) {
		vyosData["stopdate"] = o.LeafPolicyRouteRuleTimeStopdate.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeStoptime.IsNull() || o.LeafPolicyRouteRuleTimeStoptime.IsUnknown()) {
		vyosData["stoptime"] = o.LeafPolicyRouteRuleTimeStoptime.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeUtc.IsNull() || o.LeafPolicyRouteRuleTimeUtc.IsUnknown()) {
		vyosData["utc"] = o.LeafPolicyRouteRuleTimeUtc.ValueString()
	}
	if !(o.LeafPolicyRouteRuleTimeWeekdays.IsNull() || o.LeafPolicyRouteRuleTimeWeekdays.IsUnknown()) {
		vyosData["weekdays"] = o.LeafPolicyRouteRuleTimeWeekdays.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRouteRuleTime) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route", "rule", "time"}})

	// Leafs
	if value, ok := vyosData["monthdays"]; ok {
		o.LeafPolicyRouteRuleTimeMonthdays = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeMonthdays = basetypes.NewStringNull()
	}
	if value, ok := vyosData["startdate"]; ok {
		o.LeafPolicyRouteRuleTimeStartdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeStartdate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["starttime"]; ok {
		o.LeafPolicyRouteRuleTimeStarttime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeStarttime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["stopdate"]; ok {
		o.LeafPolicyRouteRuleTimeStopdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeStopdate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["stoptime"]; ok {
		o.LeafPolicyRouteRuleTimeStoptime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeStoptime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["utc"]; ok {
		o.LeafPolicyRouteRuleTimeUtc = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeUtc = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weekdays"]; ok {
		o.LeafPolicyRouteRuleTimeWeekdays = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleTimeWeekdays = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route", "rule", "time"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRouteRuleTime) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"monthdays": types.StringType,
		"startdate": types.StringType,
		"starttime": types.StringType,
		"stopdate":  types.StringType,
		"stoptime":  types.StringType,
		"utc":       types.StringType,
		"weekdays":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRuleTime) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"monthdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Monthdays to match rule on

`,
		},

		"startdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to start matching rule

`,
		},

		"starttime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to start matching rule

`,
		},

		"stopdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to stop matching rule

`,
		},

		"stoptime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to stop matching rule

`,
		},

		"utc": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interpret times for startdate, stopdate, starttime and stoptime to be UTC

`,
		},

		"weekdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Weekdays to match rule on

`,
		},

		// TagNodes

		// Nodes

	}
}
