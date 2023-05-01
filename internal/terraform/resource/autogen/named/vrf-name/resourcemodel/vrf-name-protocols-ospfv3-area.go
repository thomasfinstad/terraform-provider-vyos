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

// VrfNameProtocolsOspfvthreeArea describes the resource data model.
type VrfNameProtocolsOspfvthreeArea struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaExportList types.String `tfsdk:"export_list"`
	LeafVrfNameProtocolsOspfvthreeAreaImportList types.String `tfsdk:"import_list"`

	// TagNodes
	TagVrfNameProtocolsOspfvthreeAreaRange types.Map `tfsdk:"range"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAreaAreaType types.Object `tfsdk:"area_type"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfvthreeArea) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "area"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfvthreeAreaExportList.IsNull() || o.LeafVrfNameProtocolsOspfvthreeAreaExportList.IsUnknown()) {
		vyosData["export-list"] = o.LeafVrfNameProtocolsOspfvthreeAreaExportList.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfvthreeAreaImportList.IsNull() || o.LeafVrfNameProtocolsOspfvthreeAreaImportList.IsUnknown()) {
		vyosData["import-list"] = o.LeafVrfNameProtocolsOspfvthreeAreaImportList.ValueString()
	}

	// Tags
	if !(o.TagVrfNameProtocolsOspfvthreeAreaRange.IsNull() || o.TagVrfNameProtocolsOspfvthreeAreaRange.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsOspfvthreeAreaRange)
		diags.Append(o.TagVrfNameProtocolsOspfvthreeAreaRange.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["range"] = subData
	}

	// Nodes
	if !(o.NodeVrfNameProtocolsOspfvthreeAreaAreaType.IsNull() || o.NodeVrfNameProtocolsOspfvthreeAreaAreaType.IsUnknown()) {
		var subModel VrfNameProtocolsOspfvthreeAreaAreaType
		diags.Append(o.NodeVrfNameProtocolsOspfvthreeAreaAreaType.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["area-type"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfvthreeArea) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "area"}})

	// Leafs
	if value, ok := vyosData["export-list"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeAreaExportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeAreaExportList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import-list"]; ok {
		o.LeafVrfNameProtocolsOspfvthreeAreaImportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfvthreeAreaImportList = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["range"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeAreaRange{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsOspfvthreeAreaRange = data
	} else {
		o.TagVrfNameProtocolsOspfvthreeAreaRange = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["area-type"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfvthreeAreaAreaType{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfvthreeAreaAreaType = data

	} else {
		o.NodeVrfNameProtocolsOspfvthreeAreaAreaType = basetypes.NewObjectNull(VrfNameProtocolsOspfvthreeAreaAreaType{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospfv3", "area"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfvthreeArea) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"export_list": types.StringType,
		"import_list": types.StringType,

		// Tags
		"range": types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeAreaRange{}.AttributeTypes()}},

		// Nodes
		"area_type": types.ObjectType{AttrTypes: VrfNameProtocolsOspfvthreeAreaAreaType{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeArea) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of export-list

`,
		},

		"import_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of import-list

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfvthreeAreaRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Specify IPv6 prefix (border routers only)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Specify IPv6 prefix (border routers only)  |

`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAreaAreaType{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `OSPFv3 Area type

`,
		},
	}
}