// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesTunnelParametersErspan describes the resource data model.
type InterfacesTunnelParametersErspan struct {
	// LeafNodes
	InterfacesTunnelParametersErspanDirection customtypes.CustomStringValue `tfsdk:"direction" json:"direction,omitempty"`
	InterfacesTunnelParametersErspanHwID      customtypes.CustomStringValue `tfsdk:"hw_id" json:"hw-id,omitempty"`
	InterfacesTunnelParametersErspanIndex     customtypes.CustomStringValue `tfsdk:"index" json:"index,omitempty"`
	InterfacesTunnelParametersErspanVersion   customtypes.CustomStringValue `tfsdk:"version" json:"version,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesTunnelParametersErspan) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"direction": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Mirrored traffic direction

|  Format  |  Description  |
|----------|---------------|
|  ingress  |  Mirror ingress traffic  |
|  egress  |  Mirror egress traffic  |
`,
		},

		"hw_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Unique identifier of an ERSPAN engine within a system

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048575  |  Unique identifier of an ERSPAN engine  |
`,
		},

		"index": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `ERSPAN version 1 index field

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  Platform-depedent field for specifying port number and direction  |
`,
		},

		"version": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Protocol version

|  Format  |  Description  |
|----------|---------------|
|  1  |  ERSPAN Type II  |
|  2  |  ERSPAN Type III  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
