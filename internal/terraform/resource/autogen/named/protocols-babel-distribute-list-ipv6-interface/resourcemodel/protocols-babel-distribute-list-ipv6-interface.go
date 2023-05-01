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

// ProtocolsBabelDistributeListIPvsixInterface describes the resource data model.
type ProtocolsBabelDistributeListIPvsixInterface struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList types.Object `tfsdk:"access_list"`
	NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList types.Object `tfsdk:"prefix_list"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelDistributeListIPvsixInterface) GetVyosPath() []string {
	return []string{
		"protocols",
		"babel",
		"distribute-list",
		"ipv6",
		"interface",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBabelDistributeListIPvsixInterface) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv6", "interface"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList.IsNull() || o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList.IsUnknown()) {
		var subModel ProtocolsBabelDistributeListIPvsixInterfaceAccessList
		diags.Append(o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["access-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList.IsNull() || o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList.IsUnknown()) {
		var subModel ProtocolsBabelDistributeListIPvsixInterfacePrefixList
		diags.Append(o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBabelDistributeListIPvsixInterface) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv6", "interface"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["access-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList = data

	} else {
		o.NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList = basetypes.NewObjectNull(ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.AttributeTypes())
	}
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList = data

	} else {
		o.NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList = basetypes.NewObjectNull(ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "babel", "distribute-list", "ipv6", "interface"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBabelDistributeListIPvsixInterface) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"access_list": types.ObjectType{AttrTypes: ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.AttributeTypes()},
		"prefix_list": types.ObjectType{AttrTypes: ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvsixInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list

`,
		},
	}
}