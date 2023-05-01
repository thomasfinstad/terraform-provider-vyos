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

// VrfNameProtocolsOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthree struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeRouteMap types.String `tfsdk:"route_map"`

	// TagNodes
	TagVrfNameProtocolsOspfvthreeArea      types.Map `tfsdk:"area"`
	TagVrfNameProtocolsOspfvthreeInterface types.Map `tfsdk:"interface"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAutoCost            types.Object `tfsdk:"auto_cost"`
	NodeVrfNameProtocolsOspfvthreeDefaultInformation  types.Object `tfsdk:"default_information"`
	NodeVrfNameProtocolsOspfvthreeDistance            types.Object `tfsdk:"distance"`
	NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges types.Object `tfsdk:"log_adjacency_changes"`
	NodeVrfNameProtocolsOspfvthreeParameters          types.Object `tfsdk:"parameters"`
	NodeVrfNameProtocolsOspfvthreeRedistribute        types.Object `tfsdk:"redistribute"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfvthree) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfvthreeRouteMap.IsNull() || o.LeafVrfNameProtocolsOspfvthreeRouteMap.IsUnknown()) {
		vyosData["route-map"] = o.LeafVrfNameProtocolsOspfvthreeRouteMap.ValueString()
	}

	// Tags
	if !(o.TagVrfNameProtocolsOspfvthreeArea.IsNull() || o.TagVrfNameProtocolsOspfvthreeArea.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsOspfvthreeArea)
		diags.Append(o.TagVrfNameProtocolsOspfvthreeArea.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["area"] = subData
	}
	if !(o.TagVrfNameProtocolsOspfvthreeInterface.IsNull() || o.TagVrfNameProtocolsOspfvthreeInterface.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsOspfvthreeInterface)
		diags.Append(o.TagVrfNameProtocolsOspfvthreeInterface.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["interface"] = subData
	}

	// Nodes
	if !(o.NodeVrfNameProtocolsOspfvthreeAutoCost.IsNull() || o.NodeVrfNameProtocolsOspfvthreeAutoCost.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeAutoCost
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeAutoCost.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["auto-cost"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsOspfvthreeDefaultInformation.IsNull() || o.NodeVrfNameProtocolsOspfvthreeDefaultInformation.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeDefaultInformation
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeDefaultInformation.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["default-information"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsOspfvthreeDistance.IsNull() || o.NodeVrfNameProtocolsOspfvthreeDistance.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeDistance
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeDistance.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["distance"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges.IsNull() || o.NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeLogAdjacencyChanges
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["log-adjacency-changes"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsOspfvthreeParameters.IsNull() || o.NodeVrfNameProtocolsOspfvthreeParameters.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeParameters
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeParameters.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["parameters"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsOspfvthreeRedistribute.IsNull() || o.NodeVrfNameProtocolsOspfvthreeRedistribute.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeRedistribute
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeRedistribute.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["redistribute"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfvthree) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3"}})

	// Leafs
	if value, ok := vyosData["route-map"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeRouteMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeRouteMap = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["area"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeArea{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsOspfvthreeArea = data
	} else {
		o.TagVrfNameProtocolsOspfvthreeArea = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["interface"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeInterface{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsOspfvthreeInterface = data
	} else {
		o.TagVrfNameProtocolsOspfvthreeInterface = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["auto-cost"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeAutoCost{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeAutoCost = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeAutoCost = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeAutoCost{}.AttributeTypes())
	}
	if value, ok := vyosData["default-information"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeDefaultInformation{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeDefaultInformation = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeDefaultInformation = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeDefaultInformation{}.AttributeTypes())
	}
	if value, ok := vyosData["distance"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeDistance{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeDistance = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeDistance = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeDistance{}.AttributeTypes())
	}
	if value, ok := vyosData["log-adjacency-changes"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.AttributeTypes())
	}
	if value, ok := vyosData["parameters"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeParameters{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeParameters = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeParameters = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeParameters{}.AttributeTypes())
	}
	if value, ok := vyosData["redistribute"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeRedistribute{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeRedistribute = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeRedistribute = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeRedistribute{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfvthree) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"route_map": types.StringType,

		// Tags
		"area":      types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeArea{}.AttributeTypes()}},
		"interface": types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeInterface{}.AttributeTypes()}},

		// Nodes
		"auto_cost":             types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeAutoCost{}.AttributeTypes()},
		"default_information":   types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeDefaultInformation{}.AttributeTypes()},
		"distance":              types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeDistance{}.AttributeTypes()},
		"log_adjacency_changes": types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.AttributeTypes()},
		"parameters":            types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeParameters{}.AttributeTypes()},
		"redistribute":          types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeRedistribute{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthree) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		// TagNodes

		"area": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfvthreeArea{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `OSPFv3 Area

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Area ID as a decimal value  |
|  ipv4  |  Area ID in IP address forma  |

`,
		},

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfvthreeInterface{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Enable routing on an IPv6 interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface used for routing information exchange  |

`,
		},

		// Nodes

		"auto_cost": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAutoCost{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Calculate interface cost according to bandwidth

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDefaultInformation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Default route advertisment settings

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDistance{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distance

`,
		},

		"log_adjacency_changes": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log adjacency state changes

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeParameters{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `OSPFv3 specific parameters

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistribute{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
		},
	}
}