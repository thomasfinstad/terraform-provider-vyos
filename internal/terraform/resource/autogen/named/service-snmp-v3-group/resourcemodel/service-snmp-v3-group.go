// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceSnmpVthreeGroup describes the resource data model.
type ServiceSnmpVthreeGroup struct {
	// LeafNodes
	ServiceSnmpVthreeGroupMode     customtypes.CustomStringValue `tfsdk:"mode" json:"mode,omitempty"`
	ServiceSnmpVthreeGroupSeclevel customtypes.CustomStringValue `tfsdk:"seclevel" json:"seclevel,omitempty"`
	ServiceSnmpVthreeGroupView     customtypes.CustomStringValue `tfsdk:"view" json:"view,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceSnmpVthreeGroup) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Define access permission

|  Format  |  Description  |
|----------|---------------|
|  ro  |  Read-Only  |
|  rw  |  read write  |
`,

			// Default:          stringdefault.StaticString(`ro`),
			Computed: true,
		},

		"seclevel": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Security levels

|  Format  |  Description  |
|----------|---------------|
|  noauth  |  Messages not authenticated and not encrypted (noAuthNoPriv)  |
|  auth  |  Messages are authenticated but not encrypted (authNoPriv)  |
|  priv  |  Messages are authenticated and encrypted (authPriv)  |
`,

			// Default:          stringdefault.StaticString(`auth`),
			Computed: true,
		},

		"view": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Defines the name of view

`,
		},

		// TagNodes

		// Nodes

	}
}
