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

// FirewallNameRuleTime describes the resource data model.
type FirewallNameRuleTime struct {
	// LeafNodes
	LeafFirewallNameRuleTimeStartdate types.String `tfsdk:"startdate"`
	LeafFirewallNameRuleTimeStarttime types.String `tfsdk:"starttime"`
	LeafFirewallNameRuleTimeStopdate  types.String `tfsdk:"stopdate"`
	LeafFirewallNameRuleTimeStoptime  types.String `tfsdk:"stoptime"`
	LeafFirewallNameRuleTimeWeekdays  types.String `tfsdk:"weekdays"`

	// TagNodes

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallNameRuleTime) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "time"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallNameRuleTimeStartdate.IsNull() || o.LeafFirewallNameRuleTimeStartdate.IsUnknown()) {
		vyosData["startdate"] = o.LeafFirewallNameRuleTimeStartdate.ValueString()
	}
	if !(o.LeafFirewallNameRuleTimeStarttime.IsNull() || o.LeafFirewallNameRuleTimeStarttime.IsUnknown()) {
		vyosData["starttime"] = o.LeafFirewallNameRuleTimeStarttime.ValueString()
	}
	if !(o.LeafFirewallNameRuleTimeStopdate.IsNull() || o.LeafFirewallNameRuleTimeStopdate.IsUnknown()) {
		vyosData["stopdate"] = o.LeafFirewallNameRuleTimeStopdate.ValueString()
	}
	if !(o.LeafFirewallNameRuleTimeStoptime.IsNull() || o.LeafFirewallNameRuleTimeStoptime.IsUnknown()) {
		vyosData["stoptime"] = o.LeafFirewallNameRuleTimeStoptime.ValueString()
	}
	if !(o.LeafFirewallNameRuleTimeWeekdays.IsNull() || o.LeafFirewallNameRuleTimeWeekdays.IsUnknown()) {
		vyosData["weekdays"] = o.LeafFirewallNameRuleTimeWeekdays.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallNameRuleTime) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "time"}})

	// Leafs
	if value, ok := vyosData["startdate"]; ok {
		o.LeafFirewallNameRuleTimeStartdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTimeStartdate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["starttime"]; ok {
		o.LeafFirewallNameRuleTimeStarttime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTimeStarttime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["stopdate"]; ok {
		o.LeafFirewallNameRuleTimeStopdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTimeStopdate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["stoptime"]; ok {
		o.LeafFirewallNameRuleTimeStoptime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTimeStoptime = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weekdays"]; ok {
		o.LeafFirewallNameRuleTimeWeekdays = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleTimeWeekdays = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "name", "rule", "time"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallNameRuleTime) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"startdate": types.StringType,
		"starttime": types.StringType,
		"stopdate":  types.StringType,
		"stoptime":  types.StringType,
		"weekdays":  types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRuleTime) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"startdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to start matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter date using following notation - YYYY-MM-DD  |

`,
		},

		"starttime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to start matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter time using using 24 hour notation - hh:mm:ss  |

`,
		},

		"stopdate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Date to stop matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter date using following notation - YYYY-MM-DD  |

`,
		},

		"stoptime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time of day to stop matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter time using using 24 hour notation - hh:mm:ss  |

`,
		},

		"weekdays": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Comma separated weekdays to match rule on

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday,                           Saturday, Sunday)  |
|  u32:0-6  |  Day number (0 = Sunday ... 6 = Saturday)  |

`,
		},

		// TagNodes

		// Nodes

	}
}
