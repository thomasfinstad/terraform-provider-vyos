// Package namednetnsname code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednetnsname

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &netnsName{}
	_ resource.ResourceWithConfigure = &netnsName{}
)

// var _ resource.ResourceWithConfigValidators = &netnsName{}
// var _ resource.ResourceWithModifyPlan = &netnsName{}
// var _ resource.ResourceWithUpgradeState = &netnsName{}
// var _ resource.ResourceWithValidateConfig = &netnsName{}
// var _ resource.ResourceWithImportState = &netnsName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &netnsName{}
