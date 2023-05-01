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

// InterfacesPseudoEthernetVifSVifCDhcpvsixOptions describes the resource data model.
type InterfacesPseudoEthernetVifSVifCDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID           types.String `tfsdk:"duid"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary"`

	// TagNodes
	TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd types.Map `tfsdk:"pd"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesPseudoEthernetVifSVifCDhcpvsixOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "pseudo-ethernet", "vif-s", "vif-c", "dhcpv6-options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID.IsNull() || o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID.IsUnknown()) {
		vyosData["duid"] = o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID.ValueString()
	}
	if !(o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly.IsNull() || o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly.IsUnknown()) {
		vyosData["parameters-only"] = o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly.ValueString()
	}
	if !(o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.IsNull() || o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.IsUnknown()) {
		vyosData["rapid-commit"] = o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit.ValueString()
	}
	if !(o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary.IsNull() || o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary.IsUnknown()) {
		vyosData["temporary"] = o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags
	if !(o.TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd.IsNull() || o.TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd.IsUnknown()) {
		subModel := make(map[string]InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd)
		diags.Append(o.TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd.ElementsAs(ctx, &subModel, false)...)

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
func (o *InterfacesPseudoEthernetVifSVifCDhcpvsixOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "pseudo-ethernet", "vif-s", "vif-c", "dhcpv6-options"}})

	// Leafs
	if value, ok := vyosData["duid"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["parameters-only"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rapid-commit"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["temporary"]; ok {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["pd"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd = data
	} else {
		o.TagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "pseudo-ethernet", "vif-s", "vif-c", "dhcpv6-options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCDhcpvsixOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"duid":            types.StringType,
		"parameters_only": types.StringType,
		"rapid_commit":    types.StringType,
		"temporary":       types.StringType,

		// Tags
		"pd": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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
