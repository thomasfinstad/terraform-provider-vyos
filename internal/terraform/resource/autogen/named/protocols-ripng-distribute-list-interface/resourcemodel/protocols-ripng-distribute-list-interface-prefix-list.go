// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsRIPngDistributeListInterfacePrefixList describes the resource data model.
type ProtocolsRIPngDistributeListInterfacePrefixList struct {
	// LeafNodes
	ProtocolsRIPngDistributeListInterfacePrefixListIn  customtypes.CustomStringValue `tfsdk:"in" json:"in,omitempty"`
	ProtocolsRIPngDistributeListInterfacePrefixListOut customtypes.CustomStringValue `tfsdk:"out" json:"out,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterfacePrefixList) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"in": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefix-list to apply to input packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to input packets  |
`,
		},

		"out": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Prefix-list to apply to output packets

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Prefix-list to apply to output packets  |
`,
		},

		// TagNodes

		// Nodes

	}
}
