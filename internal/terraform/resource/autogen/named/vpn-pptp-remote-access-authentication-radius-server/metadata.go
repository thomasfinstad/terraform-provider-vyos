// Package namedvpnpptpremoteaccessauthenticationradiusserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnPptpRemoteAccessAuthenticationRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_vpn_pptp_remote_access_authentication_radius_server"
	resp.TypeName = r.ResourceName
}
