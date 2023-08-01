// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperClassMatch describes the resource data model.
type QosPolicyShaperClassMatch struct {
	SelfIdentifier types.String `tfsdk:"match_id" vyos:",self-id"`

	ParentIDQosPolicyShaper types.String `tfsdk:"shaper" vyos:"shaper,parent-id"`

	ParentIDQosPolicyShaperClass types.String `tfsdk:"class" vyos:"class,parent-id"`

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

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperClassMatch) GetVyosPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper",
		o.ParentIDQosPolicyShaper.ValueString(),

		"class",
		o.ParentIDQosPolicyShaperClass.ValueString(),

		"match",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `match_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"match_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class matching rule name

`,
		},

		"shaper_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Traffic shaping based policy (Hierarchy Token Bucket)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
		},

		"class_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-4095  &emsp; |  Class Identifier  |

`,
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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperClassMatch) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperClassMatch) UnmarshalJSON(_ []byte) error {
	return nil
}
