// Package namedfirewallipvsixinputfilterrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixinputfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvsixInputFilterRule{}
	_ resource.ResourceWithConfigure = &firewallIPvsixInputFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixInputFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixInputFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixInputFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixInputFilterRule{}
// var _ resource.ResourceWithImportState = &firewallIPvsixInputFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixInputFilterRule{}
