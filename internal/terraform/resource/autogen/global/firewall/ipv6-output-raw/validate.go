// Package globalfirewallipvsixoutputraw code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvsixoutputraw

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvsixOutputRaw{}
	_ resource.ResourceWithConfigure = &firewallIPvsixOutputRaw{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixOutputRaw{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixOutputRaw{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixOutputRaw{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixOutputRaw{}
// var _ resource.ResourceWithImportState = &firewallIPvsixOutputRaw{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixOutputRaw{}
