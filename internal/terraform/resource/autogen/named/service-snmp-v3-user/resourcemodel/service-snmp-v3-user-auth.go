// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceSnmpVthreeUserAuth describes the resource data model.
type ServiceSnmpVthreeUserAuth struct {
	// LeafNodes
	ServiceSnmpVthreeUserAuthEncryptedPassword customtypes.CustomStringValue `tfsdk:"encrypted_password" json:"encrypted-password,omitempty"`
	ServiceSnmpVthreeUserAuthPlaintextPassword customtypes.CustomStringValue `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`
	ServiceSnmpVthreeUserAuthType              customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceSnmpVthreeUserAuth) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the encrypted key for authentication

`,
		},

		"plaintext_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the clear text key for authentication

`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Define used protocol

|  Format  |  Description  |
|----------|---------------|
|  md5  |  Message Digest 5  |
|  sha  |  Secure Hash Algorithm  |
`,

			// Default:          stringdefault.StaticString(`md5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
