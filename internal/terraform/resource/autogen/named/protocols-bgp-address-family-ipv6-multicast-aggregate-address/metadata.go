// Package namedprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r protocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.ResourceName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv6_multicast_aggregate_address"
	resp.TypeName = r.ResourceName
}
