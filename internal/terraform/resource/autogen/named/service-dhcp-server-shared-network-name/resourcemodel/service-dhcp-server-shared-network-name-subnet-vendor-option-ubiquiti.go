// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti struct {
	// LeafNodes
	ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquitiUnifiController customtypes.CustomStringValue `tfsdk:"unifi_controller" json:"unifi-controller,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unifi_controller": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Address of UniFi controller

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of UniFi controller  |
`,
		},

		// TagNodes

		// Nodes

	}
}
