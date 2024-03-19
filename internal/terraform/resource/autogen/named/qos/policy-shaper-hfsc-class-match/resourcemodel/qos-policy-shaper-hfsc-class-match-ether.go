// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyShaperHfscClassMatchEther{}

// QosPolicyShaperHfscClassMatchEther describes the resource data model.
type QosPolicyShaperHfscClassMatchEther struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchEtherDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafQosPolicyShaperHfscClassMatchEtherProtocol    types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafQosPolicyShaperHfscClassMatchEtherSource      types.String `tfsdk:"source" vyos:"source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchEther) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet destination address for this match

    |  Format   |  Description           |
    |-----------|------------------------|
    |  macaddr  |  MAC address to match  |
`,
			Description: `Ethernet destination address for this match

    |  Format   |  Description           |
    |-----------|------------------------|
    |  macaddr  |  MAC address to match  |
`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet protocol for this match

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  0-65535  |  Ethernet protocol number         |
    |  txt      |  Ethernet protocol name           |
    |  all      |  Any protocol                     |
    |  ip       |  Internet IP (IPv4)               |
    |  ipv6     |  Internet IP (IPv6)               |
    |  arp      |  Address Resolution Protocol      |
    |  atalk    |  Appletalk                        |
    |  ipx      |  Novell Internet Packet Exchange  |
    |  802.1Q   |  802.1Q VLAN tag                  |
`,
			Description: `Ethernet protocol for this match

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  0-65535  |  Ethernet protocol number         |
    |  txt      |  Ethernet protocol name           |
    |  all      |  Any protocol                     |
    |  ip       |  Internet IP (IPv4)               |
    |  ipv6     |  Internet IP (IPv6)               |
    |  arp      |  Address Resolution Protocol      |
    |  atalk    |  Appletalk                        |
    |  ipx      |  Novell Internet Packet Exchange  |
    |  802.1Q   |  802.1Q VLAN tag                  |
`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet source address for this match

    |  Format   |  Description           |
    |-----------|------------------------|
    |  macaddr  |  MAC address to match  |
`,
			Description: `Ethernet source address for this match

    |  Format   |  Description           |
    |-----------|------------------------|
    |  macaddr  |  MAC address to match  |
`,
		},

		// Nodes

	}
}
