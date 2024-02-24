// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnIPsec describes the resource data model.
type VpnIPsec struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnIPsecDisableUniqreqIDs types.Bool `tfsdk:"disable_uniqreqids" vyos:"disable-uniqreqids,omitempty"`
	LeafVpnIPsecInterface         types.List `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVpnIPsecEspGroup bool `tfsdk:"-" vyos:"esp-group,ignore,child"`
	ExistsTagVpnIPsecIkeGroup bool `tfsdk:"-" vyos:"ike-group,ignore,child"`
	ExistsTagVpnIPsecProfile  bool `tfsdk:"-" vyos:"profile,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnIPsecAuthentication bool `tfsdk:"-" vyos:"authentication,ignore,omitempty"`
	ExistsNodeVpnIPsecLog            bool `tfsdk:"-" vyos:"log,ignore,omitempty"`
	ExistsNodeVpnIPsecOptions        bool `tfsdk:"-" vyos:"options,ignore,omitempty"`
	ExistsNodeVpnIPsecRemoteAccess   bool `tfsdk:"-" vyos:"remote-access,ignore,omitempty"`
	ExistsNodeVpnIPsecSiteToSite     bool `tfsdk:"-" vyos:"site-to-site,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *VpnIPsec) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsec) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"ipsec",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"disable_uniqreqids": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable requirement for unique IDs in the Security Database

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},
	}
}
