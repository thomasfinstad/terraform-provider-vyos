// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsStaticTable describes the resource data model.
type ProtocolsStaticTable struct {
	// LeafNodes
	ProtocolsStaticTableDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	ProtocolsStaticTableRoute    types.Map `tfsdk:"route" json:"route,omitempty"`
	ProtocolsStaticTableRoutesix types.Map `tfsdk:"route6" json:"route6,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsStaticTable) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		// TagNodes

		"route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticTableRoute{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Static IPv4 route

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 static route  |
`,
		},

		"route6": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsStaticTableRoutesix{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Static IPv6 route

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 static route  |
`,
		},

		// Nodes

	}
}
