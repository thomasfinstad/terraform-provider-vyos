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

// QosPolicyShaperClassMatch describes the resource data model.
type QosPolicyShaperClassMatch struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchDescrIPtion types.String `tfsdk:"description"`
	LeafQosPolicyShaperClassMatchInterface   types.String `tfsdk:"interface"`
	LeafQosPolicyShaperClassMatchMark        types.String `tfsdk:"mark"`
	LeafQosPolicyShaperClassMatchVif         types.String `tfsdk:"vif"`

	// TagNodes

	// Nodes
	NodeQosPolicyShaperClassMatchEther  types.Object `tfsdk:"ether"`
	NodeQosPolicyShaperClassMatchIP     types.Object `tfsdk:"ip"`
	NodeQosPolicyShaperClassMatchIPvsix types.Object `tfsdk:"ipv6"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyShaperClassMatch) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyShaperClassMatchDescrIPtion.IsNull() || o.LeafQosPolicyShaperClassMatchDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafQosPolicyShaperClassMatchDescrIPtion.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchInterface.IsNull() || o.LeafQosPolicyShaperClassMatchInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafQosPolicyShaperClassMatchInterface.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchMark.IsNull() || o.LeafQosPolicyShaperClassMatchMark.IsUnknown()) {
		vyosData["mark"] = o.LeafQosPolicyShaperClassMatchMark.ValueString()
	}
	if !(o.LeafQosPolicyShaperClassMatchVif.IsNull() || o.LeafQosPolicyShaperClassMatchVif.IsUnknown()) {
		vyosData["vif"] = o.LeafQosPolicyShaperClassMatchVif.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeQosPolicyShaperClassMatchEther.IsNull() || o.NodeQosPolicyShaperClassMatchEther.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchEther
		diags.Append(o.NodeQosPolicyShaperClassMatchEther.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ether"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperClassMatchIP.IsNull() || o.NodeQosPolicyShaperClassMatchIP.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchIP
		diags.Append(o.NodeQosPolicyShaperClassMatchIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyShaperClassMatchIPvsix.IsNull() || o.NodeQosPolicyShaperClassMatchIPvsix.IsUnknown()) {
		var subModel QosPolicyShaperClassMatchIPvsix
		diags.Append(o.NodeQosPolicyShaperClassMatchIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyShaperClassMatch) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafQosPolicyShaperClassMatchDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface"]; ok {
		o.LeafQosPolicyShaperClassMatchInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mark"]; ok {
		o.LeafQosPolicyShaperClassMatchMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vif"]; ok {
		o.LeafQosPolicyShaperClassMatchVif = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyShaperClassMatchVif = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["ether"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchEther{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchEther = data

	} else {
		o.NodeQosPolicyShaperClassMatchEther = basetypes.NewObjectNull(QosPolicyShaperClassMatchEther{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchIP = data

	} else {
		o.NodeQosPolicyShaperClassMatchIP = basetypes.NewObjectNull(QosPolicyShaperClassMatchIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyShaperClassMatchIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyShaperClassMatchIPvsix = data

	} else {
		o.NodeQosPolicyShaperClassMatchIPvsix = basetypes.NewObjectNull(QosPolicyShaperClassMatchIPvsix{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "shaper", "class", "match"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyShaperClassMatch) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description": types.StringType,
		"interface":   types.StringType,
		"mark":        types.StringType,
		"vif":         types.StringType,

		// Tags

		// Nodes
		"ether": types.ObjectType{AttrTypes: QosPolicyShaperClassMatchEther{}.AttributeTypes()},
		"ip":    types.ObjectType{AttrTypes: QosPolicyShaperClassMatchIP{}.AttributeTypes()},
		"ipv6":  types.ObjectType{AttrTypes: QosPolicyShaperClassMatchIPvsix{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on mark applied by firewall

|  Format  |  Description  |
|----------|---------------|
|  u32  |  FW mark to match  |

`,
		},

		"vif": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4095  |  Virtual Local Area Network (VLAN) tag   |

`,
		},

		// TagNodes

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}