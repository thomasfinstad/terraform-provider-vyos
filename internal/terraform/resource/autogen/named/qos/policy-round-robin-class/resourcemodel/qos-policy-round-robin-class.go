// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

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

var _ helpers.VyosTopResourceDataModel = &QosPolicyRoundRobinClass{}

// QosPolicyRoundRobinClass describes the resource data model.
type QosPolicyRoundRobinClass struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:"-,self-id"`

	ParentIDQosPolicyRoundRobin types.String `tfsdk:"round_robin_id" vyos:"round-robin,parent-id"`

	// LeafNodes
	LeafQosPolicyRoundRobinClassDescrIPtion  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyRoundRobinClassCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyRoundRobinClassFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyRoundRobinClassInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyRoundRobinClassQuantum      types.Number `tfsdk:"quantum" vyos:"quantum,omitempty"`
	LeafQosPolicyRoundRobinClassQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyRoundRobinClassQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyRoundRobinClassTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyRoundRobinClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyRoundRobinClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyRoundRobinClass) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRoundRobinClass) GetVyosPath() []string {
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
func (o *QosPolicyRoundRobinClass) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"round-robin",
		o.ParentIDQosPolicyRoundRobin.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyRoundRobinClass) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"round-robin",
		o.ParentIDQosPolicyRoundRobin.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClass) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format  &emsp;|  Description       |
    |----------------|--------------------|
    |  1-4095  &emsp;|  Class Identifier  |
`,
			Description: `Class ID

    |  Format  |  Description       |
    |----------------|--------------------|
    |  1-4095  |  Class Identifier  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"round_robin_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Deficit Round Robin Scheduler

    |  Format  &emsp;|  Description  |
    |----------------|---------------|
    |  txt     &emsp;|  Policy name  |
`,
			Description: `Deficit Round Robin Scheduler

    |  Format  |  Description  |
    |----------------|---------------|
    |  txt     |  Policy name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in round_robin_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  round_robin_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  &emsp;|  Description  |
    |----------------|---------------|
    |  txt     &emsp;|  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------------|---------------|
    |  txt     |  Description  |
`,
		},

		"codel_quantum": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

    |  Format     &emsp;|  Description                        |
    |-------------------|-------------------------------------|
    |  0-1048576  &emsp;|  Number of bytes used as 'deficit'  |
`,
			Description: `Deficit in the fair queuing algorithm

    |  Format     |  Description                        |
    |-------------------|-------------------------------------|
    |  0-1048576  |  Number of bytes used as 'deficit'  |
`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

    |  Format   &emsp;|  Description      |
    |-----------------|-------------------|
    |  1-65536  &emsp;|  Number of flows  |
`,
			Description: `Number of flows into which the incoming packets are classified

    |  Format   |  Description      |
    |-----------------|-------------------|
    |  1-65536  |  Number of flows  |
`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

    |  Format  &emsp;|  Description               |
    |----------------|----------------------------|
    |  u32     &emsp;|  Interval in milliseconds  |
`,
			Description: `Interval used to measure the delay

    |  Format  |  Description               |
    |----------------|----------------------------|
    |  u32     |  Interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"quantum": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Packet scheduling quantum

    |  Format        &emsp;|  Description                        |
    |----------------------|-------------------------------------|
    |  1-4294967295  &emsp;|  Packet scheduling quantum (bytes)  |
`,
			Description: `Packet scheduling quantum

    |  Format        |  Description                        |
    |----------------------|-------------------------------------|
    |  1-4294967295  |  Packet scheduling quantum (bytes)  |
`,
		},

		"queue_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format        &emsp;|  Description            |
    |----------------------|-------------------------|
    |  1-4294967295  &emsp;|  Queue size in packets  |
`,
			Description: `Maximum queue size

    |  Format        |  Description            |
    |----------------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
		},

		"queue_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

    |  Format         &emsp;|  Description                   |
    |-----------------------|--------------------------------|
    |  drop-tail      &emsp;|  First-In-First-Out (FIFO)     |
    |  fair-queue     &emsp;|  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       &emsp;|  Fair Queue Codel              |
    |  priority       &emsp;|  Priority queuing              |
    |  random-detect  &emsp;|  Random Early Detection (RED)  |
`,
			Description: `Queue type for default traffic

    |  Format         |  Description                   |
    |-----------------------|--------------------------------|
    |  drop-tail      |  First-In-First-Out (FIFO)     |
    |  fair-queue     |  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       |  Fair Queue Codel              |
    |  priority       |  Priority queuing              |
    |  random-detect  |  Random Early Detection (RED)  |
`,

			// Default:          stringdefault.StaticString(`drop-tail`),
			Computed: true,
		},

		"target": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

    |  Format  &emsp;|  Description                  |
    |----------------|-------------------------------|
    |  u32     &emsp;|  Queue delay in milliseconds  |
`,
			Description: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description                  |
    |----------------|-------------------------------|
    |  u32     |  Queue delay in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}
