// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// PolicyAccessListsix describes the resource data model.
type PolicyAccessListsix struct {
	// LeafNodes
	PolicyAccessListsixDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`

	// TagNodes
	PolicyAccessListsixRule types.Map `tfsdk:"rule" json:"rule,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o PolicyAccessListsix) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		// TagNodes

		"rule": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: PolicyAccessListsixRule{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Rule for this access-list6

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Access-list6 rule number  |
`,
		},

		// Nodes

	}
}
