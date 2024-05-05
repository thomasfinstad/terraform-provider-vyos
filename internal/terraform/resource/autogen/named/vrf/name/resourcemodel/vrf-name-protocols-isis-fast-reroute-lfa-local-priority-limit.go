// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit{}

// VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium   *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium   `tfsdk:"medium" vyos:"medium,omitempty"`
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh     *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh     `tfsdk:"high" vyos:"high,omitempty"`
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical *VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical `tfsdk:"critical" vyos:"critical,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"medium": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitMedium{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical, high, and medium priority prefixes

`,
			Description: `Compute for critical, high, and medium priority prefixes

`,
		},

		"high": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitHigh{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical, and high priority prefixes

`,
			Description: `Compute for critical, and high priority prefixes

`,
		},

		"critical": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalPriorityLimitCritical{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Compute for critical priority prefixes only

`,
			Description: `Compute for critical priority prefixes only

`,
		},
	}
}
