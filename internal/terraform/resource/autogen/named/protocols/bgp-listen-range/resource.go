// Package namedprotocolsbgplistenrange code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgplistenrange

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/bgp-listen-range/resourcemodel"
)

// NewProtocolsBgpListenRange method to return the example resource reference
func NewProtocolsBgpListenRange() resource.Resource {
	return &protocolsBgpListenRange{
		model: &resourcemodel.ProtocolsBgpListenRange{},
	}
}

// protocolsBgpListenRange defines the resource implementation.
type protocolsBgpListenRange struct {
	client *client.Client
	model  *resourcemodel.ProtocolsBgpListenRange
}

// GetClient returns the vyos api client
func (r *protocolsBgpListenRange) GetClient() *client.Client {
	return r.client
}

func (r *protocolsBgpListenRange) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
