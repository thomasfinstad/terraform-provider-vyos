// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsStaticRoutesixNextHop describes the resource data model.
type VrfNameProtocolsStaticRoutesixNextHop struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"next_hop_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsStaticRoutesix types.String `tfsdk:"route6" vyos:"route6,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsStaticRoutesixNextHopDisable   types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixNextHopDistance  types.Number `tfsdk:"distance" vyos:"distance,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixNextHopInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafVrfNameProtocolsStaticRoutesixNextHopVrf       types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsStaticRoutesixNextHop) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsStaticRoutesixNextHop) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"static",

		"route6",
		o.ParentIDVrfNameProtocolsStaticRoutesix.ValueString(),

		"next-hop",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsStaticRoutesixNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"next_hop_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 gateway address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Next-hop IPv6 router  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for this route  |

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Gateway interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF to leak route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of VRF to leak to  |

`,
		},

		// Nodes

	}
}
