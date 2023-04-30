// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceSnmpVthreeUserPrivacy describes the resource data model.
type ServiceSnmpVthreeUserPrivacy struct {
	// LeafNodes
	ServiceSnmpVthreeUserPrivacyEncryptedPassword customtypes.CustomStringValue `tfsdk:"encrypted_password" json:"encrypted-password,omitempty"`
	ServiceSnmpVthreeUserPrivacyPlaintextPassword customtypes.CustomStringValue `tfsdk:"plaintext_password" json:"plaintext-password,omitempty"`
	ServiceSnmpVthreeUserPrivacyType              customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceSnmpVthreeUserPrivacy) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"encrypted_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the encrypted key for privacy protocol

`,
		},

		"plaintext_password": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the clear text key for privacy protocol

`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the protocol for privacy

|  Format  |  Description  |
|----------|---------------|
|  des  |  Data Encryption Standard  |
|  aes  |  Advanced Encryption Standard  |
`,

			// Default:          stringdefault.StaticString(`des`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
