// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfTimers{}

// VrfNameProtocolsOspfTimers describes the resource data model.
type VrfNameProtocolsOspfTimers struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfTimersThroTTLe *VrfNameProtocolsOspfTimersThroTTLe `tfsdk:"throttle" vyos:"throttle,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfTimers) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"throttle": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfTimersThroTTLe{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Throttling adaptive timers

`,
		},
	}
}
