// Package namedqospolicyshaperclass code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyshaperclass

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyShaperClass{}
	_ resource.ResourceWithConfigure = &qosPolicyShaperClass{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyShaperClass{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyShaperClass{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyShaperClass{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyShaperClass{}
// var _ resource.ResourceWithImportState = &qosPolicyShaperClass{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyShaperClass{}
