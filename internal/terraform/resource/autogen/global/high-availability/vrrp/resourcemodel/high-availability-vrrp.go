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
var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrp{}

// HighAvailabilityVrrp describes the resource data model.
type HighAvailabilityVrrp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpSnmp types.Bool `tfsdk:"snmp" vyos:"snmp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVrrpGroup     bool `tfsdk:"-" vyos:"group,child"`
	ExistsTagHighAvailabilityVrrpSyncGroup bool `tfsdk:"-" vyos:"sync-group,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeHighAvailabilityVrrpGlobalParameters bool `tfsdk:"-" vyos:"global-parameters,child"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"vrrp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"snmp": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable SNMP

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
