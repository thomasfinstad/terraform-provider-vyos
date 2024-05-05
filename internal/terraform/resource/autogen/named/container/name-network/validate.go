// Package namedcontainernamenetwork code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernamenetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &containerNameNetwork{}
	_ resource.ResourceWithConfigure = &containerNameNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &containerNameNetwork{}
// var _ resource.ResourceWithModifyPlan = &containerNameNetwork{}
// var _ resource.ResourceWithUpgradeState = &containerNameNetwork{}
// var _ resource.ResourceWithValidateConfig = &containerNameNetwork{}
// var _ resource.ResourceWithImportState = &containerNameNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerNameNetwork{}
