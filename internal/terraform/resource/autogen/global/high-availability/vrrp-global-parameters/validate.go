// Package globalhighavailabilityvrrpglobalparameters code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrpglobalparameters

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &highAvailabilityVrrpGlobalParameters{}
	_ resource.ResourceWithConfigure = &highAvailabilityVrrpGlobalParameters{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpGlobalParameters{}
// var _ resource.ResourceWithImportState = &highAvailabilityVrrpGlobalParameters{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpGlobalParameters{}
