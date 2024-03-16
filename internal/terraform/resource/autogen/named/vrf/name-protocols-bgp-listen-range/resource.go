// Package namedvrfnameprotocolsbgplistenrange code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgplistenrange

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-listen-range/resourcemodel"
)

// NewVrfNameProtocolsBgpListenRange method to return the example resource reference
func NewVrfNameProtocolsBgpListenRange() resource.Resource {
	return &vrfNameProtocolsBgpListenRange{
		model: &resourcemodel.VrfNameProtocolsBgpListenRange{},
	}
}

// vrfNameProtocolsBgpListenRange defines the resource implementation.
type vrfNameProtocolsBgpListenRange struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpListenRange
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpListenRange) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpListenRange) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsBgpListenRange) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsBgpListenRange) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
