// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyLimiterClass{}

// QosPolicyLimiterClass describes the resource data model.
type QosPolicyLimiterClass struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:"-,self-id"`

	ParentIDQosPolicyLimiter types.String `tfsdk:"limiter_id" vyos:"limiter,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyLimiterClassDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyLimiterClassBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyLimiterClassBurst       types.String `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafQosPolicyLimiterClassMtu         types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafQosPolicyLimiterClassExceed      types.String `tfsdk:"exceed" vyos:"exceed,omitempty"`
	LeafQosPolicyLimiterClassNotExceed   types.String `tfsdk:"not_exceed" vyos:"not-exceed,omitempty"`
	LeafQosPolicyLimiterClassPriority    types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyLimiterClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyLimiterClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyLimiterClass) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyLimiterClass) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyLimiterClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"class",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyLimiterClass) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"limiter",
		o.ParentIDQosPolicyLimiter.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyLimiterClass) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"limiter",
		o.ParentIDQosPolicyLimiter.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClass) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4090  |  Class Identifier  |
`,
			Description: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4090  |  Class Identifier  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"limiter_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic input limiting policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Traffic input limiting policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in limiter_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  limiter_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
			Description: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

    |  Format            |  Description                             |
    |--------------------|------------------------------------------|
    |  <number>          |  Bytes                                   |
    |  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,
			Description: `Burst size for this class

    |  Format            |  Description                             |
    |--------------------|------------------------------------------|
    |  <number>          |  Bytes                                   |
    |  <number><suffix>  |  Bytes with scaling suffix (kb, mb, gb)  |
`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MTU size for this class

    |  Format     |  Description  |
    |-------------|---------------|
    |  256-65535  |  Bytes        |
`,
			Description: `MTU size for this class

    |  Format     |  Description  |
    |-------------|---------------|
    |  256-65535  |  Bytes        |
`,
		},

		"exceed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action for packets exceeding the limiter

    |  Format      |  Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    |  continue    |  Do not do anything, just continue with the next action in line                                                                |
    |  drop        |  Drop the packet immediately                                                                                                   |
    |  ok          |  Accept the packet                                                                                                             |
    |  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe        |  Pass the packet to the next action in line                                                                                    |
`,
			Description: `Default action for packets exceeding the limiter

    |  Format      |  Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    |  continue    |  Do not do anything, just continue with the next action in line                                                                |
    |  drop        |  Drop the packet immediately                                                                                                   |
    |  ok          |  Accept the packet                                                                                                             |
    |  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe        |  Pass the packet to the next action in line                                                                                    |
`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"not_exceed": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action for packets not exceeding the limiter

    |  Format      |  Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    |  continue    |  Do not do anything, just continue with the next action in line                                                                |
    |  drop        |  Drop the packet immediately                                                                                                   |
    |  ok          |  Accept the packet                                                                                                             |
    |  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe        |  Pass the packet to the next action in line                                                                                    |
`,
			Description: `Default action for packets not exceeding the limiter

    |  Format      |  Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    |  continue    |  Do not do anything, just continue with the next action in line                                                                |
    |  drop        |  Drop the packet immediately                                                                                                   |
    |  ok          |  Accept the packet                                                                                                             |
    |  reclassify  |  Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    |  pipe        |  Pass the packet to the next action in line                                                                                    |
`,

			// Default:          stringdefault.StaticString(`ok`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Priority for rule evaluation

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  0-20    |  Priority for match rule evaluation  |
`,
			Description: `Priority for rule evaluation

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  0-20    |  Priority for match rule evaluation  |
`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		// Nodes

	}
}
