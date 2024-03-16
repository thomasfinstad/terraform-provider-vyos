// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvfour{}

// VrfNameProtocolsIsisRedistributeIPvfour describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvfour struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvfourBgp       *VrfNameProtocolsIsisRedistributeIPvfourBgp       `tfsdk:"bgp" vyos:"bgp,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourConnected *VrfNameProtocolsIsisRedistributeIPvfourConnected `tfsdk:"connected" vyos:"connected,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourKernel    *VrfNameProtocolsIsisRedistributeIPvfourKernel    `tfsdk:"kernel" vyos:"kernel,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourOspf      *VrfNameProtocolsIsisRedistributeIPvfourOspf      `tfsdk:"ospf" vyos:"ospf,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourRIP       *VrfNameProtocolsIsisRedistributeIPvfourRIP       `tfsdk:"rip" vyos:"rip,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourBabel     *VrfNameProtocolsIsisRedistributeIPvfourBabel     `tfsdk:"babel" vyos:"babel,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvfourStatic    *VrfNameProtocolsIsisRedistributeIPvfourStatic    `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfour) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourBgp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Border Gateway Protocol (BGP)

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes into IS-IS

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes into IS-IS

`,
		},

		"ospf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourOspf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute OSPF routes into IS-IS

`,
		},

		"rip": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourRIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIP routes into IS-IS

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes into IS-IS

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes into IS-IS

`,
		},
	}
}
