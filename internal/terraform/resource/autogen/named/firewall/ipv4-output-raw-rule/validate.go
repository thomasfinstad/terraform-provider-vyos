// Package namedfirewallipvfouroutputrawrule code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfouroutputrawrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourOutputRawRule{}
	_ resource.ResourceWithConfigure = &firewallIPvfourOutputRawRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourOutputRawRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourOutputRawRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourOutputRawRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourOutputRawRule{}
// var _ resource.ResourceWithImportState = &firewallIPvfourOutputRawRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourOutputRawRule{}
