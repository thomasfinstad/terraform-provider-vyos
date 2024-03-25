// Package namedfirewallipvsixforwardfilterrule code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixforwardfilterrule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv6-forward-filter-rule/resourcemodel"
)

// NewFirewallIPvsixForwardFilterRule method to return the example resource reference
func NewFirewallIPvsixForwardFilterRule() resource.Resource {
	return &firewallIPvsixForwardFilterRule{
		model: &resourcemodel.FirewallIPvsixForwardFilterRule{},
	}
}

// firewallIPvsixForwardFilterRule defines the resource implementation.
type firewallIPvsixForwardFilterRule struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallIPvsixForwardFilterRule
}

// GetClient returns the vyos api client
func (r *firewallIPvsixForwardFilterRule) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallIPvsixForwardFilterRule) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallIPvsixForwardFilterRule) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallIPvsixForwardFilterRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
