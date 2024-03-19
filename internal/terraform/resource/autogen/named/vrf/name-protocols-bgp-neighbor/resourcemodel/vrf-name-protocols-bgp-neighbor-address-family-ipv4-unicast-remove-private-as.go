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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAsAll types.Bool `tfsdk:"all" vyos:"all,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"all": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Remove private AS numbers to all AS numbers in outbound route updates

`,
			Description: `Remove private AS numbers to all AS numbers in outbound route updates

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
