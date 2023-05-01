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

// NatDestinationRuleTranSLAtion describes the resource data model.
type NatDestinationRuleTranSLAtion struct {
	// LeafNodes
	LeafNatDestinationRuleTranSLAtionAddress types.String `tfsdk:"address"`
	LeafNatDestinationRuleTranSLAtionPort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
	NodeNatDestinationRuleTranSLAtionOptions types.Object `tfsdk:"options"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *NatDestinationRuleTranSLAtion) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafNatDestinationRuleTranSLAtionAddress.IsNull() || o.LeafNatDestinationRuleTranSLAtionAddress.IsUnknown()) {
		vyosData["address"] = o.LeafNatDestinationRuleTranSLAtionAddress.ValueString()
	}
	if !(o.LeafNatDestinationRuleTranSLAtionPort.IsNull() || o.LeafNatDestinationRuleTranSLAtionPort.IsUnknown()) {
		vyosData["port"] = o.LeafNatDestinationRuleTranSLAtionPort.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeNatDestinationRuleTranSLAtionOptions.IsNull() || o.NodeNatDestinationRuleTranSLAtionOptions.IsUnknown()) {
		var subModel NatDestinationRuleTranSLAtionOptions
		diags.Append(o.NodeNatDestinationRuleTranSLAtionOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["options"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *NatDestinationRuleTranSLAtion) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafNatDestinationRuleTranSLAtionAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleTranSLAtionAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafNatDestinationRuleTranSLAtionPort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleTranSLAtionPort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, NatDestinationRuleTranSLAtionOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeNatDestinationRuleTranSLAtionOptions = data

	} else {
		o.NodeNatDestinationRuleTranSLAtionOptions = basetypes.NewObjectNull(NatDestinationRuleTranSLAtionOptions{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "translation"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o NatDestinationRuleTranSLAtion) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes
		"options": types.ObjectType{AttrTypes: NatDestinationRuleTranSLAtionOptions{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleTranSLAtion) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
|  range  |  Numbered port range (e.g., 1001-1005)  |

`,
		},

		// TagNodes

		// Nodes

		"options": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleTranSLAtionOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Translation options

`,
		},
	}
}
