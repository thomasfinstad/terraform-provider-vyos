// Package namedvrfnameprotocolsstaticroutesixnexthopbfdmultihopsource code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesixnexthopbfdmultihopsource

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRoutesixNextHopBfdMultiHopSource{}
