// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperHfscClass{}

// QosPolicyShaperHfscClass describes the resource data model.
type QosPolicyShaperHfscClass struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"class_id" vyos:"-,self-id"`

	ParentIDQosPolicyShaperHfsc types.String `tfsdk:"shaper_hfsc_id" vyos:"shaper-hfsc,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyShaperHfscClassDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagQosPolicyShaperHfscClassMatch bool `tfsdk:"-" vyos:"match,child"`

	// Nodes
	NodeQosPolicyShaperHfscClassLinkshare  *QosPolicyShaperHfscClassLinkshare  `tfsdk:"linkshare" vyos:"linkshare,omitempty"`
	NodeQosPolicyShaperHfscClassRealtime   *QosPolicyShaperHfscClassRealtime   `tfsdk:"realtime" vyos:"realtime,omitempty"`
	NodeQosPolicyShaperHfscClassUpperlimit *QosPolicyShaperHfscClassUpperlimit `tfsdk:"upperlimit" vyos:"upperlimit,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyShaperHfscClass) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyShaperHfscClass) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyShaperHfscClass) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperHfscClass) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"class",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyShaperHfscClass) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper-hfsc",
		o.ParentIDQosPolicyShaperHfsc.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyShaperHfscClass) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper-hfsc",
		o.ParentIDQosPolicyShaperHfsc.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClass) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4095  |  Class Identifier  |
`,
			Description: `Class ID

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-4095  |  Class Identifier  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"shaper_hfsc_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Hierarchical Fair Service Curve's policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Hierarchical Fair Service Curve's policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in shaper_hfsc_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  shaper_hfsc_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		// Nodes

		"linkshare": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassLinkshare{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Linkshare class settings

`,
			Description: `Linkshare class settings

`,
		},

		"realtime": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassRealtime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Realtime class settings

`,
			Description: `Realtime class settings

`,
		},

		"upperlimit": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassUpperlimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Upperlimit class settings

`,
			Description: `Upperlimit class settings

`,
		},
	}
}
