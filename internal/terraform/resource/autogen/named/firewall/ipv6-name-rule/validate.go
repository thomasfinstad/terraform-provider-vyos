// Package namedfirewallipvsixnamerule code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixnamerule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvsixNameRule{}
	_ resource.ResourceWithConfigure = &firewallIPvsixNameRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixNameRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixNameRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixNameRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixNameRule{}
// var _ resource.ResourceWithImportState = &firewallIPvsixNameRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixNameRule{}
