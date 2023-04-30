// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceIPoeServerInterfaceExternalDhcp describes the resource data model.
type ServiceIPoeServerInterfaceExternalDhcp struct {
	// LeafNodes
	ServiceIPoeServerInterfaceExternalDhcpDhcpRelay customtypes.CustomStringValue `tfsdk:"dhcp_relay" json:"dhcp-relay,omitempty"`
	ServiceIPoeServerInterfaceExternalDhcpGiaddr    customtypes.CustomStringValue `tfsdk:"giaddr" json:"giaddr,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceIPoeServerInterfaceExternalDhcp) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dhcp_relay": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DHCP Server the request will be redirected to.

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of the DHCP Server  |
`,
		},

		"giaddr": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Relay Agent IPv4 Address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Gateway IP address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
