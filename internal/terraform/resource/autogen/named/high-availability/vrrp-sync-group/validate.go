// Package namedhighavailabilityvrrpsyncgroup code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvrrpsyncgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &highAvailabilityVrrpSyncGroup{}
	_ resource.ResourceWithConfigure = &highAvailabilityVrrpSyncGroup{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVrrpSyncGroup{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVrrpSyncGroup{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVrrpSyncGroup{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVrrpSyncGroup{}
// var _ resource.ResourceWithImportState = &highAvailabilityVrrpSyncGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVrrpSyncGroup{}
