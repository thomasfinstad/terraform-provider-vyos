// Package namednatdestinationrule code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatdestinationrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &natDestinationRule{}
	_ resource.ResourceWithConfigure = &natDestinationRule{}
)

// var _ resource.ResourceWithConfigValidators = &natDestinationRule{}
// var _ resource.ResourceWithModifyPlan = &natDestinationRule{}
// var _ resource.ResourceWithUpgradeState = &natDestinationRule{}
// var _ resource.ResourceWithValidateConfig = &natDestinationRule{}
// var _ resource.ResourceWithImportState = &natDestinationRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &natDestinationRule{}
