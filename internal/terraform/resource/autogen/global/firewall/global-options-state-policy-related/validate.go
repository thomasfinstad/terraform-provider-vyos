// Package globalfirewallglobaloptionsstatepolicyrelated code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyrelated

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGlobalOptionsStatePolicyRelated{}
	_ resource.ResourceWithConfigure = &firewallGlobalOptionsStatePolicyRelated{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptionsStatePolicyRelated{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptionsStatePolicyRelated{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptionsStatePolicyRelated{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptionsStatePolicyRelated{}
// var _ resource.ResourceWithImportState = &firewallGlobalOptionsStatePolicyRelated{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptionsStatePolicyRelated{}
