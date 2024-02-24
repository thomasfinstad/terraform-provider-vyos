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

// ProtocolsStaticArpInterfaceAddress describes the resource data model.
type ProtocolsStaticArpInterfaceAddress struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"address_id" vyos:"-,self-id"`

	ParentIDProtocolsStaticArpInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafProtocolsStaticArpInterfaceAddressDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafProtocolsStaticArpInterfaceAddressMac         types.String `tfsdk:"mac" vyos:"mac,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsStaticArpInterfaceAddress) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsStaticArpInterfaceAddress) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"static",

		"arp",

		"interface",
		o.ParentIDProtocolsStaticArpInterface.ValueString(),

		"address",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticArpInterfaceAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"address_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP address for static ARP entry

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 destination address  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

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

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  Hardware (MAC) address  |

`,
		},

		// Nodes

	}
}
