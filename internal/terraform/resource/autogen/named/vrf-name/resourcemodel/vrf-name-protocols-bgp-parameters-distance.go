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

// VrfNameProtocolsBgpParametersDistance describes the resource data model.
type VrfNameProtocolsBgpParametersDistance struct {
	// LeafNodes

	// TagNodes
	TagVrfNameProtocolsBgpParametersDistancePrefix types.Map `tfsdk:"prefix"`

	// Nodes
	NodeVrfNameProtocolsBgpParametersDistanceGlobal types.Object `tfsdk:"global"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpParametersDistance) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "distance"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags
	if !(o.TagVrfNameProtocolsBgpParametersDistancePrefix.IsNull() || o.TagVrfNameProtocolsBgpParametersDistancePrefix.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsBgpParametersDistancePrefix)
		diags.Append(o.TagVrfNameProtocolsBgpParametersDistancePrefix.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["prefix"] = subData
	}

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpParametersDistanceGlobal.IsNull() || o.NodeVrfNameProtocolsBgpParametersDistanceGlobal.IsUnknown()) {
		var subModel VrfNameProtocolsBgpParametersDistanceGlobal
		diags.Append(o.NodeVrfNameProtocolsBgpParametersDistanceGlobal.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["global"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpParametersDistance) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "distance"}})

	// Leafs

	// Tags
	if value, ok := vyosData["prefix"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsBgpParametersDistancePrefix{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsBgpParametersDistancePrefix = data
	} else {
		o.TagVrfNameProtocolsBgpParametersDistancePrefix = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["global"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpParametersDistanceGlobal{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpParametersDistanceGlobal = data

	} else {
		o.NodeVrfNameProtocolsBgpParametersDistanceGlobal = basetypes.NewObjectNull(VrfNameProtocolsBgpParametersDistanceGlobal{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "parameters", "distance"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpParametersDistance) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags
		"prefix": types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsBgpParametersDistancePrefix{}.AttributeTypes()}},

		// Nodes
		"global": types.ObjectType{AttrTypes: VrfNameProtocolsBgpParametersDistanceGlobal{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpParametersDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		"prefix": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsBgpParametersDistancePrefix{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |

`,
		},

		// Nodes

		"global": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDistanceGlobal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Global administratives distances for BGP routes

`,
		},
	}
}
