// Package namedvrfnameprotocolsbgppeergrouplocalas code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgppeergrouplocalas

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpPeerGroupLocalAs{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpPeerGroupLocalAs{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpPeerGroupLocalAs{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpPeerGroupLocalAs{}
