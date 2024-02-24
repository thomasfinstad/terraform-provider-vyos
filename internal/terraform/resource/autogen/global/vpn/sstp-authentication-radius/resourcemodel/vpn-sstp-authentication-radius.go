// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnSstpAuthenticationRadius describes the resource data model.
type VpnSstpAuthenticationRadius struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnSstpAuthenticationRadiusSourceAddress     types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafVpnSstpAuthenticationRadiusAcctInterimJitter types.Number `tfsdk:"acct_interim_jitter" vyos:"acct-interim-jitter,omitempty"`
	LeafVpnSstpAuthenticationRadiusTimeout           types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafVpnSstpAuthenticationRadiusAcctTimeout       types.Number `tfsdk:"acct_timeout" vyos:"acct-timeout,omitempty"`
	LeafVpnSstpAuthenticationRadiusMaxTry            types.Number `tfsdk:"max_try" vyos:"max-try,omitempty"`
	LeafVpnSstpAuthenticationRadiusNasIDentifier     types.String `tfsdk:"nas_identifier" vyos:"nas-identifier,omitempty"`
	LeafVpnSstpAuthenticationRadiusNasIPAddress      types.String `tfsdk:"nas_ip_address" vyos:"nas-ip-address,omitempty"`
	LeafVpnSstpAuthenticationRadiusPreallocateVif    types.Bool   `tfsdk:"preallocate_vif" vyos:"preallocate-vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVpnSstpAuthenticationRadiusServer bool `tfsdk:"-" vyos:"server,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnSstpAuthenticationRadiusDynamicAuthor bool `tfsdk:"-" vyos:"dynamic-author,ignore,omitempty"`
	ExistsNodeVpnSstpAuthenticationRadiusRateLimit     bool `tfsdk:"-" vyos:"rate-limit,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *VpnSstpAuthenticationRadius) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpAuthenticationRadius) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"sstp",

		"authentication",

		"radius",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthenticationRadius) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 source address used to initiate connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 source address  |

`,
		},

		"acct_interim_jitter": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum jitter value in seconds to be applied to accounting information interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-60  &emsp; |  Maximum jitter value in seconds  |

`,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout in seconds to wait response from RADIUS server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-60  &emsp; |  Timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"acct_timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for Interim-Update packets, terminate session afterwards

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60  &emsp; |  Timeout in seconds, 0 to keep active  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"max_try": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of tries to send Access-Request/Accounting-Request queries

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-20  &emsp; |  Maximum tries  |

`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"nas_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-Identifier attribute sent to RADIUS

`,
		},

		"nas_ip_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-IP-Address attribute sent to RADIUS

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  NAS-IP-Address attribute  |

`,
		},

		"preallocate_vif": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable attribute NAS-Port-Id in Access-Request

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
