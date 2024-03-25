// Package namedvrfnameprotocolsbgppeergrouplocalrole code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgppeergrouplocalrole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-peer-group-local-role/resourcemodel"
)

// NewVrfNameProtocolsBgpPeerGroupLocalRole method to return the example resource reference
func NewVrfNameProtocolsBgpPeerGroupLocalRole() resource.Resource {
	return &vrfNameProtocolsBgpPeerGroupLocalRole{
		model: &resourcemodel.VrfNameProtocolsBgpPeerGroupLocalRole{},
	}
}

// vrfNameProtocolsBgpPeerGroupLocalRole defines the resource implementation.
type vrfNameProtocolsBgpPeerGroupLocalRole struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpPeerGroupLocalRole
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpPeerGroupLocalRole) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpPeerGroupLocalRole) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsBgpPeerGroupLocalRole) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsBgpPeerGroupLocalRole) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
