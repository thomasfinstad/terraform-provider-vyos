// Package namedqospolicyroundrobinclassmatch code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobinclassmatch

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/qos/policy-round-robin-class-match/resourcemodel"
)

// NewQosPolicyRoundRobinClassMatch method to return the example resource reference
func NewQosPolicyRoundRobinClassMatch() resource.Resource {
	return &qosPolicyRoundRobinClassMatch{
		model: &resourcemodel.QosPolicyRoundRobinClassMatch{},
	}
}

// qosPolicyRoundRobinClassMatch defines the resource implementation.
type qosPolicyRoundRobinClassMatch struct {
	providerData data.ProviderData
	model        *resourcemodel.QosPolicyRoundRobinClassMatch
}

// GetClient returns the vyos api client
func (r *qosPolicyRoundRobinClassMatch) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *qosPolicyRoundRobinClassMatch) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *qosPolicyRoundRobinClassMatch) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *qosPolicyRoundRobinClassMatch) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
