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

// ProtocolsFailoverRouteNextHop describes the resource data model.
type ProtocolsFailoverRouteNextHop struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"next_hop_id" vyos:"-,self-id"`

	ParentIDProtocolsFailoverRoute types.String `tfsdk:"route" vyos:"route,parent-id"`

	// LeafNodes
	LeafProtocolsFailoverRouteNextHopInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafProtocolsFailoverRouteNextHopMetric    types.Number `tfsdk:"metric" vyos:"metric,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsFailoverRouteNextHopCheck *ProtocolsFailoverRouteNextHopCheck `tfsdk:"check" vyos:"check,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsFailoverRouteNextHop) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsFailoverRouteNextHop) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"failover",

		"route",
		o.ParentIDProtocolsFailoverRoute.ValueString(),

		"next-hop",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHop) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"next_hop_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Next-hop IPv4 router address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Next-hop router address  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Failover IPv4 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 failover route  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway interface name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Gateway interface name  |

`,
		},

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Route metric for this gateway

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Route metric  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// Nodes

		"check": schema.SingleNestedAttribute{
			Attributes: ProtocolsFailoverRouteNextHopCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Check target options

`,
		},
	}
}
