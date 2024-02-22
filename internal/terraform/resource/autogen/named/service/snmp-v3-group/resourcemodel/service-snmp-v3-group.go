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

// ServiceSnmpVthreeGroup describes the resource data model.
type ServiceSnmpVthreeGroup struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"group_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceSnmpVthreeGroupMode     types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafServiceSnmpVthreeGroupSeclevel types.String `tfsdk:"seclevel" vyos:"seclevel,omitempty"`
	LeafServiceSnmpVthreeGroupView     types.String `tfsdk:"view" vyos:"view,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSnmpVthreeGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpVthreeGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"snmp",

		"v3",

		"group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Specifies the group with name groupname

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define access permission

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ro  &emsp; |  Read-Only  |
    |  rw  &emsp; |  read write  |

`,

			// Default:          stringdefault.StaticString(`ro`),
			Computed: true,
		},

		"seclevel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Security levels

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  noauth  &emsp; |  Messages not authenticated and not encrypted (noAuthNoPriv)  |
    |  auth  &emsp; |  Messages are authenticated but not encrypted (authNoPriv)  |
    |  priv  &emsp; |  Messages are authenticated and encrypted (authPriv)  |

`,

			// Default:          stringdefault.StaticString(`auth`),
			Computed: true,
		},

		"view": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the name of view

`,
		},

		// Nodes

	}
}
