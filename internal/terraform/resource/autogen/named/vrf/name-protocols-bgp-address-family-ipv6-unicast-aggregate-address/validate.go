// Package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastaggregateaddress code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastaggregateaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastAggregateAddress{}
