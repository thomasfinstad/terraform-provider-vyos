// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}

// VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall *VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall `tfsdk:"local_install" vyos:"local-install,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"local_install": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspecLocalInstall{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Apply local policy routing to interface

`,
			Description: `Apply local policy routing to interface

`,
		},
	}
}
