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

// VpnIPsecRemoteAccessPool describes the resource data model.
type VpnIPsecRemoteAccessPool struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"pool_id" vyos:"-,self-id"`

	// LeafNodes
	LeafVpnIPsecRemoteAccessPoolExclude    types.List   `tfsdk:"exclude" vyos:"exclude,omitempty"`
	LeafVpnIPsecRemoteAccessPoolPrefix     types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`
	LeafVpnIPsecRemoteAccessPoolNameServer types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VpnIPsecRemoteAccessPool) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecRemoteAccessPool) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"ipsec",

		"remote-access",

		"pool",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessPool) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"pool_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IP address pool for remote access users

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IPv4 or IPv6 pool prefix exclusions

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Local IPv4 pool prefix exclusion  |
    |  ipv6net  &emsp; |  Local IPv6 pool prefix exclusion  |

`,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IPv4 or IPv6 pool prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Local IPv4 pool prefix  |
    |  ipv6net  &emsp; |  Local IPv6 pool prefix  |

`,
		},

		"name_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6  &emsp; |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		// Nodes

	}
}
