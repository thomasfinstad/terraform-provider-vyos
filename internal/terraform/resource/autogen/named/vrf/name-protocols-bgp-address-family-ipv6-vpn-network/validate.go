// Package namedvrfnameprotocolsbgpaddressfamilyipvsixvpnnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixvpnnetwork

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixVpnNetwork{}
