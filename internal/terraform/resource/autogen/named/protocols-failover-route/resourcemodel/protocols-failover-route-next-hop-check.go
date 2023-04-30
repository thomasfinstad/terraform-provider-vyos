// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsFailoverRouteNextHopCheck describes the resource data model.
type ProtocolsFailoverRouteNextHopCheck struct {
	// LeafNodes
	ProtocolsFailoverRouteNextHopCheckPort    customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`
	ProtocolsFailoverRouteNextHopCheckTarget  customtypes.CustomStringValue `tfsdk:"target" json:"target,omitempty"`
	ProtocolsFailoverRouteNextHopCheckTimeout customtypes.CustomStringValue `tfsdk:"timeout" json:"timeout,omitempty"`
	ProtocolsFailoverRouteNextHopCheckType    customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHopCheck) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
`,
		},

		"target": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Check target address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to check  |
`,
		},

		"timeout": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Timeout between checks

|  Format  |  Description  |
|----------|---------------|
|  u32:1-300  |  Timeout in seconds between checks  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Check type

|  Format  |  Description  |
|----------|---------------|
|  arp  |  Check target by ARP  |
|  icmp  |  Check target by ICMP  |
|  tcp  |  Check target by TCP  |
`,

			// Default:          stringdefault.StaticString(`icmp`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
