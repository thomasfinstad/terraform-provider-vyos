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

// ProtocolsNhrpTunnelMap describes the resource data model.
type ProtocolsNhrpTunnelMap struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"map_id" vyos:"-,self-id"`

	ParentIDProtocolsNhrpTunnel types.String `tfsdk:"tunnel" vyos:"tunnel,parent-id"`

	// LeafNodes
	LeafProtocolsNhrpTunnelMapCisco       types.Bool   `tfsdk:"cisco" vyos:"cisco,omitempty"`
	LeafProtocolsNhrpTunnelMapNbmaAddress types.String `tfsdk:"nbma_address" vyos:"nbma-address,omitempty"`
	LeafProtocolsNhrpTunnelMapRegister    types.Bool   `tfsdk:"register" vyos:"register,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsNhrpTunnelMap) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsNhrpTunnelMap) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"nhrp",

		"tunnel",
		o.ParentIDProtocolsNhrpTunnel.ValueString(),

		"map",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsNhrpTunnelMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"map_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Set an HUB tunnel address

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"tunnel_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Tunnel for NHRP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tunN  &emsp; |  NHRP tunnel name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"cisco": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `If the statically mapped peer is running Cisco IOS, specify this

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nbma_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set HUB address (nbma-address - external hub address or fqdn)

`,
		},

		"register": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Specifies that Registration Request should be sent to this peer on startup

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
