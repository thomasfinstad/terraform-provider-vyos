// Package namedqospolicyroundrobinclassmatch code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobinclassmatch

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &qosPolicyRoundRobinClassMatch{}
	_ resource.ResourceWithConfigure = &qosPolicyRoundRobinClassMatch{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRoundRobinClassMatch{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRoundRobinClassMatch{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRoundRobinClassMatch{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRoundRobinClassMatch{}
// var _ resource.ResourceWithImportState = &qosPolicyRoundRobinClassMatch{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRoundRobinClassMatch{}
