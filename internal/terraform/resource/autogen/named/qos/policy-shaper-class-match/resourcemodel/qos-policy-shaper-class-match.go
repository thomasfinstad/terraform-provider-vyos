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

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperClassMatch{}

// QosPolicyShaperClassMatch describes the resource data model.
type QosPolicyShaperClassMatch struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"match_id" vyos:"-,self-id"`

	ParentIDQosPolicyShaper types.String `tfsdk:"shaper_id" vyos:"shaper,parent-id"`

	ParentIDQosPolicyShaperClass types.Number `tfsdk:"class_id" vyos:"class,parent-id"`

	// LeafNodes
	LeafQosPolicyShaperClassMatchDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperClassMatchInterface   types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafQosPolicyShaperClassMatchMark        types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafQosPolicyShaperClassMatchVif         types.Number `tfsdk:"vif" vyos:"vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyShaperClassMatchEther  *QosPolicyShaperClassMatchEther  `tfsdk:"ether" vyos:"ether,omitempty"`
	NodeQosPolicyShaperClassMatchIP     *QosPolicyShaperClassMatchIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeQosPolicyShaperClassMatchIPvsix *QosPolicyShaperClassMatchIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyShaperClassMatch) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperClassMatch) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"qos",

		"policy",

		"shaper",
		o.ParentIDQosPolicyShaper.ValueString(),

		"class",
		o.ParentIDQosPolicyShaperClass.ValueBigFloat().String(),

		"match",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"match_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class matching rule name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in match_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  match_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"shaper_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic shaping based policy (Hierarchy Token Bucket)

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
							"double underscores in shaper_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  shaper_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-4095  &emsp; |  Class Identifier  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
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

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		"mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on mark applied by firewall

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  FW mark to match  |

`,
		},

		"vif": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4095  &emsp; |  Virtual Local Area Network (VLAN) tag   |

`,
		},

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}
