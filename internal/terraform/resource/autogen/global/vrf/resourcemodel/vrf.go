// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Vrf describes the resource data model.
type Vrf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVrfBindToAll types.Bool `tfsdk:"bind_to_all" vyos:"bind-to-all,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfName bool `tfsdk:"-" vyos:"name,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *Vrf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *Vrf) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o Vrf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"bind_to_all": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable binding services to all VRFs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
