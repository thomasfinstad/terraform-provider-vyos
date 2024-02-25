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

// InterfacesPseudoEthernet describes the resource data model.
type InterfacesPseudoEthernet struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"pseudo_ethernet_id" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesPseudoEthernetAddress           types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesPseudoEthernetDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesPseudoEthernetDisableLinkDetect types.Bool   `tfsdk:"disable_link_detect" vyos:"disable-link-detect,omitempty"`
	LeafInterfacesPseudoEthernetDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesPseudoEthernetVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesPseudoEthernetSourceInterface   types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesPseudoEthernetMac               types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesPseudoEthernetMode              types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesPseudoEthernetMtu               types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesPseudoEthernetRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesPseudoEthernetVifS bool `tfsdk:"-" vyos:"vif-s,child"`
	ExistsTagInterfacesPseudoEthernetVif  bool `tfsdk:"-" vyos:"vif,child"`

	// Nodes
	NodeInterfacesPseudoEthernetDhcpOptions     *InterfacesPseudoEthernetDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesPseudoEthernetDhcpvsixOptions *InterfacesPseudoEthernetDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesPseudoEthernetIP              *InterfacesPseudoEthernetIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesPseudoEthernetIPvsix          *InterfacesPseudoEthernetIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesPseudoEthernetMirror          *InterfacesPseudoEthernetMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesPseudoEthernet) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPseudoEthernet) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"pseudo-ethernet",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
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

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |
    |  dhcp  &emsp; |  Dynamic Host Configuration Protocol  |
    |  dhcpv6  &emsp; |  Dynamic Host Configuration Protocol for IPv6  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable_link_detect": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical interface the traffic will go through

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  interface  &emsp; |  Physical interface used for traffic forwarding  |

`,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Receive mode (default: private)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  private  &emsp; |  No communication with other pseudo-devices  |
    |  vepa  &emsp; |  Virtual Ethernet Port Aggregator reflective relay  |
    |  bridge  &emsp; |  Simple bridge between pseudo-devices  |
    |  passthru  &emsp; |  Promicious mode passthrough of underlying device  |

`,

			// Default:          stringdefault.StaticString(`private`),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPseudoEthernetMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}
