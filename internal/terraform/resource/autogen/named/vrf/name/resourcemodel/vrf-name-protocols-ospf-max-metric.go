// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfMaxMetric{}

// VrfNameProtocolsOspfMaxMetric describes the resource data model.
type VrfNameProtocolsOspfMaxMetric struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfMaxMetricRouterLsa *VrfNameProtocolsOspfMaxMetricRouterLsa `tfsdk:"router_lsa" vyos:"router-lsa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfMaxMetric) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"router_lsa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfMaxMetricRouterLsa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise own Router-LSA with infinite distance (stub router)

`,
		},
	}
}
