// Package namedprotocolsnhrptunnel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsNhrpTunnel) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_nhrp_tunnel"
	resp.TypeName = r.ResourceName
}
