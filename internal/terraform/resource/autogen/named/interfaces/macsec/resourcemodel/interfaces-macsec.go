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

// InterfacesMacsec describes the resource data model.
type InterfacesMacsec struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"macsec_id" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesMacsecAddress         types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesMacsecDescrIPtion     types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesMacsecDisable         types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesMacsecMtu             types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesMacsecSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesMacsecRedirect        types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesMacsecVrf             types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesMacsecDhcpOptions     *InterfacesMacsecDhcpOptions     `tfsdk:"dhcp_options" vyos:"dhcp-options,omitempty"`
	NodeInterfacesMacsecDhcpvsixOptions *InterfacesMacsecDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`
	NodeInterfacesMacsecIP              *InterfacesMacsecIP              `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesMacsecIPvsix          *InterfacesMacsecIPvsix          `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesMacsecMirror          *InterfacesMacsecMirror          `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesMacsecSecURIty        *InterfacesMacsecSecURIty        `tfsdk:"security" vyos:"security,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesMacsec) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesMacsec) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"macsec",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"macsec_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `MACsec Interface (802.1ae)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macsecN  &emsp; |  MACsec interface name  |

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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1460`),
			Computed: true,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical interface the traffic will go through

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  interface  &emsp; |  Physical interface used for traffic forwarding  |

`,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Security/Encryption Settings

`,
		},
	}
}
