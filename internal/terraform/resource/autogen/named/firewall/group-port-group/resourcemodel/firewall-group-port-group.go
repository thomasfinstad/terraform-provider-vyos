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

// FirewallGroupPortGroup describes the resource data model.
type FirewallGroupPortGroup struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"port_group_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallGroupPortGroupDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallGroupPortGroupPort        types.List   `tfsdk:"port" vyos:"port,omitempty"`
	LeafFirewallGroupPortGroupInclude     types.List   `tfsdk:"include" vyos:"include,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallGroupPortGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupPortGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"group",

		"port-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupPortGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"port_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall port-group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"port": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Port-group member

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Named port (any name in /etc/services, e.g., http)  |
    |  number: 1-65535  &emsp; |  Numbered port  |
    |  start-end  &emsp; |  Numbered port range (e.g. 1001-1050)  |

`,
		},

		"include": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Include another port-group

`,
		},

		// Nodes

	}
}
