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

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PolicyExtcommunityListRule{}

// PolicyExtcommunityListRule describes the resource data model.
type PolicyExtcommunityListRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDPolicyExtcommunityList types.String `tfsdk:"extcommunity_list_id" vyos:"extcommunity-list,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyExtcommunityListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyExtcommunityListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyExtcommunityListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyExtcommunityListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyExtcommunityListRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyExtcommunityListRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyExtcommunityListRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyExtcommunityListRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"extcommunity-list",
		o.ParentIDPolicyExtcommunityList.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyExtcommunityListRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"extcommunity-list",
		o.ParentIDPolicyExtcommunityList.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyExtcommunityListRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Rule for this BGP extended community list

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Extended community-list rule number  |
`,
			Description: `Rule for this BGP extended community list

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Extended community-list rule number  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"extcommunity_list_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Add a BGP extended community list entry

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  BGP extended community-list name  |
`,
			Description: `Add a BGP extended community list entry

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  BGP extended community-list name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in extcommunity_list_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  extcommunity_list_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

    |  Format  |  Description              |
    |----------|---------------------------|
    |  permit  |  Permit matching entries  |
    |  deny    |  Deny matching entries    |
`,
			Description: `Action to take on entries matching this rule

    |  Format  |  Description              |
    |----------|---------------------------|
    |  permit  |  Permit matching entries  |
    |  deny    |  Deny matching entries    |
`,
		},

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

		"regex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against an extended community list

    |  Format          |  Description                                 |
    |------------------|----------------------------------------------|
    |  <aa:nn:nn>      |  Extended community list regular expression  |
    |  <rt aa:nn:nn>   |  Route Target regular expression             |
    |  <soo aa:nn:nn>  |  Site of Origin regular expression           |
`,
			Description: `Regular expression to match against an extended community list

    |  Format          |  Description                                 |
    |------------------|----------------------------------------------|
    |  <aa:nn:nn>      |  Extended community list regular expression  |
    |  <rt aa:nn:nn>   |  Route Target regular expression             |
    |  <soo aa:nn:nn>  |  Site of Origin regular expression           |
`,
		},

		// Nodes

	}
}
