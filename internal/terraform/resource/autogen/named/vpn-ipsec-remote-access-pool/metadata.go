// Package namedvpnipsecremoteaccesspool code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccesspool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r vpnIPsecRemoteAccessPool) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_vpn_ipsec_remote_access_pool"
	resp.TypeName = r.ResourceName
}
