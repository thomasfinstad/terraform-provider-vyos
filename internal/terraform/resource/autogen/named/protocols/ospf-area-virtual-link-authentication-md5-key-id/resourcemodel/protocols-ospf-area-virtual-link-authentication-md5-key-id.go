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

// ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID describes the resource data model.
type ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"key_id_id" vyos:"-,self-id"`

	ParentIDProtocolsOspfArea types.String `tfsdk:"area" vyos:"area,parent-id"`

	ParentIDProtocolsOspfAreaVirtualLink types.String `tfsdk:"virtual_link" vyos:"virtual-link,parent-id"`

	// LeafNodes
	LeafProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey types.String `tfsdk:"md5_key" vyos:"md5-key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"ospf",

		"area",
		o.ParentIDProtocolsOspfArea.ValueString(),

		"virtual-link",
		o.ParentIDProtocolsOspfAreaVirtualLink.ValueString(),

		"authentication",

		"md5",

		"key-id",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area number in decimal notation  |
    |  ipv4  &emsp; |  OSPF area number in dotted decimal notation  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"virtual_link_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual link

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  OSPF area in dotted decimal notation  |

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
