// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyNetworkEmulator{}

// QosPolicyNetworkEmulator describes the resource data model.
type QosPolicyNetworkEmulator struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"network_emulator_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyNetworkEmulatorDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyNetworkEmulatorBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyNetworkEmulatorDelay       types.String `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafQosPolicyNetworkEmulatorCorruption  types.String `tfsdk:"corruption" vyos:"corruption,omitempty"`
	LeafQosPolicyNetworkEmulatorDuplicate   types.String `tfsdk:"duplicate" vyos:"duplicate,omitempty"`
	LeafQosPolicyNetworkEmulatorLoss        types.String `tfsdk:"loss" vyos:"loss,omitempty"`
	LeafQosPolicyNetworkEmulatorReordering  types.String `tfsdk:"reordering" vyos:"reordering,omitempty"`
	LeafQosPolicyNetworkEmulatorQueueLimit  types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyNetworkEmulator) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyNetworkEmulator) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyNetworkEmulator) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyNetworkEmulator) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"network-emulator",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyNetworkEmulator) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyNetworkEmulator) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyNetworkEmulator) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"network_emulator_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Network emulator policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Network emulator policy

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
							"double underscores in network_emulator_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  network_emulator_id, value must match: ^[a-zA-Z0-9-_]*$",
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

		"delay": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adds delay to packets outgoing to chosen network interface

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
			Description: `Adds delay to packets outgoing to chosen network interface

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
		},

		"corruption": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Introducing error in a random position for chosen percent of packets

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Introducing error in a random position for chosen percent of packets

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"duplicate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cosen percent of packets is duplicated before queuing them

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Cosen percent of packets is duplicated before queuing them

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"loss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Add independent loss probability to the packets outgoing to chosen network interface

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Add independent loss probability to the packets outgoing to chosen network interface

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"reordering": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Emulated packet reordering percentage

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Emulated packet reordering percentage

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
			Description: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
		},

		// Nodes

	}
}
