// Package namedqospolicylimiterclass code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicylimiterclass

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyLimiterClass{}
	_ resource.ResourceWithConfigure = &qosPolicyLimiterClass{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyLimiterClass{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyLimiterClass{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyLimiterClass{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyLimiterClass{}
// var _ resource.ResourceWithImportState = &qosPolicyLimiterClass{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyLimiterClass{}
