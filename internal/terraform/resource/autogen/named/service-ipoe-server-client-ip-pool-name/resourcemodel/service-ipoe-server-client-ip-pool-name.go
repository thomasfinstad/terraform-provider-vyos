// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceIPoeServerClientIPPoolName describes the resource data model.
type ServiceIPoeServerClientIPPoolName struct {
	// LeafNodes
	ServiceIPoeServerClientIPPoolNameGatewayAddress customtypes.CustomStringValue `tfsdk:"gateway_address" json:"gateway-address,omitempty"`
	ServiceIPoeServerClientIPPoolNameSubnet         customtypes.CustomStringValue `tfsdk:"subnet" json:"subnet,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceIPoeServerClientIPPoolName) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"gateway_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Gateway IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default Gateway send to the client  |
`,
		},

		"subnet": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client IP subnet (CIDR notation)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
`,
		},

		// TagNodes

		// Nodes

	}
}
