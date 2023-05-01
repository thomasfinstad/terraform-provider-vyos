// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnPptpRemoteAccessAuthenticationLocalUsersUsername describes the resource data model.
type VpnPptpRemoteAccessAuthenticationLocalUsersUsername struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable  types.String `tfsdk:"disable"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip"`

	// TagNodes

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) GetVyosPath() []string {
	return []string{
		"vpn",
		"pptp",
		"remote-access",
		"authentication",
		"local-users",
		"username",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vpn", "pptp", "remote-access", "authentication", "local-users", "username"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable.IsNull() || o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable.ValueString()
	}
	if !(o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword.IsNull() || o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword.IsUnknown()) {
		vyosData["password"] = o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword.ValueString()
	}
	if !(o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP.IsNull() || o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP.IsUnknown()) {
		vyosData["static-ip"] = o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP.ValueString()
	}

	// Tags

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vpn", "pptp", "remote-access", "authentication", "local-users", "username"}})

	// Leafs
	if value, ok := vyosData["disable"]; ok {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["password"]; ok {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword = basetypes.NewStringNull()
	}
	if value, ok := vyosData["static-ip"]; ok {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP = basetypes.NewStringNull()
	}

	// Tags

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vpn", "pptp", "remote-access", "authentication", "local-users", "username"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VpnPptpRemoteAccessAuthenticationLocalUsersUsername) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"disable":   types.StringType,
		"password":  types.StringType,
		"static_ip": types.StringType,

		// Tags

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessAuthenticationLocalUsersUsername) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `User name for authentication

`,
		},

		// LeafNodes

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
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

		// TagNodes

		// Nodes

	}
}
