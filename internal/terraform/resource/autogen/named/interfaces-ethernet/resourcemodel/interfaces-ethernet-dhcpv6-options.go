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

// InterfacesEthernetDhcpvsixOptions describes the resource data model.
type InterfacesEthernetDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesEthernetDhcpvsixOptionsDuID           types.String `tfsdk:"duid"`
	LeafInterfacesEthernetDhcpvsixOptionsParametersOnly types.String `tfsdk:"parameters_only"`
	LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit    types.String `tfsdk:"rapid_commit"`
	LeafInterfacesEthernetDhcpvsixOptionsTemporary      types.String `tfsdk:"temporary"`

	// TagNodes
	TagInterfacesEthernetDhcpvsixOptionsPd types.Map `tfsdk:"pd"`

	// Nodes
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesEthernetDhcpvsixOptions) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "dhcpv6-options"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesEthernetDhcpvsixOptionsDuID.IsNull() || o.LeafInterfacesEthernetDhcpvsixOptionsDuID.IsUnknown()) {
		vyosData["duid"] = o.LeafInterfacesEthernetDhcpvsixOptionsDuID.ValueString()
	}
	if !(o.LeafInterfacesEthernetDhcpvsixOptionsParametersOnly.IsNull() || o.LeafInterfacesEthernetDhcpvsixOptionsParametersOnly.IsUnknown()) {
		vyosData["parameters-only"] = o.LeafInterfacesEthernetDhcpvsixOptionsParametersOnly.ValueString()
	}
	if !(o.LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit.IsNull() || o.LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit.IsUnknown()) {
		vyosData["rapid-commit"] = o.LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit.ValueString()
	}
	if !(o.LeafInterfacesEthernetDhcpvsixOptionsTemporary.IsNull() || o.LeafInterfacesEthernetDhcpvsixOptionsTemporary.IsUnknown()) {
		vyosData["temporary"] = o.LeafInterfacesEthernetDhcpvsixOptionsTemporary.ValueString()
	}

	// Tags
	if !(o.TagInterfacesEthernetDhcpvsixOptionsPd.IsNull() || o.TagInterfacesEthernetDhcpvsixOptionsPd.IsUnknown()) {
		subModel := make(map[string]InterfacesEthernetDhcpvsixOptionsPd)
		diags.Append(o.TagInterfacesEthernetDhcpvsixOptionsPd.ElementsAs(ctx, &subModel, false)...)

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
func (o *InterfacesEthernetDhcpvsixOptions) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "dhcpv6-options"}})

	// Leafs
	if value, ok := vyosData["duid"]; ok {
		o.LeafInterfacesEthernetDhcpvsixOptionsDuID = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetDhcpvsixOptionsDuID = basetypes.NewStringNull()
	}
	if value, ok := vyosData["parameters-only"]; ok {
		o.LeafInterfacesEthernetDhcpvsixOptionsParametersOnly = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetDhcpvsixOptionsParametersOnly = basetypes.NewStringNull()
	}
	if value, ok := vyosData["rapid-commit"]; ok {
		o.LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetDhcpvsixOptionsRAPIDCommit = basetypes.NewStringNull()
	}
	if value, ok := vyosData["temporary"]; ok {
		o.LeafInterfacesEthernetDhcpvsixOptionsTemporary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesEthernetDhcpvsixOptionsTemporary = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["pd"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesEthernetDhcpvsixOptionsPd{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesEthernetDhcpvsixOptionsPd = data
	} else {
		o.TagInterfacesEthernetDhcpvsixOptionsPd = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "ethernet", "dhcpv6-options"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesEthernetDhcpvsixOptions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"duid":            types.StringType,
		"parameters_only": types.StringType,
		"rapid_commit":    types.StringType,
		"temporary":       types.StringType,

		// Tags
		"pd": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesEthernetDhcpvsixOptionsPd{}.AttributeTypes()}},

		// Nodes

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
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
				Attributes: InterfacesEthernetDhcpvsixOptionsPd{}.ResourceSchemaAttributes(),
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