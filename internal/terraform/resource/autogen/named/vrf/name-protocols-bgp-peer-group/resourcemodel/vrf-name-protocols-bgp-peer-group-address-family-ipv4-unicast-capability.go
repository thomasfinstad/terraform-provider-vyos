// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf `tfsdk:"orf" vyos:"orf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"orf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastCapabilityOrf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise ORF capability to this peer

`,
		},
	}
}