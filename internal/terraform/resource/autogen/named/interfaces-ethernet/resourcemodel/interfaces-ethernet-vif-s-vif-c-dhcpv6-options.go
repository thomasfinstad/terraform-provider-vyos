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

// InterfacesEthernetVifSVifCDhcpvsixOptions describes the resource data model.
type InterfacesEthernetVifSVifCDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID           types.String `tfsdk:"duid"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary"`

	// TagNodes
	TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd types.Map `tfsdk:"pd"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesEthernetVifSVifCDhcpvsixOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "vif-c", "dhcpv6-options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID.IsNull() || o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID.IsUnknown()) {
		vyosData["duid"] = o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID.ValueString()
	}
	if !(o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly.IsNull() || o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly.IsUnknown()) {
		vyosData["parameters-only"] = o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly.ValueString()
	}
	if !(o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.IsNull() || o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.IsUnknown()) {
		vyosData["rapid-commit"] = o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.ValueString()
	}
	if !(o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary.IsNull() || o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary.IsUnknown()) {
		vyosData["temporary"] = o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags
	if !(o.TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd.IsNull() || o.TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd.IsUnknown()) {
		subModel := make(map[string]InterfacesEthernetVifSVifCDhcpvsixOptionsPd)
		diags.Append(o.TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["pd"] = subData
	}

	// Nodes

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesEthernetVifSVifCDhcpvsixOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "vif-c", "dhcpv6-options"}})

	// Leafs
	if value, ok := vyosData["duid"]; ok {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["parameters-only"]; ok {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rapid-commit"]; ok {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["temporary"]; ok {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["pd"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesEthernetVifSVifCDhcpvsixOptionsPd{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd = data
	} else {
		o.TagInterfacesEthernetVifSVifCDhcpvsixOptionsPd = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "vif-s", "vif-c", "dhcpv6-options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesEthernetVifSVifCDhcpvsixOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"duid":            types.StringType,
		"parameters_only": types.StringType,
		"rapid_commit":    types.StringType,
		"temporary":       types.StringType,

		// Tags
		"pd": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesEthernetVifSVifCDhcpvsixOptionsPd{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSVifCDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

|  Format  |  Description  |
|----------|---------------|
|  duid  |  DHCP unique identifier (DUID)  |

`,
		},

		"parameters_only": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
		},

		"rapid_commit": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
		},

		"temporary": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 temporary address

`,
		},

		// TagNodes

		"pd": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesEthernetVifSVifCDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCPv6 prefix delegation interface statement

|  Format  |  Description  |
|----------|---------------|
|  instance number  |  Prefix delegation instance (>= 0)  |

`,
		},

		// Nodes

	}
}