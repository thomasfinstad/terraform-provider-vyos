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

// InterfacesWirelessVifSVifC describes the resource data model.
type InterfacesWirelessVifSVifC struct {
	// LeafNodes
	LeafInterfacesWirelessVifSVifCDescrIPtion       types.String `tfsdk:"description"`
	LeafInterfacesWirelessVifSVifCAddress           types.String `tfsdk:"address"`
	LeafInterfacesWirelessVifSVifCDisableLinkDetect types.String `tfsdk:"disable_link_detect"`
	LeafInterfacesWirelessVifSVifCDisable           types.String `tfsdk:"disable"`
	LeafInterfacesWirelessVifSVifCMac               types.String `tfsdk:"mac"`
	LeafInterfacesWirelessVifSVifCMtu               types.String `tfsdk:"mtu"`
	LeafInterfacesWirelessVifSVifCRedirect          types.String `tfsdk:"redirect"`
	LeafInterfacesWirelessVifSVifCVrf               types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
	NodeInterfacesWirelessVifSVifCDhcpOptions     types.Object `tfsdk:"dhcp_options"`
	NodeInterfacesWirelessVifSVifCDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options"`
	NodeInterfacesWirelessVifSVifCIP              types.Object `tfsdk:"ip"`
	NodeInterfacesWirelessVifSVifCIPvsix          types.Object `tfsdk:"ipv6"`
	NodeInterfacesWirelessVifSVifCMirror          types.Object `tfsdk:"mirror"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWirelessVifSVifC) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWirelessVifSVifCDescrIPtion.IsNull() || o.LeafInterfacesWirelessVifSVifCDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafInterfacesWirelessVifSVifCDescrIPtion.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCAddress.IsNull() || o.LeafInterfacesWirelessVifSVifCAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesWirelessVifSVifCAddress.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCDisableLinkDetect.IsNull() || o.LeafInterfacesWirelessVifSVifCDisableLinkDetect.IsUnknown()) {
		vyosData["disable-link-detect"] = o.LeafInterfacesWirelessVifSVifCDisableLinkDetect.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCDisable.IsNull() || o.LeafInterfacesWirelessVifSVifCDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesWirelessVifSVifCDisable.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCMac.IsNull() || o.LeafInterfacesWirelessVifSVifCMac.IsUnknown()) {
		vyosData["mac"] = o.LeafInterfacesWirelessVifSVifCMac.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCMtu.IsNull() || o.LeafInterfacesWirelessVifSVifCMtu.IsUnknown()) {
		vyosData["mtu"] = o.LeafInterfacesWirelessVifSVifCMtu.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCRedirect.IsNull() || o.LeafInterfacesWirelessVifSVifCRedirect.IsUnknown()) {
		vyosData["redirect"] = o.LeafInterfacesWirelessVifSVifCRedirect.ValueString()
	}
	if !(o.LeafInterfacesWirelessVifSVifCVrf.IsNull() || o.LeafInterfacesWirelessVifSVifCVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafInterfacesWirelessVifSVifCVrf.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeInterfacesWirelessVifSVifCDhcpOptions.IsNull() || o.NodeInterfacesWirelessVifSVifCDhcpOptions.IsUnknown()) {
		var subModel InterfacesWirelessVifSVifCDhcpOptions
		diags.Append(o.NodeInterfacesWirelessVifSVifCDhcpOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcp-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWirelessVifSVifCDhcpvsixOptions.IsNull() || o.NodeInterfacesWirelessVifSVifCDhcpvsixOptions.IsUnknown()) {
		var subModel InterfacesWirelessVifSVifCDhcpvsixOptions
		diags.Append(o.NodeInterfacesWirelessVifSVifCDhcpvsixOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcpv6-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWirelessVifSVifCIP.IsNull() || o.NodeInterfacesWirelessVifSVifCIP.IsUnknown()) {
		var subModel InterfacesWirelessVifSVifCIP
		diags.Append(o.NodeInterfacesWirelessVifSVifCIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWirelessVifSVifCIPvsix.IsNull() || o.NodeInterfacesWirelessVifSVifCIPvsix.IsUnknown()) {
		var subModel InterfacesWirelessVifSVifCIPvsix
		diags.Append(o.NodeInterfacesWirelessVifSVifCIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWirelessVifSVifCMirror.IsNull() || o.NodeInterfacesWirelessVifSVifCMirror.IsUnknown()) {
		var subModel InterfacesWirelessVifSVifCMirror
		diags.Append(o.NodeInterfacesWirelessVifSVifCMirror.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["mirror"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWirelessVifSVifC) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafInterfacesWirelessVifSVifCDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesWirelessVifSVifCAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-link-detect"]; ok {
		o.LeafInterfacesWirelessVifSVifCDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDisableLinkDetect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesWirelessVifSVifCDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mac"]; ok {
		o.LeafInterfacesWirelessVifSVifCMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCMac = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mtu"]; ok {
		o.LeafInterfacesWirelessVifSVifCMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCMtu = basetypes.NewStringNull()
	}
	if value, ok := vyosData["redirect"]; ok {
		o.LeafInterfacesWirelessVifSVifCRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCRedirect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafInterfacesWirelessVifSVifCVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWirelessVifSVifCVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["dhcp-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessVifSVifCDhcpOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessVifSVifCDhcpOptions = data

	} else {
		o.NodeInterfacesWirelessVifSVifCDhcpOptions = basetypes.NewObjectNull(InterfacesWirelessVifSVifCDhcpOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["dhcpv6-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessVifSVifCDhcpvsixOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessVifSVifCDhcpvsixOptions = data

	} else {
		o.NodeInterfacesWirelessVifSVifCDhcpvsixOptions = basetypes.NewObjectNull(InterfacesWirelessVifSVifCDhcpvsixOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessVifSVifCIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessVifSVifCIP = data

	} else {
		o.NodeInterfacesWirelessVifSVifCIP = basetypes.NewObjectNull(InterfacesWirelessVifSVifCIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessVifSVifCIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessVifSVifCIPvsix = data

	} else {
		o.NodeInterfacesWirelessVifSVifCIPvsix = basetypes.NewObjectNull(InterfacesWirelessVifSVifCIPvsix{}.AttributeTypes())
	}
	if value, ok := vyosData["mirror"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWirelessVifSVifCMirror{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWirelessVifSVifCMirror = data

	} else {
		o.NodeInterfacesWirelessVifSVifCMirror = basetypes.NewObjectNull(InterfacesWirelessVifSVifCMirror{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wireless", "vif-s", "vif-c"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWirelessVifSVifC) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description":         types.StringType,
		"address":             types.StringType,
		"disable_link_detect": types.StringType,
		"disable":             types.StringType,
		"mac":                 types.StringType,
		"mtu":                 types.StringType,
		"redirect":            types.StringType,
		"vrf":                 types.StringType,

		// Tags

		// Nodes
		"dhcp_options":   types.ObjectType{AttrTypes: InterfacesWirelessVifSVifCDhcpOptions{}.AttributeTypes()},
		"dhcpv6_options": types.ObjectType{AttrTypes: InterfacesWirelessVifSVifCDhcpvsixOptions{}.AttributeTypes()},
		"ip":             types.ObjectType{AttrTypes: InterfacesWirelessVifSVifCIP{}.AttributeTypes()},
		"ipv6":           types.ObjectType{AttrTypes: InterfacesWirelessVifSVifCIPvsix{}.AttributeTypes()},
		"mirror":         types.ObjectType{AttrTypes: InterfacesWirelessVifSVifCMirror{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSVifC) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
|  dhcp  |  Dynamic Host Configuration Protocol  |
|  dhcpv6  |  Dynamic Host Configuration Protocol for IPv6  |

`,
		},

		"disable_link_detect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		// TagNodes

		// Nodes

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessVifSVifCMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}