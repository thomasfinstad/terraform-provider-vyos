// Package namedpkica code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkica

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &pkiCa{}
	_ resource.ResourceWithConfigure = &pkiCa{}
)

// var _ resource.ResourceWithConfigValidators = &pkiCa{}
// var _ resource.ResourceWithModifyPlan = &pkiCa{}
// var _ resource.ResourceWithUpgradeState = &pkiCa{}
// var _ resource.ResourceWithValidateConfig = &pkiCa{}
// var _ resource.ResourceWithImportState = &pkiCa{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiCa{}
