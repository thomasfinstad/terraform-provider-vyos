// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesVirtualEthernetDhcpvsixOptionsPd describes the resource data model.
type InterfacesVirtualEthernetDhcpvsixOptionsPd struct {
	SelfIdentifier types.String `tfsdk:"pd_id" vyos:",self-id"`

	ParentIDInterfacesVirtualEthernet types.String `tfsdk:"virtual_ethernet" vyos:"virtual-ethernet,parent-id"`

	// LeafNodes
	LeafInterfacesVirtualEthernetDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesVirtualEthernetDhcpvsixOptionsPdInterface bool `tfsdk:"interface" vyos:"interface,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesVirtualEthernetDhcpvsixOptionsPd) GetVyosPath() []string {
	return []string{
		"interfaces",

		"virtual-ethernet",
		o.ParentIDInterfacesVirtualEthernet.ValueString(),

		"dhcpv6-options",

		"pd",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVirtualEthernetDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `pd_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"pd_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |

`,
		},

		"virtual_ethernet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Ethernet (veth) Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  vethN  &emsp; |  Virtual Ethernet interface name  |

`,
		},

		// LeafNodes

		"length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 32-64  &emsp; |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesVirtualEthernetDhcpvsixOptionsPd) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesVirtualEthernetDhcpvsixOptionsPd) UnmarshalJSON(_ []byte) error {
	return nil
}
