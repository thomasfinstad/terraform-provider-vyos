// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixListReceive types.Bool `tfsdk:"receive" vyos:"receive,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixListSend    types.Bool `tfsdk:"send" vyos:"send,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixLabeledUnicastCapabilityOrfPrefixList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"receive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Capability to receive the ORF

`,
			Description: `Capability to receive the ORF

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"send": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Capability to send the ORF

`,
			Description: `Capability to send the ORF

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
