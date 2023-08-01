// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesOpenvpnServerPushRoute describes the resource data model.
type InterfacesOpenvpnServerPushRoute struct {
	SelfIdentifier types.String `tfsdk:"push_route_id" vyos:",self-id"`

	ParentIDInterfacesOpenvpn types.String `tfsdk:"openvpn" vyos:"openvpn,parent-id"`

	// LeafNodes
	LeafInterfacesOpenvpnServerPushRouteMetric types.Number `tfsdk:"metric" vyos:"metric,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesOpenvpnServerPushRoute) GetVyosPath() []string {
	return []string{
		"interfaces",

		"openvpn",
		o.ParentIDInterfacesOpenvpn.ValueString(),

		"server",

		"push-route",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerPushRoute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `push_route_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"push_route_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Route to be pushed to all clients

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 network and prefix length  |
    |  ipv6net  &emsp; |  IPv6 network and prefix length  |

`,
		},

		"openvpn_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OpenVPN Tunnel Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  vtunN  &emsp; |  OpenVPN interface name  |

`,
		},

		// LeafNodes

		"metric": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set metric for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4294967295  &emsp; |  Metric for this route  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnServerPushRoute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesOpenvpnServerPushRoute) UnmarshalJSON(_ []byte) error {
	return nil
}
