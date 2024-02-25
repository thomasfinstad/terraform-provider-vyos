// Package namedcontainernetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/container/network/resourcemodel"
)

// NewContainerNetwork method to return the example resource reference
func NewContainerNetwork() resource.Resource {
	return &containerNetwork{
		model: &resourcemodel.ContainerNetwork{},
	}
}

// containerNetwork defines the resource implementation.
type containerNetwork struct {
	client *client.Client
	model  *resourcemodel.ContainerNetwork
}

// GetClient returns the vyos api client
func (r *containerNetwork) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *containerNetwork) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *containerNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
