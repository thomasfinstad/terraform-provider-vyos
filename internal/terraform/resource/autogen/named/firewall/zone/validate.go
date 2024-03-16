// Package namedfirewallzone code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallzone

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallZone{}
	_ resource.ResourceWithConfigure = &firewallZone{}
)

// var _ resource.ResourceWithConfigValidators = &firewallZone{}
// var _ resource.ResourceWithModifyPlan = &firewallZone{}
// var _ resource.ResourceWithUpgradeState = &firewallZone{}
// var _ resource.ResourceWithValidateConfig = &firewallZone{}
// var _ resource.ResourceWithImportState = &firewallZone{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallZone{}
