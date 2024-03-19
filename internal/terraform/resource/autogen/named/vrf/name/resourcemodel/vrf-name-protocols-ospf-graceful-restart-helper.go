// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfGracefulRestartHelper{}

// VrfNameProtocolsOspfGracefulRestartHelper describes the resource data model.
type VrfNameProtocolsOspfGracefulRestartHelper struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfGracefulRestartHelperPlannedOnly         types.Bool   `tfsdk:"planned_only" vyos:"planned-only,omitempty"`
	LeafVrfNameProtocolsOspfGracefulRestartHelperSupportedGraceTime  types.Number `tfsdk:"supported_grace_time" vyos:"supported-grace-time,omitempty"`
	LeafVrfNameProtocolsOspfGracefulRestartHelperNoStrictLsaChecking types.Bool   `tfsdk:"no_strict_lsa_checking" vyos:"no-strict-lsa-checking,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfGracefulRestartHelperEnable *VrfNameProtocolsOspfGracefulRestartHelperEnable `tfsdk:"enable" vyos:"enable,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfGracefulRestartHelper) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"planned_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Supported only planned restart

`,
			Description: `Supported only planned restart

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"supported_grace_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Supported grace timer

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  10-1800  |  Grace interval in seconds  |
`,
			Description: `Supported grace timer

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  10-1800  |  Grace interval in seconds  |
`,
		},

		"no_strict_lsa_checking": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable strict LSA check

`,
			Description: `Disable strict LSA check

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"enable": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfGracefulRestartHelperEnable{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable helper support

`,
			Description: `Enable helper support

`,
		},
	}
}
