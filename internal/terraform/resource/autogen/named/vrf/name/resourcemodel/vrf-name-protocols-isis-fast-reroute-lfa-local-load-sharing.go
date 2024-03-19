// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing{}

// VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable *VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable `tfsdk:"disable" vyos:"disable,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"disable": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable load sharing

`,
			Description: `Disable load sharing

`,
		},
	}
}
