// Package globalfirewallglobaloptionsstatepolicyinvalid code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyinvalid

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGlobalOptionsStatePolicyInvalID{}
	_ resource.ResourceWithConfigure = &firewallGlobalOptionsStatePolicyInvalID{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptionsStatePolicyInvalID{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptionsStatePolicyInvalID{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptionsStatePolicyInvalID{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptionsStatePolicyInvalID{}
// var _ resource.ResourceWithImportState = &firewallGlobalOptionsStatePolicyInvalID{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptionsStatePolicyInvalID{}
