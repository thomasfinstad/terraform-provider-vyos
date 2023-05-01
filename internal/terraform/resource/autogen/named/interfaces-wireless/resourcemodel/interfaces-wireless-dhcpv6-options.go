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

// InterfacesWirelessDhcpvsixOptions describes the resource data model.
type InterfacesWirelessDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesWirelessDhcpvsixOptionsDuID           types.String `tfsdk:"duid"`
	LeafInterfacesWirelessDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only"`
	LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit"`
	LeafInterfacesWirelessDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary"`

	// TagNodes
	TagInterfacesWirelessDhcpvsixOptionsPd types.Map `tfsdk:"pd"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessDhcpvsixOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "dhcpv6-options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessDhcpvsixOptionsDuID.IsNull() || o.LeafInterfacesWirelessDhcpvsixOptionsDuID.IsUnknown()) {
		vyosData["duid"] = o.LeafInterfacesWirelessDhcpvsixOptionsDuID.ValueString()
	}
	if !(o.LeafInterfacesWirelessDhcpvsixOptionsParametersOnly.IsNull() || o.LeafInterfacesWirelessDhcpvsixOptionsParametersOnly.IsUnknown()) {
		vyosData["parameters-only"] = o.LeafInterfacesWirelessDhcpvsixOptionsParametersOnly.ValueString()
	}
	if !(o.LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit.IsNull() || o.LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit.IsUnknown()) {
		vyosData["rapid-commit"] = o.LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit.ValueString()
	}
	if !(o.LeafInterfacesWirelessDhcpvsixOptionsTemporary.IsNull() || o.LeafInterfacesWirelessDhcpvsixOptionsTemporary.IsUnknown()) {
		vyosData["temporary"] = o.LeafInterfacesWirelessDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags
	if !(o.TagInterfacesWirelessDhcpvsixOptionsPd.IsNull() || o.TagInterfacesWirelessDhcpvsixOptionsPd.IsUnknown()) {
		subModel := make(map[string]InterfacesWirelessDhcpvsixOptionsPd)
		diags.Append(o.TagInterfacesWirelessDhcpvsixOptionsPd.ElementsAs(ctx, &subModel, false)...)

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
func (o *InterfacesWirelessDhcpvsixOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "dhcpv6-options"}})

	// Leafs
	if value, ok := vyosData["duid"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["parameters-only"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rapid-commit"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["temporary"]; ok {
		o.LeafInterfacesWirelessDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["pd"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesWirelessDhcpvsixOptionsPd{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesWirelessDhcpvsixOptionsPd = data
	} else {
		o.TagInterfacesWirelessDhcpvsixOptionsPd = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "dhcpv6-options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessDhcpvsixOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"duid":            types.StringType,
		"parameters_only": types.StringType,
		"rapid_commit":    types.StringType,
		"temporary":       types.StringType,

		// Tags
		"pd": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesWirelessDhcpvsixOptionsPd{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesWirelessDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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
