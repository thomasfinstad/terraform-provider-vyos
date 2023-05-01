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

// InterfacesWirelessIPvsix describes the resource data model.
type InterfacesWirelessIPvsix struct {
	// LeafNodes
	LeafInterfacesWirelessIPvsixAdjustMss              types.String `tfsdk:"adjust_mss"`
	LeafInterfacesWirelessIPvsixDisableForwarding      types.String `tfsdk:"disable_forwarding"`
	LeafInterfacesWirelessIPvsixDupAddrDetectTransmits types.String `tfsdk:"dup_addr_detect_transmits"`

	// TagNodes

	// Nodes
	NodeInterfacesWirelessIPvsixAddress types.Object `tfsdk:"address"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessIPvsix) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessIPvsixAdjustMss.IsNull() || o.LeafInterfacesWirelessIPvsixAdjustMss.IsUnknown()) {
		vyosData["adjust-mss"] = o.LeafInterfacesWirelessIPvsixAdjustMss.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPvsixDisableForwarding.IsNull() || o.LeafInterfacesWirelessIPvsixDisableForwarding.IsUnknown()) {
		vyosData["disable-forwarding"] = o.LeafInterfacesWirelessIPvsixDisableForwarding.ValueString()
	}
	if !(o.LeafInterfacesWirelessIPvsixDupAddrDetectTransmits.IsNull() || o.LeafInterfacesWirelessIPvsixDupAddrDetectTransmits.IsUnknown()) {
		vyosData["dup-addr-detect-transmits"] = o.LeafInterfacesWirelessIPvsixDupAddrDetectTransmits.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeInterfacesWirelessIPvsixAddress.IsNull() || o.NodeInterfacesWirelessIPvsixAddress.IsUnknown()) {
		var subModel InterfacesWirelessIPvsixAddress
		diags.Append(o.NodeInterfacesWirelessIPvsixAddress.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["address"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessIPvsix) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6"}})

	// Leafs
	if value, ok := vyosData["adjust-mss"]; ok {
		o.LeafInterfacesWirelessIPvsixAdjustMss = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixAdjustMss = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-forwarding"]; ok {
		o.LeafInterfacesWirelessIPvsixDisableForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixDisableForwarding = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dup-addr-detect-transmits"]; ok {
		o.LeafInterfacesWirelessIPvsixDupAddrDetectTransmits = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessIPvsixDupAddrDetectTransmits = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["address"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessIPvsixAddress{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessIPvsixAddress = data

	} else {
		o.NodeInterfacesWirelessIPvsixAddress = basetypes.NewObjectNull(InterfacesWirelessIPvsixAddress{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "ipv6"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessIPvsix) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"adjust_mss":                types.StringType,
		"disable_forwarding":        types.StringType,
		"dup_addr_detect_transmits": types.StringType,

		// Tags

		// Nodes
		"address": types.ObjectType{AttrTypes: InterfacesWirelessIPvsixAddress{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: InterfacesWirelessIPvsixAddress{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
		},
	}
}