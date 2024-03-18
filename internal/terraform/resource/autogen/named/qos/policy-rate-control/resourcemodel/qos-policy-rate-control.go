// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

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

var _ helpers.VyosTopResourceDataModel = &QosPolicyRateControl{}

// QosPolicyRateControl describes the resource data model.
type QosPolicyRateControl struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"rate_control_id" vyos:"-,self-id"`

	// LeafNodes
	LeafQosPolicyRateControlDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyRateControlBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyRateControlBurst       types.String `tfsdk:"burst" vyos:"burst,omitempty"`
	LeafQosPolicyRateControlLatency     types.String `tfsdk:"latency" vyos:"latency,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyRateControl) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyRateControl) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRateControl) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rate-control",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyRateControl) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyRateControl) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRateControl) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rate_control_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rate limiting policy (Token Bucket Filter)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in rate_control_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  rate_control_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"bandwidth": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |

`,
		},

		"burst": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Burst size for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bytes  |
    |  <number><suffix>  &emsp; |  Bytes with scaling suffix (kb, mb, gb)  |

`,

			// Default:          stringdefault.StaticString(`15k`),
			Computed: true,
		},

		"latency": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum latency

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Time in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`50`),
			Computed: true,
		},

		// Nodes

	}
}
