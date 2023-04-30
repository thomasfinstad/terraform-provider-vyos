// Package namedinterfaceswireguard code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswireguard

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/interfaces-wireguard/resourcemodel"
)

// NewInterfacesWireguard method to return the example resource reference
func NewInterfacesWireguard() resource.Resource {
	return &interfacesWireguard{
		model: resourcemodel.InterfacesWireguard{},
	}
}

// interfacesWireguard defines the resource implementation.
type interfacesWireguard struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.InterfacesWireguard
}

func (r *interfacesWireguard) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
