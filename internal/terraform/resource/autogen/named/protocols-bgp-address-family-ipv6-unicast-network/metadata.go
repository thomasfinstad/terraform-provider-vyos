// Package namedprotocolsbgpaddressfamilyipvsixunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvsixunicastnetwork

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvsixUnicastNetwork) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv6_unicast_network"
	resp.TypeName = r.ResourceName
}
