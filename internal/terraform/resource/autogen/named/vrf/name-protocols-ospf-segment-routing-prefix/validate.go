// Package namedvrfnameprotocolsospfsegmentroutingprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfsegmentroutingprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsOspfSegmentRoutingPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsOspfSegmentRoutingPrefix{}
