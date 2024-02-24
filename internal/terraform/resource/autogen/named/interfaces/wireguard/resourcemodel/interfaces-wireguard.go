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

// InterfacesWireguard describes the resource data model.
type InterfacesWireguard struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"wireguard_id" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesWireguardAddress     types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWireguardDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesWireguardDisable     types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWireguardPort        types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesWireguardMtu         types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesWireguardFwmark      types.String `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafInterfacesWireguardPrivateKey  types.String `tfsdk:"private_key" vyos:"private-key,omitempty"`
	LeafInterfacesWireguardRedirect    types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesWireguardVrf         types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesWireguardPeer bool `tfsdk:"-" vyos:"peer,ignore,child"`

	// Nodes
	NodeInterfacesWireguardMirror *InterfacesWireguardMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
	NodeInterfacesWireguardIP     *InterfacesWireguardIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesWireguardIPvsix *InterfacesWireguardIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesWireguard) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWireguard) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"wireguard",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWireguard) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"wireguard_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `WireGuard Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  wgN  &emsp; |  WireGuard interface name  |

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

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1420`),
			Computed: true,
		},

		"fwmark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A 32-bit fwmark value set on all outgoing packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number  &emsp; |  value which marks the packet for QoS/shaper  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"private_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Base64 encoded private key

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

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWireguardMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWireguardIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWireguardIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},
	}
}
