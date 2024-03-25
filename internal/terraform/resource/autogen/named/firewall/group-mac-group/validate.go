// Package namedfirewallgroupmacgroup code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupmacgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGroupMacGroup{}
	_ resource.ResourceWithConfigure = &firewallGroupMacGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupMacGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupMacGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupMacGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupMacGroup{}
// var _ resource.ResourceWithImportState = &firewallGroupMacGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupMacGroup{}
