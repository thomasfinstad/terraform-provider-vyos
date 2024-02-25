// Package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-isis-fast-reroute-lfa-local-tiebreaker-node-protecting-index/resourcemodel"
)

// NewVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex method to return the example resource reference
func NewVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex() resource.Resource {
	return &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{
		model: &resourcemodel.VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{},
	}
}

// vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex defines the resource implementation.
type vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex struct {
	client *client.Client
	model  *resourcemodel.VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
