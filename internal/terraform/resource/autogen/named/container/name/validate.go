// Package namedcontainername code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainername

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &containerName{}
	_ resource.ResourceWithConfigure = &containerName{}
)

// var _ resource.ResourceWithConfigValidators = &containerName{}
// var _ resource.ResourceWithModifyPlan = &containerName{}
// var _ resource.ResourceWithUpgradeState = &containerName{}
// var _ resource.ResourceWithValidateConfig = &containerName{}
// var _ resource.ResourceWithImportState = &containerName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerName{}
