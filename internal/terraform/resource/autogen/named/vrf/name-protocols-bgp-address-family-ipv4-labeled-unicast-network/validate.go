// Package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourlabeledunicastnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicastNetwork{}
