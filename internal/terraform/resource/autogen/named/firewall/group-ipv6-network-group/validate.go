// Package namedfirewallgroupipvsixnetworkgroup code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallgroupipvsixnetworkgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGroupIPvsixNetworkGroup{}
	_ resource.ResourceWithConfigure = &firewallGroupIPvsixNetworkGroup{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithModifyPlan = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithUpgradeState = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithValidateConfig = &firewallGroupIPvsixNetworkGroup{}
// var _ resource.ResourceWithImportState = &firewallGroupIPvsixNetworkGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGroupIPvsixNetworkGroup{}
