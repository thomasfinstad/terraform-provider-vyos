// Package namedprotocolsstaticmulticastroutenexthop code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticmulticastroutenexthop

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/static-multicast-route-next-hop/resourcemodel"
)

// NewProtocolsStaticMulticastRouteNextHop method to return the example resource reference
func NewProtocolsStaticMulticastRouteNextHop() resource.Resource {
	return &protocolsStaticMulticastRouteNextHop{
		model: resourcemodel.ProtocolsStaticMulticastRouteNextHop{},
	}
}

// protocolsStaticMulticastRouteNextHop defines the resource implementation.
type protocolsStaticMulticastRouteNextHop struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ProtocolsStaticMulticastRouteNextHop
}

func (r *protocolsStaticMulticastRouteNextHop) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
