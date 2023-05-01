// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRoutesixRuleTime describes the resource data model.
type PolicyRoutesixRuleTime struct {
	// LeafNodes
	LeafPolicyRoutesixRuleTimeMonthdays types.String `tfsdk:"monthdays" json:"monthdays,omitempty"`
	LeafPolicyRoutesixRuleTimeStartdate types.String `tfsdk:"startdate" json:"startdate,omitempty"`
	LeafPolicyRoutesixRuleTimeStarttime types.String `tfsdk:"starttime" json:"starttime,omitempty"`
	LeafPolicyRoutesixRuleTimeStopdate  types.String `tfsdk:"stopdate" json:"stopdate,omitempty"`
	LeafPolicyRoutesixRuleTimeStoptime  types.String `tfsdk:"stoptime" json:"stoptime,omitempty"`
	LeafPolicyRoutesixRuleTimeUtc       types.String `tfsdk:"utc" json:"utc,omitempty"`
	LeafPolicyRoutesixRuleTimeWeekdays  types.String `tfsdk:"weekdays" json:"weekdays,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleTime) ResourceSchemaAttributes() map[string]schema.Attribute {
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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRoutesixRuleTime) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRoutesixRuleTimeMonthdays.IsNull() && !o.LeafPolicyRoutesixRuleTimeMonthdays.IsUnknown() {
		jsonData["monthdays"] = o.LeafPolicyRoutesixRuleTimeMonthdays.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeStartdate.IsNull() && !o.LeafPolicyRoutesixRuleTimeStartdate.IsUnknown() {
		jsonData["startdate"] = o.LeafPolicyRoutesixRuleTimeStartdate.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeStarttime.IsNull() && !o.LeafPolicyRoutesixRuleTimeStarttime.IsUnknown() {
		jsonData["starttime"] = o.LeafPolicyRoutesixRuleTimeStarttime.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeStopdate.IsNull() && !o.LeafPolicyRoutesixRuleTimeStopdate.IsUnknown() {
		jsonData["stopdate"] = o.LeafPolicyRoutesixRuleTimeStopdate.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeStoptime.IsNull() && !o.LeafPolicyRoutesixRuleTimeStoptime.IsUnknown() {
		jsonData["stoptime"] = o.LeafPolicyRoutesixRuleTimeStoptime.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeUtc.IsNull() && !o.LeafPolicyRoutesixRuleTimeUtc.IsUnknown() {
		jsonData["utc"] = o.LeafPolicyRoutesixRuleTimeUtc.ValueString()
	}

	if !o.LeafPolicyRoutesixRuleTimeWeekdays.IsNull() && !o.LeafPolicyRoutesixRuleTimeWeekdays.IsUnknown() {
		jsonData["weekdays"] = o.LeafPolicyRoutesixRuleTimeWeekdays.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PolicyRoutesixRuleTime) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["monthdays"]; ok {
		o.LeafPolicyRoutesixRuleTimeMonthdays = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeMonthdays = basetypes.NewStringNull()
	}

	if value, ok := jsonData["startdate"]; ok {
		o.LeafPolicyRoutesixRuleTimeStartdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeStartdate = basetypes.NewStringNull()
	}

	if value, ok := jsonData["starttime"]; ok {
		o.LeafPolicyRoutesixRuleTimeStarttime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeStarttime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["stopdate"]; ok {
		o.LeafPolicyRoutesixRuleTimeStopdate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeStopdate = basetypes.NewStringNull()
	}

	if value, ok := jsonData["stoptime"]; ok {
		o.LeafPolicyRoutesixRuleTimeStoptime = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeStoptime = basetypes.NewStringNull()
	}

	if value, ok := jsonData["utc"]; ok {
		o.LeafPolicyRoutesixRuleTimeUtc = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeUtc = basetypes.NewStringNull()
	}

	if value, ok := jsonData["weekdays"]; ok {
		o.LeafPolicyRoutesixRuleTimeWeekdays = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleTimeWeekdays = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	return nil
}
