// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecAuthenticationPsk describes the resource data model.
type VpnIPsecAuthenticationPsk struct {
	SelfIdentifier types.String `tfsdk:"psk_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnIPsecAuthenticationPskDhcpInterface types.List   `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafVpnIPsecAuthenticationPskID            types.List   `tfsdk:"id_param" vyos:"id,omitempty"`
	LeafVpnIPsecAuthenticationPskSecret        types.String `tfsdk:"secret" vyos:"secret,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecAuthenticationPsk) GetVyosPath() []string {
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
			MarkdownDescription: "Resource ID, an amalgamation of the `psk_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"psk_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pre-shared key name

`,
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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecAuthenticationPsk) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecAuthenticationPsk) UnmarshalJSON(_ []byte) error {
	return nil
}
