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

// ServiceDhcpServerSharedNetworkNameSubnetStaticMapping describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetStaticMapping struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"static_mapping_id" vyos:",self-id"`

	ParentIDServiceDhcpServerSharedNetworkName types.String `tfsdk:"shared_network_name" vyos:"shared-network-name,parent-id"`

	ParentIDServiceDhcpServerSharedNetworkNameSubnet types.String `tfsdk:"subnet" vyos:"subnet,parent-id"`

	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable                 types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress               types.String `tfsdk:"ip_address" vyos:"ip-address,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMacAddress              types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingStaticMappingParameters types.List   `tfsdk:"static_mapping_parameters" vyos:"static-mapping-parameters,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"dhcp-server",

		"shared-network-name",
		o.ParentIDServiceDhcpServerSharedNetworkName.ValueString(),

		"subnet",
		o.ParentIDServiceDhcpServerSharedNetworkNameSubnet.ValueString(),

		"static-mapping",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"static_mapping_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of static mapping

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"shared_network_name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of DHCP shared network

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"subnet_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `DHCP subnet for shared network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ip_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fixed IP address of static mapping

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address used in static mapping  |

`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		"static_mapping_parameters": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Additional static-mapping parameters for DHCP server. Will be placed inside the "host" block of the mapping. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		// Nodes

	}
}
