// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrpSyncGroup{}

// HighAvailabilityVrrpSyncGroup describes the resource data model.
type HighAvailabilityVrrpSyncGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"sync_group_id" vyos:"-,self-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupMember types.List `tfsdk:"member" vyos:"member,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeHighAvailabilityVrrpSyncGroupHealthCheck      *HighAvailabilityVrrpSyncGroupHealthCheck      `tfsdk:"health_check" vyos:"health-check,omitempty"`
	NodeHighAvailabilityVrrpSyncGroupTransitionScrIPt *HighAvailabilityVrrpSyncGroupTransitionScrIPt `tfsdk:"transition_script" vyos:"transition-script,omitempty"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpSyncGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpSyncGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"vrrp",

		"sync-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"sync_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VRRP sync group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in sync_group_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  sync_group_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"member": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Sync group member

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRRP group name  |

`,
		},

		// Nodes

		"health_check": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupHealthCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Health check

`,
		},

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpSyncGroupTransitionScrIPt{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
		},
	}
}
