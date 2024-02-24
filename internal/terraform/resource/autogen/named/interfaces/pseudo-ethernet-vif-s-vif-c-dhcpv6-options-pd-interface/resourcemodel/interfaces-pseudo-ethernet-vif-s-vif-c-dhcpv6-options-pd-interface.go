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

// InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:"-,self-id"`

	ParentIDInterfacesPseudoEthernet types.String `tfsdk:"pseudo_ethernet" vyos:"pseudo-ethernet,parent-id"`

	ParentIDInterfacesPseudoEthernetVifS types.String `tfsdk:"vif_s" vyos:"vif-s,parent-id"`

	ParentIDInterfacesPseudoEthernetVifSVifC types.String `tfsdk:"vif_c" vyos:"vif-c,parent-id"`

	ParentIDInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"pd,parent-id"`

	// LeafNodes
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"pseudo-ethernet",
		o.ParentIDInterfacesPseudoEthernet.ValueString(),

		"vif-s",
		o.ParentIDInterfacesPseudoEthernetVifS.ValueString(),

		"vif-c",
		o.ParentIDInterfacesPseudoEthernetVifSVifC.ValueString(),

		"dhcpv6-options",

		"pd",
		o.ParentIDInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd.ValueString(),

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPdInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"pseudo_ethernet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pseudo Ethernet Interface (Macvlan)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  pethN  &emsp; |  Pseudo Ethernet interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"vif_s_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"vif_c_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  >0  &emsp; |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// Nodes

	}
}
