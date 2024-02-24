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

// InterfacesBrIDgeMemberInterface describes the resource data model.
type InterfacesBrIDgeMemberInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:"-,self-id"`

	ParentIDInterfacesBrIDge types.String `tfsdk:"bridge" vyos:"bridge,parent-id"`

	// LeafNodes
	LeafInterfacesBrIDgeMemberInterfaceNativeVlan  types.Number `tfsdk:"native_vlan" vyos:"native-vlan,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceAllowedVlan types.List   `tfsdk:"allowed_vlan" vyos:"allowed-vlan,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceCost        types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafInterfacesBrIDgeMemberInterfacePriority    types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafInterfacesBrIDgeMemberInterfaceIsolated    types.Bool   `tfsdk:"isolated" vyos:"isolated,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesBrIDgeMemberInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBrIDgeMemberInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"interfaces",

		"bridge",
		o.ParentIDInterfacesBrIDge.ValueString(),

		"member",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeMemberInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Member interface name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"bridge_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bridge Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  brN  &emsp; |  Bridge interface name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"native_vlan": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specify VLAN id which should natively be present on the link

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4094  &emsp; |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		"allowed_vlan": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify VLAN id which is allowed in this trunk interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <id>  &emsp; |  VLAN id allowed to pass this interface  |
    |  <idN>-<idM>  &emsp; |  VLAN id range allowed on this interface (use '-' as delimiter)  |

`,
		},

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Bridge port cost

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Path cost value for Spanning Tree Protocol  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Bridge port priority

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  Bridge port priority  |

`,

			// Default:          stringdefault.StaticString(`32`),
			Computed: true,
		},

		"isolated": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Port is isolated (also known as Private-VLAN)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
