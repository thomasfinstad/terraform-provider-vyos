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

// ServiceIPoeServerClientIPvsixPoolPrefix describes the resource data model.
type ServiceIPoeServerClientIPvsixPoolPrefix struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceIPoeServerClientIPvsixPoolPrefixMask types.Number `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceIPoeServerClientIPvsixPoolPrefix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerClientIPvsixPoolPrefix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"ipoe-server",

		"client-ipv6-pool",

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerClientIPvsixPoolPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pool of addresses used to assign to clients

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"mask": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 48-128  &emsp; |  Client prefix length  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}
