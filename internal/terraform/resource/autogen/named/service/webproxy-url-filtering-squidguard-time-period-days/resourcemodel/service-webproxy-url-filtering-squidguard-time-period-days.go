// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceWebproxyURLFilteringSquIDguardTimePeriodDays describes the resource data model.
type ServiceWebproxyURLFilteringSquIDguardTimePeriodDays struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"days_id" vyos:",self-id"`

	ParentIDServiceWebproxyURLFilteringSquIDguardTimePeriod types.String `tfsdk:"time_period" vyos:"time-period,parent-id"`

	// LeafNodes
	LeafServiceWebproxyURLFilteringSquIDguardTimePeriodDaysTime types.String `tfsdk:"time" vyos:"time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceWebproxyURLFilteringSquIDguardTimePeriodDays) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyURLFilteringSquIDguardTimePeriodDays) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"webproxy",

		"url-filtering",

		"squidguard",

		"time-period",
		o.ParentIDServiceWebproxyURLFilteringSquIDguardTimePeriod.ValueString(),

		"days",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyURLFilteringSquIDguardTimePeriodDays) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"days_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Time-period days

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  Sun  &emsp; |  Sunday  |
    |  Mon  &emsp; |  Monday  |
    |  Tue  &emsp; |  Tuesday  |
    |  Wed  &emsp; |  Wednesday  |
    |  Thu  &emsp; |  Thursday  |
    |  Fri  &emsp; |  Friday  |
    |  Sat  &emsp; |  Saturday  |
    |  weekdays  &emsp; |  Monday through Friday  |
    |  weekend  &emsp; |  Saturday and Sunday  |
    |  all  &emsp; |  All days of the week  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"time_period_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Time period name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"time": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time for time-period

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <hh:mm - hh:mm>  &emsp; |  Time range in 24hr time  |

`,
		},

		// Nodes

	}
}
