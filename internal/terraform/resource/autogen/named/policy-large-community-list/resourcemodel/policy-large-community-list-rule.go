// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyLargeCommunityListRule describes the resource data model.
type PolicyLargeCommunityListRule struct {
	// LeafNodes
	PolicyLargeCommunityListRuleAction      customtypes.CustomStringValue `tfsdk:"action" json:"action,omitempty"`
	PolicyLargeCommunityListRuleDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	PolicyLargeCommunityListRuleRegex       customtypes.CustomStringValue `tfsdk:"regex" json:"regex,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyLargeCommunityListRule) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Action to take on entries matching this rule

|  Format  |  Description  |
|----------|---------------|
|  permit  |  Permit matching entries  |
|  deny  |  Deny matching entries  |
`,
		},

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"regex": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Regular expression to match against a large community list

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN:NN  |  BGP large-community-list filter  |
|  IP:NN:NN  |  BGP large-community-list filter (IPv4 address format)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
