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

// ProtocolsRIPngDistributeListInterface describes the resource data model.
type ProtocolsRIPngDistributeListInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsRIPngDistributeListInterfaceAccessList types.Object `tfsdk:"access_list"`
	NodeProtocolsRIPngDistributeListInterfacePrefixList types.Object `tfsdk:"prefix_list"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPngDistributeListInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"ripng",
		"distribute-list",
		"interface",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsRIPngDistributeListInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeProtocolsRIPngDistributeListInterfaceAccessList.IsNull() || o.NodeProtocolsRIPngDistributeListInterfaceAccessList.IsUnknown()) {
		var subModel ProtocolsRIPngDistributeListInterfaceAccessList
		diags.Append(o.NodeProtocolsRIPngDistributeListInterfaceAccessList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["access-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsRIPngDistributeListInterfacePrefixList.IsNull() || o.NodeProtocolsRIPngDistributeListInterfacePrefixList.IsUnknown()) {
		var subModel ProtocolsRIPngDistributeListInterfacePrefixList
		diags.Append(o.NodeProtocolsRIPngDistributeListInterfacePrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsRIPngDistributeListInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["access-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsRIPngDistributeListInterfaceAccessList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsRIPngDistributeListInterfaceAccessList = data

	} else {
		o.NodeProtocolsRIPngDistributeListInterfaceAccessList = basetypes.NewObjectNull(ProtocolsRIPngDistributeListInterfaceAccessList{}.AttributeTypes())
	}
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsRIPngDistributeListInterfacePrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsRIPngDistributeListInterfacePrefixList = data

	} else {
		o.NodeProtocolsRIPngDistributeListInterfacePrefixList = basetypes.NewObjectNull(ProtocolsRIPngDistributeListInterfacePrefixList{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "ripng", "distribute-list", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsRIPngDistributeListInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"access_list": types.ObjectType{AttrTypes: ProtocolsRIPngDistributeListInterfaceAccessList{}.AttributeTypes()},
		"prefix_list": types.ObjectType{AttrTypes: ProtocolsRIPngDistributeListInterfacePrefixList{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Apply filtering to an interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Apply filtering to an interface  |

`,
		},

		// LeafNodes

		// TagNodes

		// Nodes

		"access_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPngDistributeListInterfaceAccessList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPngDistributeListInterfacePrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list

`,
		},
	}
}
