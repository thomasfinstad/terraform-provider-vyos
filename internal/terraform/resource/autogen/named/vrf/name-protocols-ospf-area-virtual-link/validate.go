// Package namedvrfnameprotocolsospfareavirtuallink code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfareavirtuallink

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfAreaVirtualLink{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfAreaVirtualLink{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfAreaVirtualLink{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfAreaVirtualLink{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfAreaVirtualLink{}
