// Package namedfirewallzonefrom code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallzonefrom

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallZoneFrom{}
	_ resource.ResourceWithConfigure = &firewallZoneFrom{}
)

// var _ resource.ResourceWithConfigValidators = &firewallZoneFrom{}
// var _ resource.ResourceWithModifyPlan = &firewallZoneFrom{}
// var _ resource.ResourceWithUpgradeState = &firewallZoneFrom{}
// var _ resource.ResourceWithValidateConfig = &firewallZoneFrom{}
// var _ resource.ResourceWithImportState = &firewallZoneFrom{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallZoneFrom{}
