// Package namedqospolicylimiterclassmatch code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicylimiterclassmatch

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyLimiterClassMatch{}
	_ resource.ResourceWithConfigure = &qosPolicyLimiterClassMatch{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyLimiterClassMatch{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyLimiterClassMatch{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyLimiterClassMatch{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyLimiterClassMatch{}
// var _ resource.ResourceWithImportState = &qosPolicyLimiterClassMatch{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyLimiterClassMatch{}
