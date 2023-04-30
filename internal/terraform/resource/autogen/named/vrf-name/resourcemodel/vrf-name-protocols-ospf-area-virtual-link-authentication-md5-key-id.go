// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID describes the resource data model.
type VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID struct {
	// LeafNodes
	VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyIDMdfiveKey customtypes.CustomStringValue `tfsdk:"md5_key" json:"md5-key,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"md5_key": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `MD5 authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  MD5 Key (16 characters or less)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
