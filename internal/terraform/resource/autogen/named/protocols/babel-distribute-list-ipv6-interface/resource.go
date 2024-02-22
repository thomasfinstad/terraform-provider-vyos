// Package namedprotocolsbabeldistributelistipvsixinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbabeldistributelistipvsixinterface

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/babel-distribute-list-ipv6-interface/resourcemodel"
)

// NewProtocolsBabelDistributeListIPvsixInterface method to return the example resource reference
func NewProtocolsBabelDistributeListIPvsixInterface() resource.Resource {
	return &protocolsBabelDistributeListIPvsixInterface{
		model: &resourcemodel.ProtocolsBabelDistributeListIPvsixInterface{},
	}
}

// protocolsBabelDistributeListIPvsixInterface defines the resource implementation.
type protocolsBabelDistributeListIPvsixInterface struct {
	client *client.Client
	model  *resourcemodel.ProtocolsBabelDistributeListIPvsixInterface
}

// GetClient returns the vyos api client
func (r *protocolsBabelDistributeListIPvsixInterface) GetClient() *client.Client {
	return r.client
}

func (r *protocolsBabelDistributeListIPvsixInterface) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
