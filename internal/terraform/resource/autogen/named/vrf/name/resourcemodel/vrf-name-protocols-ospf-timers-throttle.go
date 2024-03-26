// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfTimersThroTTLe{}

// VrfNameProtocolsOspfTimersThroTTLe describes the resource data model.
type VrfNameProtocolsOspfTimersThroTTLe struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfTimersThroTTLeSpf *VrfNameProtocolsOspfTimersThroTTLeSpf `tfsdk:"spf" vyos:"spf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfTimersThroTTLe) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"spf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfTimersThroTTLeSpf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPF SPF timers

`,
			Description: `OSPF SPF timers

`,
		},
	}
}
