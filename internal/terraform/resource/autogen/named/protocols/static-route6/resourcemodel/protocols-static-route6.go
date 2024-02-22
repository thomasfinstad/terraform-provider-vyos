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

// ProtocolsStaticRoutesix describes the resource data model.
type ProtocolsStaticRoutesix struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"route6_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsStaticRoutesixDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsStaticRoutesixInterface bool `tfsdk:"-" vyos:"interface,child"`
	ExistsTagProtocolsStaticRoutesixNextHop   bool `tfsdk:"-" vyos:"next-hop,child"`

	// Nodes
	NodeProtocolsStaticRoutesixBlackhole *ProtocolsStaticRoutesixBlackhole `tfsdk:"blackhole" vyos:"blackhole,omitempty"`
	NodeProtocolsStaticRoutesixReject    *ProtocolsStaticRoutesixReject    `tfsdk:"reject" vyos:"reject,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsStaticRoutesix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticRoutesix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"static",

		"route6",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoutesix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"route6_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv6 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 static route  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
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

		// Nodes

		"blackhole": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRoutesixBlackhole{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Silently discard pkts when matched

`,
		},

		"reject": schema.SingleNestedAttribute{
			Attributes: ProtocolsStaticRoutesixReject{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Emit an ICMP unreachable when matched

`,
		},
	}
}
