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

// LoadBalancingWanInterfaceHealth describes the resource data model.
type LoadBalancingWanInterfaceHealth struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_health_id" vyos:"-,self-id"`

	// LeafNodes
	LeafLoadBalancingWanInterfaceHealthFailureCount types.Number `tfsdk:"failure_count" vyos:"failure-count,omitempty"`
	LeafLoadBalancingWanInterfaceHealthNexthop      types.String `tfsdk:"nexthop" vyos:"nexthop,omitempty"`
	LeafLoadBalancingWanInterfaceHealthSuccessCount types.Number `tfsdk:"success_count" vyos:"success-count,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagLoadBalancingWanInterfaceHealthTest bool `tfsdk:"-" vyos:"test,ignore,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingWanInterfaceHealth) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWanInterfaceHealth) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"load-balancing",

		"wan",

		"interface-health",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanInterfaceHealth) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_health_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"failure_count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Failure count

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Failure count  |

`,
		},

		"nexthop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Outbound interface nexthop address. Can be 'DHCP or IPv4 address' [REQUIRED]

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Nexthop IP address  |
    |  dhcp  &emsp; |  Set the nexthop via DHCP  |

`,
		},

		"success_count": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Success count

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-10  &emsp; |  Success count  |

`,
		},

		// Nodes

	}
}
