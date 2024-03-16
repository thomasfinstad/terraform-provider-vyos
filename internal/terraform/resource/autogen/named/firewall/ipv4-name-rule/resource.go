// Package namedfirewallipvfournamerule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfournamerule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name-rule/resourcemodel"
)

// NewFirewallIPvfourNameRule method to return the example resource reference
func NewFirewallIPvfourNameRule() resource.Resource {
	return &firewallIPvfourNameRule{
		model: &resourcemodel.FirewallIPvfourNameRule{},
	}
}

// firewallIPvfourNameRule defines the resource implementation.
type firewallIPvfourNameRule struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallIPvfourNameRule
}

// GetClient returns the vyos api client
func (r *firewallIPvfourNameRule) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallIPvfourNameRule) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallIPvfourNameRule) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallIPvfourNameRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
