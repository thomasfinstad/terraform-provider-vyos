// Package namedvpnopenconnectauthenticationlocalusersusername code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectauthenticationlocalusersusername

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnOpenconnectAuthenticationLocalUsersUsername) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_openconnect_authentication_local_users_username"
}
