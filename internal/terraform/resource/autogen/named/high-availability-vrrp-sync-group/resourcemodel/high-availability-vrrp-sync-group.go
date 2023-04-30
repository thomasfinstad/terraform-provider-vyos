// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// HighAvailabilityVrrpSyncGroup describes the resource data model.
type HighAvailabilityVrrpSyncGroup struct {
	// LeafNodes
	HighAvailabilityVrrpSyncGroupMember customtypes.CustomStringValue `tfsdk:"member" json:"member,omitempty"`

	// TagNodes

	// Nodes
	HighAvailabilityVrrpSyncGroupTransitionScrIPt types.Object `tfsdk:"transition_script" json:"transition-script,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroup) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"member": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Sync group member

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRRP group name  |
`,
		},

		// TagNodes

		// Nodes

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupTransitionScrIPt{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
		},
	}
}
