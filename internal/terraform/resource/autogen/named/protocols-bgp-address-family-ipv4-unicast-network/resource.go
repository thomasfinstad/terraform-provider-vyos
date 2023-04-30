// Package namedprotocolsbgpaddressfamilyipvfourunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvfourunicastnetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols-bgp-address-family-ipv4-unicast-network/resourcemodel"
)

// NewProtocolsBgpAddressFamilyIPvfourUnicastNetwork method to return the example resource reference
func NewProtocolsBgpAddressFamilyIPvfourUnicastNetwork() resource.Resource {
	return &protocolsBgpAddressFamilyIPvfourUnicastNetwork{
		model: resourcemodel.ProtocolsBgpAddressFamilyIPvfourUnicastNetwork{},
	}
}

// protocolsBgpAddressFamilyIPvfourUnicastNetwork defines the resource implementation.
type protocolsBgpAddressFamilyIPvfourUnicastNetwork struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ProtocolsBgpAddressFamilyIPvfourUnicastNetwork
}

func (r *protocolsBgpAddressFamilyIPvfourUnicastNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
