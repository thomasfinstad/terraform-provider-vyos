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

// ServiceIPoeServerClientIPvsixPoolDelegate describes the resource data model.
type ServiceIPoeServerClientIPvsixPoolDelegate struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"delegate_id" vyos:"-,self-id"`

	// LeafNodes
	LeafServiceIPoeServerClientIPvsixPoolDelegateDelegationPrefix types.Number `tfsdk:"delegation_prefix" vyos:"delegation-prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceIPoeServerClientIPvsixPoolDelegate) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerClientIPvsixPoolDelegate) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"ipoe-server",

		"client-ipv6-pool",

		"delegate",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerClientIPvsixPoolDelegate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"delegate_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Subnet used to delegate prefix through DHCPv6-PD (RFC3633)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"delegation_prefix": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length delegated to client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 32-64  &emsp; |  Delegated prefix length  |

`,
		},

		// Nodes

	}
}
