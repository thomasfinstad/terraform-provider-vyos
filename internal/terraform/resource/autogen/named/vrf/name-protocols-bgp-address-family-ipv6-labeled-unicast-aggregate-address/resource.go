// Package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastaggregateaddress code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastaggregateaddress

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-ipv6-labeled-unicast-aggregate-address/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress{},
	}
}

// vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastAggregateAddress) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
