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

// QosPolicyPriorityQueueClassMatch describes the resource data model.
type QosPolicyPriorityQueueClassMatch struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchDescrIPtion types.String `tfsdk:"description"`
	LeafQosPolicyPriorityQueueClassMatchInterface   types.String `tfsdk:"interface"`
	LeafQosPolicyPriorityQueueClassMatchMark        types.String `tfsdk:"mark"`
	LeafQosPolicyPriorityQueueClassMatchVif         types.String `tfsdk:"vif"`

	// TagNodes

	// Nodes
	NodeQosPolicyPriorityQueueClassMatchEther  types.Object `tfsdk:"ether"`
	NodeQosPolicyPriorityQueueClassMatchIP     types.Object `tfsdk:"ip"`
	NodeQosPolicyPriorityQueueClassMatchIPvsix types.Object `tfsdk:"ipv6"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *QosPolicyPriorityQueueClassMatch) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafQosPolicyPriorityQueueClassMatchDescrIPtion.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafQosPolicyPriorityQueueClassMatchDescrIPtion.ValueString()
	}
	if !(o.LeafQosPolicyPriorityQueueClassMatchInterface.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchInterface.IsUnknown()) {
		vyosData["interface"] = o.LeafQosPolicyPriorityQueueClassMatchInterface.ValueString()
	}
	if !(o.LeafQosPolicyPriorityQueueClassMatchMark.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchMark.IsUnknown()) {
		vyosData["mark"] = o.LeafQosPolicyPriorityQueueClassMatchMark.ValueString()
	}
	if !(o.LeafQosPolicyPriorityQueueClassMatchVif.IsNull() || o.LeafQosPolicyPriorityQueueClassMatchVif.IsUnknown()) {
		vyosData["vif"] = o.LeafQosPolicyPriorityQueueClassMatchVif.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeQosPolicyPriorityQueueClassMatchEther.IsNull() || o.NodeQosPolicyPriorityQueueClassMatchEther.IsUnknown()) {
		var subModel QosPolicyPriorityQueueClassMatchEther
		diags.Append(o.NodeQosPolicyPriorityQueueClassMatchEther.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ether"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyPriorityQueueClassMatchIP.IsNull() || o.NodeQosPolicyPriorityQueueClassMatchIP.IsUnknown()) {
		var subModel QosPolicyPriorityQueueClassMatchIP
		diags.Append(o.NodeQosPolicyPriorityQueueClassMatchIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeQosPolicyPriorityQueueClassMatchIPvsix.IsNull() || o.NodeQosPolicyPriorityQueueClassMatchIPvsix.IsUnknown()) {
		var subModel QosPolicyPriorityQueueClassMatchIPvsix
		diags.Append(o.NodeQosPolicyPriorityQueueClassMatchIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *QosPolicyPriorityQueueClassMatch) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["interface"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mark"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vif"]; ok {
		o.LeafQosPolicyPriorityQueueClassMatchVif = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafQosPolicyPriorityQueueClassMatchVif = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["ether"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyPriorityQueueClassMatchEther{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyPriorityQueueClassMatchEther = data

	} else {
		o.NodeQosPolicyPriorityQueueClassMatchEther = basetypes.NewObjectNull(QosPolicyPriorityQueueClassMatchEther{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyPriorityQueueClassMatchIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyPriorityQueueClassMatchIP = data

	} else {
		o.NodeQosPolicyPriorityQueueClassMatchIP = basetypes.NewObjectNull(QosPolicyPriorityQueueClassMatchIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, QosPolicyPriorityQueueClassMatchIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeQosPolicyPriorityQueueClassMatchIPvsix = data

	} else {
		o.NodeQosPolicyPriorityQueueClassMatchIPvsix = basetypes.NewObjectNull(QosPolicyPriorityQueueClassMatchIPvsix{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"qos", "policy", "priority-queue", "class", "match"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o QosPolicyPriorityQueueClassMatch) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description": types.StringType,
		"interface":   types.StringType,
		"mark":        types.StringType,
		"vif":         types.StringType,

		// Tags

		// Nodes
		"ether": types.ObjectType{AttrTypes: QosPolicyPriorityQueueClassMatchEther{}.AttributeTypes()},
		"ip":    types.ObjectType{AttrTypes: QosPolicyPriorityQueueClassMatchIP{}.AttributeTypes()},
		"ipv6":  types.ObjectType{AttrTypes: QosPolicyPriorityQueueClassMatchIPvsix{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: QosPolicyPriorityQueueClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}