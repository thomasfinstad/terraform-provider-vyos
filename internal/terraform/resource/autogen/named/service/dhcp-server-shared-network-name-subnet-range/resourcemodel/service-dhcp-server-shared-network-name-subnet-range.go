// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDhcpServerSharedNetworkNameSubnetRange describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetRange struct {
	SelfIdentifier types.String `tfsdk:"range_id" vyos:",self-id"`

	ParentIDServiceDhcpServerSharedNetworkName types.String `tfsdk:"shared_network_name" vyos:"shared-network-name,parent-id"`

	ParentIDServiceDhcpServerSharedNetworkNameSubnet types.String `tfsdk:"subnet" vyos:"subnet,parent-id"`

	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetRangeStart types.String `tfsdk:"start" vyos:"start,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetRangeStop  types.String `tfsdk:"stop" vyos:"stop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) GetVyosPath() []string {
	return []string{
		"service",

		"dhcp-server",

		"shared-network-name",
		o.ParentIDServiceDhcpServerSharedNetworkName.ValueString(),

		"subnet",
		o.ParentIDServiceDhcpServerSharedNetworkNameSubnet.ValueString(),

		"range",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetRange) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `range_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"range_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCP lease range

`,
		},

		"shared_network_name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of DHCP shared network

`,
		},

		"subnet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCP subnet for shared network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |

`,
		},

		// LeafNodes

		"start": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `First IP address for DHCP lease range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 start address of pool  |

`,
		},

		"stop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Last IP address for DHCP lease range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 end address of pool  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDhcpServerSharedNetworkNameSubnetRange) UnmarshalJSON(_ []byte) error {
	return nil
}
