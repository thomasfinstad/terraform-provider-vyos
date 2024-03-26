// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisInterfaceLdpSync{}

// VrfNameProtocolsIsisInterfaceLdpSync describes the resource data model.
type VrfNameProtocolsIsisInterfaceLdpSync struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceLdpSyncDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceLdpSyncHolddown types.Number `tfsdk:"holddown" vyos:"holddown,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterfaceLdpSync) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"holddown": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hold down timer for LDP-IGP cost restoration

    |  Format   |  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  |  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |
`,
			Description: `Hold down timer for LDP-IGP cost restoration

    |  Format   |  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  |  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |
`,
		},

		// Nodes

	}
}
