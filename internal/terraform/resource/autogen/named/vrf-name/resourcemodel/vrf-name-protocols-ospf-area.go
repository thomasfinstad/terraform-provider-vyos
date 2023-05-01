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

// VrfNameProtocolsOspfArea describes the resource data model.
type VrfNameProtocolsOspfArea struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAuthentication types.String `tfsdk:"authentication"`
	LeafVrfNameProtocolsOspfAreaNetwork        types.String `tfsdk:"network"`
	LeafVrfNameProtocolsOspfAreaShortcut       types.String `tfsdk:"shortcut"`
	LeafVrfNameProtocolsOspfAreaExportList     types.String `tfsdk:"export_list"`
	LeafVrfNameProtocolsOspfAreaImportList     types.String `tfsdk:"import_list"`

	// TagNodes
	TagVrfNameProtocolsOspfAreaRange       types.Map `tfsdk:"range"`
	TagVrfNameProtocolsOspfAreaVirtualLink types.Map `tfsdk:"virtual_link"`

	// Nodes
	NodeVrfNameProtocolsOspfAreaAreaType types.Object `tfsdk:"area_type"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsOspfArea) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "area"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsOspfAreaAuthentication.IsNull() || o.LeafVrfNameProtocolsOspfAreaAuthentication.IsUnknown()) {
		vyosData["authentication"] = o.LeafVrfNameProtocolsOspfAreaAuthentication.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfAreaNetwork.IsNull() || o.LeafVrfNameProtocolsOspfAreaNetwork.IsUnknown()) {
		vyosData["network"] = o.LeafVrfNameProtocolsOspfAreaNetwork.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfAreaShortcut.IsNull() || o.LeafVrfNameProtocolsOspfAreaShortcut.IsUnknown()) {
		vyosData["shortcut"] = o.LeafVrfNameProtocolsOspfAreaShortcut.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfAreaExportList.IsNull() || o.LeafVrfNameProtocolsOspfAreaExportList.IsUnknown()) {
		vyosData["export-list"] = o.LeafVrfNameProtocolsOspfAreaExportList.ValueString()
	}
	if !(o.LeafVrfNameProtocolsOspfAreaImportList.IsNull() || o.LeafVrfNameProtocolsOspfAreaImportList.IsUnknown()) {
		vyosData["import-list"] = o.LeafVrfNameProtocolsOspfAreaImportList.ValueString()
	}

	// Tags
	if !(o.TagVrfNameProtocolsOspfAreaRange.IsNull() || o.TagVrfNameProtocolsOspfAreaRange.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsOspfAreaRange)
		diags.Append(o.TagVrfNameProtocolsOspfAreaRange.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["range"] = subData
	}
	if !(o.TagVrfNameProtocolsOspfAreaVirtualLink.IsNull() || o.TagVrfNameProtocolsOspfAreaVirtualLink.IsUnknown()) {
		subModel := make(map[string]VrfNameProtocolsOspfAreaVirtualLink)
		diags.Append(o.TagVrfNameProtocolsOspfAreaVirtualLink.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["virtual-link"] = subData
	}

	// Nodes
	if !(o.NodeVrfNameProtocolsOspfAreaAreaType.IsNull() || o.NodeVrfNameProtocolsOspfAreaAreaType.IsUnknown()) {
		var subModel VrfNameProtocolsOspfAreaAreaType
		diags.Append(o.NodeVrfNameProtocolsOspfAreaAreaType.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["area-type"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsOspfArea) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "area"}})

	// Leafs
	if value, ok := vyosData["authentication"]; ok {
		o.LeafVrfNameProtocolsOspfAreaAuthentication = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaAuthentication = basetypes.NewStringNull()
	}
	if value, ok := vyosData["network"]; ok {
		o.LeafVrfNameProtocolsOspfAreaNetwork = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaNetwork = basetypes.NewStringNull()
	}
	if value, ok := vyosData["shortcut"]; ok {
		o.LeafVrfNameProtocolsOspfAreaShortcut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaShortcut = basetypes.NewStringNull()
	}
	if value, ok := vyosData["export-list"]; ok {
		o.LeafVrfNameProtocolsOspfAreaExportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaExportList = basetypes.NewStringNull()
	}
	if value, ok := vyosData["import-list"]; ok {
		o.LeafVrfNameProtocolsOspfAreaImportList = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsOspfAreaImportList = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["range"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsOspfAreaRange{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsOspfAreaRange = data
	} else {
		o.TagVrfNameProtocolsOspfAreaRange = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["virtual-link"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: VrfNameProtocolsOspfAreaVirtualLink{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagVrfNameProtocolsOspfAreaVirtualLink = data
	} else {
		o.TagVrfNameProtocolsOspfAreaVirtualLink = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["area-type"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsOspfAreaAreaType{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsOspfAreaAreaType = data

	} else {
		o.NodeVrfNameProtocolsOspfAreaAreaType = basetypes.NewObjectNull(VrfNameProtocolsOspfAreaAreaType{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "ospf", "area"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsOspfArea) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"authentication": types.StringType,
		"network":        types.StringType,
		"shortcut":       types.StringType,
		"export_list":    types.StringType,
		"import_list":    types.StringType,

		// Tags
		"range":        types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsOspfAreaRange{}.AttributeTypes()}},
		"virtual_link": types.MapType{ElemType: types.ObjectType{AttrTypes: VrfNameProtocolsOspfAreaVirtualLink{}.AttributeTypes()}},

		// Nodes
		"area_type": types.ObjectType{AttrTypes: VrfNameProtocolsOspfAreaAreaType{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfArea) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"authentication": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF area authentication type

|  Format  |  Description  |
|----------|---------------|
|  plaintext-password  |  Use plain-text authentication  |
|  md5  |  Use MD5 authentication  |

`,
		},

		"network": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  OSPF network  |

`,
		},

		"shortcut": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Area shortcut mode

|  Format  |  Description  |
|----------|---------------|
|  default  |  Set default  |
|  disable  |  Disable shortcutting mode  |
|  enable  |  Enable shortcutting mode  |

`,
		},

		"export_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks announced to other areas

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access-list number  |

`,
		},

		"import_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks from other areas announced

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access-list number  |

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfAreaRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Summarize routes matching a prefix (border routers only)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Area range prefix  |

`,
		},

		"virtual_link": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: VrfNameProtocolsOspfAreaVirtualLink{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Virtual link

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  OSPF area in dotted decimal notation  |

`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaAreaType{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Area type

`,
		},
	}
}
