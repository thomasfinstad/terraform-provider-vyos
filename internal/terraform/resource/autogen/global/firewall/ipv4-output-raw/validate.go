// Package globalfirewallipvfouroutputraw code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfouroutputraw

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourOutputRaw{}
	_ resource.ResourceWithConfigure = &firewallIPvfourOutputRaw{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourOutputRaw{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourOutputRaw{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourOutputRaw{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourOutputRaw{}
// var _ resource.ResourceWithImportState = &firewallIPvfourOutputRaw{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourOutputRaw{}