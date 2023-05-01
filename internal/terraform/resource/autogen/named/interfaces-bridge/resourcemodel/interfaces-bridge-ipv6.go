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

// InterfacesBrIDgeIPvsix describes the resource data model.
type InterfacesBrIDgeIPvsix struct {
	// LeafNodes
	LeafInterfacesBrIDgeIPvsixAdjustMss              types.String `tfsdk:"adjust_mss"`
	LeafInterfacesBrIDgeIPvsixDisableForwarding      types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits types.String `tfsdk:"dup_addr_detect_transmits"`

	// TagNodes

	// Nodes
	NodeInterfacesBrIDgeIPvsixAddress types.Object `tfsdk:"address"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBrIDgeIPvsix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bridge", "ipv6"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBrIDgeIPvsixAdjustMss.IsNull() || o.LeafInterfacesBrIDgeIPvsixAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesBrIDgeIPvsixAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeIPvsixDisableForwarding.IsNull() || o.LeafInterfacesBrIDgeIPvsixDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesBrIDgeIPvsixDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits.IsNull() || o.LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits.IsUnknown()) {
		vyosData["dup-addr-detect-transmits"] = o.LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeInterfacesBrIDgeIPvsixAddress.IsNull() || o.NodeInterfacesBrIDgeIPvsixAddress.IsUnknown()) {
		var subModel InterfacesBrIDgeIPvsixAddress
		diags.Append(o.NodeInterfacesBrIDgeIPvsixAddress.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["address"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBrIDgeIPvsix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bridge", "ipv6"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesBrIDgeIPvsixAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeIPvsixAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesBrIDgeIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeIPvsixDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dup-addr-detect-transmits"]; ok {
		o.LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBrIDgeIPvsixDupAddrDetectTransmits = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["address"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBrIDgeIPvsixAddress{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBrIDgeIPvsixAddress = data

	} else {
		o.NodeInterfacesBrIDgeIPvsixAddress = basetypes.NewObjectNull(InterfacesBrIDgeIPvsixAddress{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bridge", "ipv6"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBrIDgeIPvsix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"adjust_mss":                types.StringType,
		"disable_forwarding":        types.StringType,
		"dup_addr_detect_transmits": types.StringType,

		// Tags

		// Nodes
		"address": types.ObjectType{AttrTypes: InterfacesBrIDgeIPvsixAddress{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBrIDgeIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |

`,
		},

		"disable_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
		},

		"dup_addr_detect_transmits": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
|  u32:1-n  |  Number of NS messages to send while performing DAD  |

`,
		},

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesBrIDgeIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}
