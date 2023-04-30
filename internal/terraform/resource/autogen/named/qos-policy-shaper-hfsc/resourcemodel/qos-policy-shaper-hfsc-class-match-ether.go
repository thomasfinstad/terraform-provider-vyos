// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyShaperHfscClassMatchEther describes the resource data model.
type QosPolicyShaperHfscClassMatchEther struct {
	// LeafNodes
	QosPolicyShaperHfscClassMatchEtherDestination customtypes.CustomStringValue `tfsdk:"destination" json:"destination,omitempty"`
	QosPolicyShaperHfscClassMatchEtherProtocol    customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	QosPolicyShaperHfscClassMatchEtherSource      customtypes.CustomStringValue `tfsdk:"source" json:"source,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchEther) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"destination": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ethernet destination address for this match

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
`,
		},

		"protocol": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ethernet protocol for this match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Ethernet protocol number  |
|  txt  |  Ethernet protocol name  |
|  all  |  Any protocol  |
|  ip  |  Internet IP (IPv4)  |
|  ipv6  |  Internet IP (IPv6)  |
|  arp  |  Address Resolution Protocol  |
|  atalk  |  Appletalk  |
|  ipx  |  Novell Internet Packet Exchange  |
|  802.1Q  |  802.1Q VLAN tag  |
`,
		},

		"source": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ethernet source address for this match

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
`,
		},

		// TagNodes

		// Nodes

	}
}
