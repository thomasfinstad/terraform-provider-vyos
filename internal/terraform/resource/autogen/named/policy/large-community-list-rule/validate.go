// Package namedpolicylargecommunitylistrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylargecommunitylistrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyLargeCommunityListRule{}
	_ resource.ResourceWithConfigure = &policyLargeCommunityListRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithModifyPlan = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithUpgradeState = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithValidateConfig = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithImportState = &policyLargeCommunityListRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyLargeCommunityListRule{}
