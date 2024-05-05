// Package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-ipv6-multicast-aggregate-address/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress{},
	}
}

// vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(data.ProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.providerData = data
}
