// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesDummyIP describes the resource data model.
type InterfacesDummyIP struct {
	// LeafNodes
	InterfacesDummyIPSourceValIDation  customtypes.CustomStringValue `tfsdk:"source_validation" json:"source-validation,omitempty"`
	InterfacesDummyIPDisableForwarding customtypes.CustomStringValue `tfsdk:"disable_forwarding" json:"disable-forwarding,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesDummyIP) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source_validation": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |
`,
		},

		"disable_forwarding": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		// TagNodes

		// Nodes

	}
}
