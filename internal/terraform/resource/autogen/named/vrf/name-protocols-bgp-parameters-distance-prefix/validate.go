// Package namedvrfnameprotocolsbgpparametersdistanceprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpparametersdistanceprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpParametersDistancePrefix{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpParametersDistancePrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpParametersDistancePrefix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpParametersDistancePrefix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpParametersDistancePrefix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpParametersDistancePrefix{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpParametersDistancePrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpParametersDistancePrefix{}
