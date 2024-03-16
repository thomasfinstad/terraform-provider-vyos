// Package namedfirewallgroupdynamicgroupaddressgroup code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupdynamicgroupaddressgroup

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/group-dynamic-group-address-group/resourcemodel"
)

// NewFirewallGroupDynamicGroupAddressGroup method to return the example resource reference
func NewFirewallGroupDynamicGroupAddressGroup() resource.Resource {
	return &firewallGroupDynamicGroupAddressGroup{
		model: &resourcemodel.FirewallGroupDynamicGroupAddressGroup{},
	}
}

// firewallGroupDynamicGroupAddressGroup defines the resource implementation.
type firewallGroupDynamicGroupAddressGroup struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallGroupDynamicGroupAddressGroup
}

// GetClient returns the vyos api client
func (r *firewallGroupDynamicGroupAddressGroup) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallGroupDynamicGroupAddressGroup) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallGroupDynamicGroupAddressGroup) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallGroupDynamicGroupAddressGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
