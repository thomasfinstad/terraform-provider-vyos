// Package namedprotocolsospfaccesslist code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfaccesslist

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/ospf-access-list/resourcemodel"
)

// NewProtocolsOspfAccessList method to return the example resource reference
func NewProtocolsOspfAccessList() resource.Resource {
	return &protocolsOspfAccessList{
		model: &resourcemodel.ProtocolsOspfAccessList{},
	}
}

// protocolsOspfAccessList defines the resource implementation.
type protocolsOspfAccessList struct {
	client *client.Client
	model  *resourcemodel.ProtocolsOspfAccessList
}

// GetClient returns the vyos api client
func (r *protocolsOspfAccessList) GetClient() *client.Client {
	return r.client
}

func (r *protocolsOspfAccessList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
