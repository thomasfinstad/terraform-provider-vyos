// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServicePppoeServerClientIPPoolName describes the resource data model.
type ServicePppoeServerClientIPPoolName struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"name_id" vyos:",self-id"`

	// LeafNodes
	LeafServicePppoeServerClientIPPoolNameGatewayAddress types.String `tfsdk:"gateway_address" vyos:"gateway-address,omitempty"`
	LeafServicePppoeServerClientIPPoolNameSubnet         types.String `tfsdk:"subnet" vyos:"subnet,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServicePppoeServerClientIPPoolName) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerClientIPPoolName) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"pppoe-server",

		"client-ip-pool",

		"name",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerClientIPPoolName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pool name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IP pool  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"gateway_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Default Gateway send to the client  |

`,
		},

		"subnet": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client IP subnet (CIDR notation)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |

`,
		},

		// Nodes

	}
}
