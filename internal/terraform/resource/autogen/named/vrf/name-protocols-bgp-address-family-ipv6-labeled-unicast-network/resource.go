// Package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-ipv6-labeled-unicast-network/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{},
	}
}

// vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
