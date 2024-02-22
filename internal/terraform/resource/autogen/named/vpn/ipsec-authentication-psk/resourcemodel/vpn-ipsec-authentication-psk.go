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

// VpnIPsecAuthenticationPsk describes the resource data model.
type VpnIPsecAuthenticationPsk struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"psk_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnIPsecAuthenticationPskDhcpInterface types.List   `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafVpnIPsecAuthenticationPskID            types.List   `tfsdk:"id_param" vyos:"id,omitempty"`
	LeafVpnIPsecAuthenticationPskSecret        types.String `tfsdk:"secret" vyos:"secret,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VpnIPsecAuthenticationPsk) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecAuthenticationPsk) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"ipsec",

		"authentication",

		"psk",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecAuthenticationPsk) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"psk_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pre-shared key name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"dhcp_interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  DHCP interface name  |

`,
		},

		"id_param": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `ID for authentication

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  ID used for authentication  |

`,
		},

		"secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IKE pre-shared secret key

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  IKE pre-shared secret key  |

`,
		},

		// Nodes

	}
}
