// Package namedvrfnameprotocolsbgpneighbor code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpneighbor

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpNeighbor{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpNeighbor{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpNeighbor{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpNeighbor{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpNeighbor{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpNeighbor{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpNeighbor{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpNeighbor{}
