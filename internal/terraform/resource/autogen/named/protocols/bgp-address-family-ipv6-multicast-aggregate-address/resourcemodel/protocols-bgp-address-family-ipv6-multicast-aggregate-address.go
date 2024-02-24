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

// ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress describes the resource data model.
type ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"aggregate_address_id" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressAsSet       types.Bool   `tfsdk:"as_set" vyos:"as-set,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressRouteMap    types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddressSummaryOnly types.Bool   `tfsdk:"summary_only" vyos:"summary-only,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"address-family",

		"ipv6-multicast",

		"aggregate-address",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"aggregate_address_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP aggregate network/prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  BGP aggregate network/prefix  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"as_set": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		"summary_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Announce the aggregate summary network only

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
