// Package globalfirewallglobaloptions code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptions

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallGlobalOptions{}
	_ resource.ResourceWithConfigure = &firewallGlobalOptions{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptions{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptions{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptions{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptions{}
// var _ resource.ResourceWithImportState = &firewallGlobalOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptions{}
