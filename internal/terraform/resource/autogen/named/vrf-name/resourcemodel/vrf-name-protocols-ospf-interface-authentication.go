// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfInterfaceAuthentication describes the resource data model.
type VrfNameProtocolsOspfInterfaceAuthentication struct {
	// LeafNodes
	VrfNameProtocolsOspfInterfaceAuthenticationPlaintextPassword customtypes.CustomStringValue `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsOspfInterfaceAuthenticationMdfive types.Object `tfsdk:"md5" json:"md5,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfInterfaceAuthentication) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"plaintext_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Plain text password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Plain text password (8 characters or less)  |
`,
		},

		// TagNodes

		// Nodes

		"md5": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfInterfaceAuthenticationMdfive{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `MD5 key id

`,
		},
	}
}
