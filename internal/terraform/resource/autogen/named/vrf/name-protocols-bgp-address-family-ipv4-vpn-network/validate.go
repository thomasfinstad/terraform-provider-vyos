// Package namedvrfnameprotocolsbgpaddressfamilyipvfourvpnnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourvpnnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}
