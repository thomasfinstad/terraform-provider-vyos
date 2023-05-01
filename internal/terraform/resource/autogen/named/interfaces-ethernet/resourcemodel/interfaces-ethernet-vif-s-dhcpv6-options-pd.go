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

// InterfacesEthernetVifSDhcpvsixOptionsPd describes the resource data model.
type InterfacesEthernetVifSDhcpvsixOptionsPd struct {
	// LeafNodes
	LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength types.String `tfsdk:"length"`

	// TagNodes
	TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface types.Map `tfsdk:"interface"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesEthernetVifSDhcpvsixOptionsPd) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "dhcpv6-options", "pd"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength.IsNull() || o.LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength.IsUnknown()) {
		vyosData["length"] = o.LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength.ValueString()
	}

	// Tags
	if !(o.TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface.IsNull() || o.TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface.IsUnknown()) {
		subModel := make(map[string]InterfacesEthernetVifSDhcpvsixOptionsPdInterface)
		diags.Append(o.TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface.ElementsAs(ctx, &subModel, false)...)

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
func (o *InterfacesEthernetVifSDhcpvsixOptionsPd) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "dhcpv6-options", "pd"}})

	// Leafs
	if value, ok := vyosData["length"]; ok {
		o.LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSDhcpvsixOptionsPdLength = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["interface"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesEthernetVifSDhcpvsixOptionsPdInterface{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface = data
	} else {
		o.TagInterfacesEthernetVifSDhcpvsixOptionsPdInterface = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "dhcpv6-options", "pd"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesEthernetVifSDhcpvsixOptionsPd) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"length": types.StringType,

		// Tags
		"interface": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesEthernetVifSDhcpvsixOptionsPdInterface{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSDhcpvsixOptionsPd) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesEthernetVifSDhcpvsixOptionsPdInterface{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
		},

		// Nodes

	}
}