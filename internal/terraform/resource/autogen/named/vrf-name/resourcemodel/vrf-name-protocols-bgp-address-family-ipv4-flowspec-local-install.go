// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface customtypes.CustomStringValue `tfsdk:"interface" json:"interface,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface

`,
		},

		// TagNodes

		// Nodes

	}
}
