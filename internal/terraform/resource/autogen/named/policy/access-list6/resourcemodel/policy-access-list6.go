// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyAccessListsix describes the resource data model.
type PolicyAccessListsix struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"access_list6_id" vyos:"-,self-id"`

	// LeafNodes
	LeafPolicyAccessListsixDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagPolicyAccessListsixRule bool `tfsdk:"-" vyos:"rule,ignore,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyAccessListsix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAccessListsix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"policy",

		"access-list6",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessListsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"access_list6_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 access-list filter

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of IPv6 access-list  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		// Nodes

	}
}
