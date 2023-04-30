// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsOspfAreaAreaTypeNssa describes the resource data model.
type VrfNameProtocolsOspfAreaAreaTypeNssa struct {
	// LeafNodes
	VrfNameProtocolsOspfAreaAreaTypeNssaDefaultCost customtypes.CustomStringValue `tfsdk:"default_cost" json:"default-cost,omitempty"`
	VrfNameProtocolsOspfAreaAreaTypeNssaNoSummary   customtypes.CustomStringValue `tfsdk:"no_summary" json:"no-summary,omitempty"`
	VrfNameProtocolsOspfAreaAreaTypeNssaTranSLAte   customtypes.CustomStringValue `tfsdk:"translate" json:"translate,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaAreaTypeNssa) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_cost": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Summary-default cost of an NSSA area

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Summary default cost  |
`,
		},

		"no_summary": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not inject inter-area routes into stub

`,
		},

		"translate": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Configure NSSA-ABR

|  Format  |  Description  |
|----------|---------------|
|  always  |  Always translate LSA types  |
|  candidate  |  Translate for election  |
|  never  |  Never translate LSA types  |
`,

			// Default:          stringdefault.StaticString(`candidate`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
