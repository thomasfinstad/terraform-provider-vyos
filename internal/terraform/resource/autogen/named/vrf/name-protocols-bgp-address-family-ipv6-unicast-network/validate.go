// Package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastnetwork code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastNetwork{}
