// Package namedprotocolsnhrptunnelshortcuttarget code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnelshortcuttarget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsNhrpTunnelShortcutTarget) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_nhrp_tunnel_shortcut_target"
	resp.TypeName = r.ResourceName
}
