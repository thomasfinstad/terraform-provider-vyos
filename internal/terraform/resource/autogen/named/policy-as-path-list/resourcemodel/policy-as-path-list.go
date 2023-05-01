// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyAsPathList describes the resource data model.
type PolicyAsPathList struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafPolicyAsPathListDescrIPtion types.String `tfsdk:"description"`

	// TagNodes
	TagPolicyAsPathListRule types.Map `tfsdk:"rule"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAsPathList) GetVyosPath() []string {
	return []string{
		"policy",
		"as-path-list",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyAsPathList) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "as-path-list"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyAsPathListDescrIPtion.IsNull() || o.LeafPolicyAsPathListDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafPolicyAsPathListDescrIPtion.ValueString()
	}

	// Tags
	if !(o.TagPolicyAsPathListRule.IsNull() || o.TagPolicyAsPathListRule.IsUnknown()) {
		subModel := make(map[string]PolicyAsPathListRule)
		diags.Append(o.TagPolicyAsPathListRule.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["rule"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyAsPathList) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "as-path-list"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafPolicyAsPathListDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyAsPathListDescrIPtion = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["rule"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: PolicyAsPathListRule{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagPolicyAsPathListRule = data
	} else {
		o.TagPolicyAsPathListRule = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "as-path-list"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyAsPathList) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description": types.StringType,

		// Tags
		"rule": types.MapType{ElemType: types.ObjectType{AttrTypes: PolicyAsPathListRule{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAsPathList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Add a BGP autonomous system path filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  AS path list name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		// TagNodes

		"rule": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: PolicyAsPathListRule{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Rule for this as-path-list

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  AS path list rule number  |

`,
		},

		// Nodes

	}
}
