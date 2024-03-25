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

var _ helpers.VyosTopResourceDataModel = &QosPolicyRandomDetectPrecedence{}

// QosPolicyRandomDetectPrecedence describes the resource data model.
type QosPolicyRandomDetectPrecedence struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"precedence_id" vyos:"-,self-id"`

	ParentIDQosPolicyRandomDetect types.String `tfsdk:"random_detect_id" vyos:"random-detect,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyRandomDetectPrecedenceQueueLimit       types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyRandomDetectPrecedenceAveragePacket    types.Number `tfsdk:"average_packet" vyos:"average-packet,omitempty"`
	LeafQosPolicyRandomDetectPrecedenceMarkProbability  types.String `tfsdk:"mark_probability" vyos:"mark-probability,omitempty"`
	LeafQosPolicyRandomDetectPrecedenceMaximumThreshold types.Number `tfsdk:"maximum_threshold" vyos:"maximum-threshold,omitempty"`
	LeafQosPolicyRandomDetectPrecedenceMinimumThreshold types.Number `tfsdk:"minimum_threshold" vyos:"minimum-threshold,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyRandomDetectPrecedence) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyRandomDetectPrecedence) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyRandomDetectPrecedence) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRandomDetectPrecedence) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"precedence",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyRandomDetectPrecedence) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"random-detect",
		o.ParentIDQosPolicyRandomDetect.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyRandomDetectPrecedence) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"random-detect",
		o.ParentIDQosPolicyRandomDetect.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRandomDetectPrecedence) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"precedence_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IP precedence

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-7     |  IP precedence value  |
`,
			Description: `IP precedence

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-7     |  IP precedence value  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"random_detect_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Weighted Random Early Detect policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Weighted Random Early Detect policy

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
							"double underscores in random_detect_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  random_detect_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

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

		"average_packet": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Average packet size (bytes)

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  16-10240  |  Average packet size in bytes  |
`,
			Description: `Average packet size (bytes)

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  16-10240  |  Average packet size in bytes  |
`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"mark_probability": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mark probability for this precedence

    |  Format    |  Description          |
    |------------|-----------------------|
    |  <number>  |  Numeric value (1/N)  |
`,
			Description: `Mark probability for this precedence

    |  Format    |  Description          |
    |------------|-----------------------|
    |  <number>  |  Numeric value (1/N)  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"maximum_threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum threshold for random detection

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  0-4096  |  Maximum Threshold in packets  |
`,
			Description: `Maximum threshold for random detection

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  0-4096  |  Maximum Threshold in packets  |
`,

			// Default:          stringdefault.StaticString(`18`),
			Computed: true,
		},

		"minimum_threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum  threshold for random detection

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  0-4096  |  Maximum Threshold in packets  |
`,
			Description: `Minimum  threshold for random detection

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  0-4096  |  Maximum Threshold in packets  |
`,
		},

		// Nodes

	}
}
