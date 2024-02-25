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

// VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID describes the resource data model.
type VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"key_id_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsOspfInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyIDMdfiveKey types.String `tfsdk:"md5_key" vyos:"md5-key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"interface",
		o.ParentIDVrfNameProtocolsOspfInterface.ValueString(),

		"authentication",

		"md5",

		"key-id",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfInterfaceAuthenticationMdfiveKeyID) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"key_id_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `MD5 key id

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  MD5 key id  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

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

		"md5_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MD5 authentication type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  MD5 Key (16 characters or less)  |

`,
		},

		// Nodes

	}
}