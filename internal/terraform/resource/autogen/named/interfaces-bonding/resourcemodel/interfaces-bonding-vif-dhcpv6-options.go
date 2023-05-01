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

// InterfacesBondingVifDhcpvsixOptions describes the resource data model.
type InterfacesBondingVifDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesBondingVifDhcpvsixOptionsDuID           types.String `tfsdk:"duid"`
	LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only"`
	LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit"`
	LeafInterfacesBondingVifDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary"`

	// TagNodes
	TagInterfacesBondingVifDhcpvsixOptionsPd types.Map `tfsdk:"pd"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBondingVifDhcpvsixOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif", "dhcpv6-options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBondingVifDhcpvsixOptionsDuID.IsNull() || o.LeafInterfacesBondingVifDhcpvsixOptionsDuID.IsUnknown()) {
		vyosData["duid"] = o.LeafInterfacesBondingVifDhcpvsixOptionsDuID.ValueString()
	}
	if !(o.LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly.IsNull() || o.LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly.IsUnknown()) {
		vyosData["parameters-only"] = o.LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly.ValueString()
	}
	if !(o.LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit.IsNull() || o.LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit.IsUnknown()) {
		vyosData["rapid-commit"] = o.LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit.ValueString()
	}
	if !(o.LeafInterfacesBondingVifDhcpvsixOptionsTemporary.IsNull() || o.LeafInterfacesBondingVifDhcpvsixOptionsTemporary.IsUnknown()) {
		vyosData["temporary"] = o.LeafInterfacesBondingVifDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags
	if !(o.TagInterfacesBondingVifDhcpvsixOptionsPd.IsNull() || o.TagInterfacesBondingVifDhcpvsixOptionsPd.IsUnknown()) {
		subModel := make(map[string]InterfacesBondingVifDhcpvsixOptionsPd)
		diags.Append(o.TagInterfacesBondingVifDhcpvsixOptionsPd.ElementsAs(ctx, &subModel, false)...)

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
func (o *InterfacesBondingVifDhcpvsixOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif", "dhcpv6-options"}})

	// Leafs
	if value, ok := vyosData["duid"]; ok {
		o.LeafInterfacesBondingVifDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["parameters-only"]; ok {
		o.LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rapid-commit"]; ok {
		o.LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["temporary"]; ok {
		o.LeafInterfacesBondingVifDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVifDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["pd"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesBondingVifDhcpvsixOptionsPd{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesBondingVifDhcpvsixOptionsPd = data
	} else {
		o.TagInterfacesBondingVifDhcpvsixOptionsPd = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bonding", "vif", "dhcpv6-options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBondingVifDhcpvsixOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"duid":            types.StringType,
		"parameters_only": types.StringType,
		"rapid_commit":    types.StringType,
		"temporary":       types.StringType,

		// Tags
		"pd": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesBondingVifDhcpvsixOptionsPd{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesBondingVifDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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