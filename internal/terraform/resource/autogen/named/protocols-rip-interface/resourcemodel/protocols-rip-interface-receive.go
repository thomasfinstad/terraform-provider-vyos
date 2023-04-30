// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsRIPInterfaceReceive describes the resource data model.
type ProtocolsRIPInterfaceReceive struct {
	// LeafNodes
	ProtocolsRIPInterfaceReceiveVersion customtypes.CustomStringValue `tfsdk:"version" json:"version,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsRIPInterfaceReceive) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"version": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Limit RIP protocol version

|  Format  |  Description  |
|----------|---------------|
|  1  |  Allow RIPv1 only  |
|  2  |  Allow RIPv2 only  |
`,
		},

		// TagNodes

		// Nodes

	}
}
