// Package namedprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols-bgp-address-family-ipv4-multicast-aggregate-address/resourcemodel"
)

// NewProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress method to return the example resource reference
func NewProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress() resource.Resource {
	return &protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{
		model: resourcemodel.ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{},
	}
}

// protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress defines the resource implementation.
type protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress
}

func (r *protocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
