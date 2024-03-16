// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeTable types.String `tfsdk:"table" vyos:"table,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected `tfsdk:"connected" vyos:"connected,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeIsis      *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeIsis      `tfsdk:"isis" vyos:"isis,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeKernel    *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeKernel    `tfsdk:"kernel" vyos:"kernel,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf      *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf      `tfsdk:"ospf" vyos:"ospf,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP       *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP       `tfsdk:"rip" vyos:"rip,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel     *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel     `tfsdk:"babel" vyos:"babel,omitempty"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeStatic    *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeStatic    `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistribute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redistribute non-main Kernel Routing Table

`,
		},

		// Nodes

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes into BGP

`,
		},

		"isis": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeIsis{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute IS-IS routes into BGP

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes into BGP

`,
		},

		"ospf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeOspf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute OSPF routes into BGP

`,
		},

		"rip": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeRIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIP routes into BGP

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes into BGP

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRedistributeStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes into BGP

`,
		},
	}
}
