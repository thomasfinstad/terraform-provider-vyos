// Package namedvrfnameprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourmulticastaggregateaddress

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-ipv4-multicast-aggregate-address/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress{},
	}
}

// vrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress struct {
	client *client.Client
	model  *resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) GetClient() *client.Client {
	return r.client
}

func (r *vrfNameProtocolsBgpAddressFamilyIPvfourMulticastAggregateAddress) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
