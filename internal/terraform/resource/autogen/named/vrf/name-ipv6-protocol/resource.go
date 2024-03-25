// Package namedvrfnameipvsixprotocol code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameipvsixprotocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-ipv6-protocol/resourcemodel"
)

// NewVrfNameIPvsixProtocol method to return the example resource reference
func NewVrfNameIPvsixProtocol() resource.Resource {
	return &vrfNameIPvsixProtocol{
		model: &resourcemodel.VrfNameIPvsixProtocol{},
	}
}

// vrfNameIPvsixProtocol defines the resource implementation.
type vrfNameIPvsixProtocol struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameIPvsixProtocol
}

// GetClient returns the vyos api client
func (r *vrfNameIPvsixProtocol) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameIPvsixProtocol) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameIPvsixProtocol) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameIPvsixProtocol) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
