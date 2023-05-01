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

// InterfacesMacsecDhcpvsixOptionsPd describes the resource data model.
type InterfacesMacsecDhcpvsixOptionsPd struct {
	// LeafNodes
	LeafInterfacesMacsecDhcpvsixOptionsPdLength types.String `tfsdk:"length"`

	// TagNodes
	TagInterfacesMacsecDhcpvsixOptionsPdInterface types.Map `tfsdk:"interface"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesMacsecDhcpvsixOptionsPd) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "macsec", "dhcpv6-options", "pd"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesMacsecDhcpvsixOptionsPdLength.IsNull() || o.LeafInterfacesMacsecDhcpvsixOptionsPdLength.IsUnknown()) {
		vyosData["length"] = o.LeafInterfacesMacsecDhcpvsixOptionsPdLength.ValueString()
	}

	// Tags
	if !(o.TagInterfacesMacsecDhcpvsixOptionsPdInterface.IsNull() || o.TagInterfacesMacsecDhcpvsixOptionsPdInterface.IsUnknown()) {
		subModel := make(map[string]InterfacesMacsecDhcpvsixOptionsPdInterface)
		diags.Append(o.TagInterfacesMacsecDhcpvsixOptionsPdInterface.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["interface"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesMacsecDhcpvsixOptionsPd) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "macsec", "dhcpv6-options", "pd"}})

	// Leafs
	if value, ok := vyosData["length"]; ok {
		o.LeafInterfacesMacsecDhcpvsixOptionsPdLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDhcpvsixOptionsPdLength = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["interface"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesMacsecDhcpvsixOptionsPdInterface{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesMacsecDhcpvsixOptionsPdInterface = data
	} else {
		o.TagInterfacesMacsecDhcpvsixOptionsPdInterface = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "macsec", "dhcpv6-options", "pd"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesMacsecDhcpvsixOptionsPd) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"length": types.StringType,

		// Tags
		"interface": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesMacsecDhcpvsixOptionsPdInterface{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

|  Format  |  Description  |
|----------|---------------|
|  u32:32-64  |  Length of delegated prefix  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		"interface": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesMacsecDhcpvsixOptionsPdInterface{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		// Nodes

	}
}