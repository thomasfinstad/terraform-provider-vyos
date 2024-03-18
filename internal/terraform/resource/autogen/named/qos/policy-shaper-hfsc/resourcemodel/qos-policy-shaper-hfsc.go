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

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperHfsc{}

// QosPolicyShaperHfsc describes the resource data model.
type QosPolicyShaperHfsc struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"shaper_hfsc_id" vyos:"-,self-id"`

	// LeafNodes
	LeafQosPolicyShaperHfscDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperHfscBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyShaperHfscClass bool `tfsdk:"-" vyos:"class,child"`

	// Nodes
	NodeQosPolicyShaperHfscDefault *QosPolicyShaperHfscDefault `tfsdk:"default" vyos:"default,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyShaperHfsc) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyShaperHfsc) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperHfsc) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"shaper-hfsc",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyShaperHfsc) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyShaperHfsc) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfsc) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"shaper_hfsc_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Hierarchical Fair Service Curve's policy

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
							"double underscores in shaper_hfsc_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  shaper_hfsc_id, value must match: ^[a-zA-Z0-9-_]*$",
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
    |  auto  &emsp; |  Bandwidth matches interface speed  |
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |

`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		// Nodes

		"default": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefault{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default policy

`,
		},
	}
}
