// Package namedpolicyprefixlistrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyprefixlistrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyPrefixListRule{}
	_ resource.ResourceWithConfigure = &policyPrefixListRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyPrefixListRule{}
// var _ resource.ResourceWithModifyPlan = &policyPrefixListRule{}
// var _ resource.ResourceWithUpgradeState = &policyPrefixListRule{}
// var _ resource.ResourceWithValidateConfig = &policyPrefixListRule{}
// var _ resource.ResourceWithImportState = &policyPrefixListRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyPrefixListRule{}
