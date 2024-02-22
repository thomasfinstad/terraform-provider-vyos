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

// VpnPptpRemoteAccessAuthenticationLocalUsersUsername describes the resource data model.
type VpnPptpRemoteAccessAuthenticationLocalUsersUsername struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"username_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip" vyos:"static-ip,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"pptp",

		"remote-access",

		"authentication",

		"local-users",

		"username",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessAuthenticationLocalUsersUsername) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"username_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `User name for authentication

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for authentication

`,
		},

		"static_ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static client IP address

`,
		},

		// Nodes

	}
}
