// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnPptpRemoteAccessAuthenticationRadiusServer describes the resource data model.
type VpnPptpRemoteAccessAuthenticationRadiusServer struct {
	SelfIdentifier types.String `tfsdk:"server_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerKey               types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerPort              types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerAcctPort          types.Number `tfsdk:"acct_port" vyos:"acct-port,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerDisableAccounting types.Bool   `tfsdk:"disable_accounting" vyos:"disable-accounting,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationRadiusServerFailTime          types.Number `tfsdk:"fail_time" vyos:"fail-time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"vpn",

		"pptp",

		"remote-access",

		"authentication",

		"radius",

		"server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `server_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  RADIUS server IPv4 address  |

`,
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Authentication port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		"disable_accounting": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable accounting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fail_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Mark server unavailable for <n> seconds on failure

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600  &emsp; |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnPptpRemoteAccessAuthenticationRadiusServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnPptpRemoteAccessAuthenticationRadiusServer) UnmarshalJSON(_ []byte) error {
	return nil
}
