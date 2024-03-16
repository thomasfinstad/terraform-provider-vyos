// Package namedqospolicyfairqueue code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyfairqueue

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyFairQueue{}
	_ resource.ResourceWithConfigure = &qosPolicyFairQueue{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyFairQueue{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyFairQueue{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyFairQueue{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyFairQueue{}
// var _ resource.ResourceWithImportState = &qosPolicyFairQueue{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyFairQueue{}
