// Package namedqospolicyrandomdetectprecedence code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyrandomdetectprecedence

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyRandomDetectPrecedence{}
	_ resource.ResourceWithConfigure = &qosPolicyRandomDetectPrecedence{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRandomDetectPrecedence{}
// var _ resource.ResourceWithImportState = &qosPolicyRandomDetectPrecedence{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRandomDetectPrecedence{}
