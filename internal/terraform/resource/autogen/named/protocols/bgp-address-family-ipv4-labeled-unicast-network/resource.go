// Package namedprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/bgp-address-family-ipv4-labeled-unicast-network/resourcemodel"
)

// NewProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork method to return the example resource reference
func NewProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork() resource.Resource {
	return &protocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{
		model: &resourcemodel.ProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{},
	}
}

// protocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork defines the resource implementation.
type protocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork struct {
	client *client.Client
	model  *resourcemodel.ProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork
}

// GetClient returns the vyos api client
func (r *protocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) GetClient() *client.Client {
	return r.client
}

func (r *protocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
