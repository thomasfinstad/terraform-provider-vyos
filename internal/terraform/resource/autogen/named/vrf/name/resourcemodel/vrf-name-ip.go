// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameIP{}

// VrfNameIP describes the resource data model.
type VrfNameIP struct {
	// LeafNodes
	LeafVrfNameIPDisableForwarding types.Bool `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameIPProtocol bool `tfsdk:"protocol" vyos:"protocol,child"`

	// Nodes
	NodeVrfNameIPNht *VrfNameIPNht `tfsdk:"nht" vyos:"nht,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable_forwarding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Description: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"nht": schema.SingleNestedAttribute{
			Attributes: VrfNameIPNht{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Filter Next Hop tracking route resolution

`,
			Description: `Filter Next Hop tracking route resolution

`,
		},
	}
}
