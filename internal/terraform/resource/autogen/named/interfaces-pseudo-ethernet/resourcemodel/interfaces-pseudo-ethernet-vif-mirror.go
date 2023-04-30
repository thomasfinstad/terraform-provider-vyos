// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesPseudoEthernetVifMirror describes the resource data model.
type InterfacesPseudoEthernetVifMirror struct {
	// LeafNodes
	InterfacesPseudoEthernetVifMirrorIngress customtypes.CustomStringValue `tfsdk:"ingress" json:"ingress,omitempty"`
	InterfacesPseudoEthernetVifMirrorEgress  customtypes.CustomStringValue `tfsdk:"egress" json:"egress,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesPseudoEthernetVifMirror) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
		},

		"egress": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
