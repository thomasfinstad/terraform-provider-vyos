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

// NatDestinationRuleSource describes the resource data model.
type NatDestinationRuleSource struct {
	// LeafNodes
	LeafNatDestinationRuleSourceAddress types.String `tfsdk:"address"`
	LeafNatDestinationRuleSourcePort    types.String `tfsdk:"port"`

	// TagNodes

	// Nodes
	NodeNatDestinationRuleSourceGroup types.Object `tfsdk:"group"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *NatDestinationRuleSource) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "source"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafNatDestinationRuleSourceAddress.IsNull() || o.LeafNatDestinationRuleSourceAddress.IsUnknown()) {
		vyosData["address"] = o.LeafNatDestinationRuleSourceAddress.ValueString()
	}
	if !(o.LeafNatDestinationRuleSourcePort.IsNull() || o.LeafNatDestinationRuleSourcePort.IsUnknown()) {
		vyosData["port"] = o.LeafNatDestinationRuleSourcePort.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeNatDestinationRuleSourceGroup.IsNull() || o.NodeNatDestinationRuleSourceGroup.IsUnknown()) {
		var subModel NatDestinationRuleSourceGroup
		diags.Append(o.NodeNatDestinationRuleSourceGroup.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["group"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *NatDestinationRuleSource) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "source"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafNatDestinationRuleSourceAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleSourceAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["port"]; ok {
		o.LeafNatDestinationRuleSourcePort = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleSourcePort = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["group"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, NatDestinationRuleSourceGroup{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeNatDestinationRuleSourceGroup = data

	} else {
		o.NodeNatDestinationRuleSourceGroup = basetypes.NewObjectNull(NatDestinationRuleSourceGroup{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"nat", "destination", "rule", "source"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o NatDestinationRuleSource) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address": types.StringType,
		"port":    types.StringType,

		// Tags

		// Nodes
		"group": types.ObjectType{AttrTypes: NatDestinationRuleSourceGroup{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleSource) ResourceSchemaAttributes() map[string]schema.Attribute {
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
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numeric IP port  |
|  start-end  |  Numbered port range (e.g. 1001-1005)  |
|   |   |

`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleSourceGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
