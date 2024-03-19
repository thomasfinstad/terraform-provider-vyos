// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

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

var _ helpers.VyosTopResourceDataModel = &PolicyAsPathListRule{}

// PolicyAsPathListRule describes the resource data model.
type PolicyAsPathListRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	ParentIDPolicyAsPathList types.String `tfsdk:"as_path_list_id" vyos:"as-path-list,parent-id"`

	// LeafNodes
	LeafPolicyAsPathListRuleAction      types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafPolicyAsPathListRuleDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPolicyAsPathListRuleRegex       types.String `tfsdk:"regex" vyos:"regex,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyAsPathListRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyAsPathListRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAsPathListRule) GetVyosPath() []string {
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
func (o *PolicyAsPathListRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"as-path-list",
		o.ParentIDPolicyAsPathList.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyAsPathListRule) GetVyosNamedParentPath() []string {
	return []string{
		"policy",

		"as-path-list",
		o.ParentIDPolicyAsPathList.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAsPathListRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Rule for this as-path-list

    |  Format   &emsp;|  Description               |
    |-----------------|----------------------------|
    |  1-65535  &emsp;|  AS path list rule number  |
`,
			Description: `Rule for this as-path-list

    |  Format   |  Description               |
    |-----------------|----------------------------|
    |  1-65535  |  AS path list rule number  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"as_path_list_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Add a BGP autonomous system path filter

    |  Format  &emsp;|  Description        |
    |----------------|---------------------|
    |  txt     &emsp;|  AS path list name  |
`,
			Description: `Add a BGP autonomous system path filter

    |  Format  |  Description        |
    |----------------|---------------------|
    |  txt     |  AS path list name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in as_path_list_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  as_path_list_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take on entries matching this rule

    |  Format  &emsp;|  Description              |
    |----------------|---------------------------|
    |  permit  &emsp;|  Permit matching entries  |
    |  deny    &emsp;|  Deny matching entries    |
`,
			Description: `Action to take on entries matching this rule

    |  Format  |  Description              |
    |----------------|---------------------------|
    |  permit  |  Permit matching entries  |
    |  deny    |  Deny matching entries    |
`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  &emsp;|  Description  |
    |----------------|---------------|
    |  txt     &emsp;|  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------------|---------------|
    |  txt     |  Description  |
`,
		},

		"regex": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression to match against an AS path

    |  Format  &emsp;|  Description                                     |
    |----------------|--------------------------------------------------|
    |  txt     &emsp;|  AS path regular expression (ex: "64501 64502")  |
`,
			Description: `Regular expression to match against an AS path

    |  Format  |  Description                                     |
    |----------------|--------------------------------------------------|
    |  txt     |  AS path regular expression (ex: "64501 64502")  |
`,
		},

		// Nodes

	}
}
