// Package namedservicepppoeserverauthenticationradiusserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationradiusserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r servicePppoeServerAuthenticationRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_service_pppoe_server_authentication_radius_server"
	resp.TypeName = r.ResourceName
}
