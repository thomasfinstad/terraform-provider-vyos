// Package namedqospolicyrandomdetect code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyrandomdetect

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/qos/policy-random-detect/resourcemodel"
)

// NewQosPolicyRandomDetect method to return the example resource reference
func NewQosPolicyRandomDetect() resource.Resource {
	return &qosPolicyRandomDetect{
		model: &resourcemodel.QosPolicyRandomDetect{},
	}
}

// qosPolicyRandomDetect defines the resource implementation.
type qosPolicyRandomDetect struct {
	providerData data.ProviderData
	model        *resourcemodel.QosPolicyRandomDetect
}

// GetClient returns the vyos api client
func (r *qosPolicyRandomDetect) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *qosPolicyRandomDetect) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *qosPolicyRandomDetect) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *qosPolicyRandomDetect) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
