// Package namedpolicylargecommunitylist code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylargecommunitylist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyLargeCommunityList{}
	_ resource.ResourceWithConfigure = &policyLargeCommunityList{}
)

// var _ resource.ResourceWithConfigValidators = &policyLargeCommunityList{}
// var _ resource.ResourceWithModifyPlan = &policyLargeCommunityList{}
// var _ resource.ResourceWithUpgradeState = &policyLargeCommunityList{}
// var _ resource.ResourceWithValidateConfig = &policyLargeCommunityList{}
// var _ resource.ResourceWithImportState = &policyLargeCommunityList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyLargeCommunityList{}
