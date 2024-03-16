// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &HighAvailability{}

// HighAvailability describes the resource data model.
type HighAvailability struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafHighAvailabilityDisable types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVirtualServer bool `tfsdk:"-" vyos:"virtual-server,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeHighAvailabilityVrrp bool `tfsdk:"-" vyos:"vrrp,child"`
}

// SetID configures the resource ID
func (o *HighAvailability) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailability) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
