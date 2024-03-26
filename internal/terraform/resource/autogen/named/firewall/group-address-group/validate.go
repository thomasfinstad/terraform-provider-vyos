// Package namedfirewallgroupaddressgroup code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupaddressgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGroupAddressGroup{}
	_ resource.ResourceWithConfigure = &firewallGroupAddressGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupAddressGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupAddressGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupAddressGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupAddressGroup{}
// var _ resource.ResourceWithImportState = &firewallGroupAddressGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupAddressGroup{}
