// Package namedvrfnameprotocolsospfarearange code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfarearange

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-ospf-area-range/resourcemodel"
)

// NewVrfNameProtocolsOspfAreaRange method to return the example resource reference
func NewVrfNameProtocolsOspfAreaRange() resource.Resource {
	return &vrfNameProtocolsOspfAreaRange{
		model: &resourcemodel.VrfNameProtocolsOspfAreaRange{},
	}
}

// vrfNameProtocolsOspfAreaRange defines the resource implementation.
type vrfNameProtocolsOspfAreaRange struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsOspfAreaRange
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsOspfAreaRange) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsOspfAreaRange) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsOspfAreaRange) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsOspfAreaRange) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
