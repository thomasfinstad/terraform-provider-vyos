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

// InterfacesWwan describes the resource data model.
type InterfacesWwan struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesWwanAddress           types.String `tfsdk:"address"`
	LeafInterfacesWwanApn               types.String `tfsdk:"apn"`
	LeafInterfacesWwanDescrIPtion       types.String `tfsdk:"description"`
	LeafInterfacesWwanDisable           types.String `tfsdk:"disable"`
	LeafInterfacesWwanDisableLinkDetect types.String `tfsdk:"disable_link_detect"`
	LeafInterfacesWwanMtu               types.String `tfsdk:"mtu"`
	LeafInterfacesWwanConnectOnDemand   types.String `tfsdk:"connect_on_demand"`
	LeafInterfacesWwanRedirect          types.String `tfsdk:"redirect"`
	LeafInterfacesWwanVrf               types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
	NodeInterfacesWwanDhcpOptions     types.Object `tfsdk:"dhcp_options"`
	NodeInterfacesWwanDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options"`
	NodeInterfacesWwanAuthentication  types.Object `tfsdk:"authentication"`
	NodeInterfacesWwanMirror          types.Object `tfsdk:"mirror"`
	NodeInterfacesWwanIP              types.Object `tfsdk:"ip"`
	NodeInterfacesWwanIPvsix          types.Object `tfsdk:"ipv6"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWwan) GetVyosPath() []string {
	return []string{
		"interfaces",
		"wwan",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesWwan) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "wwan"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesWwanAddress.IsNull() || o.LeafInterfacesWwanAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesWwanAddress.ValueString()
	}
	if !(o.LeafInterfacesWwanApn.IsNull() || o.LeafInterfacesWwanApn.IsUnknown()) {
		vyosData["apn"] = o.LeafInterfacesWwanApn.ValueString()
	}
	if !(o.LeafInterfacesWwanDescrIPtion.IsNull() || o.LeafInterfacesWwanDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafInterfacesWwanDescrIPtion.ValueString()
	}
	if !(o.LeafInterfacesWwanDisable.IsNull() || o.LeafInterfacesWwanDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesWwanDisable.ValueString()
	}
	if !(o.LeafInterfacesWwanDisableLinkDetect.IsNull() || o.LeafInterfacesWwanDisableLinkDetect.IsUnknown()) {
		vyosData["disable-link-detect"] = o.LeafInterfacesWwanDisableLinkDetect.ValueString()
	}
	if !(o.LeafInterfacesWwanMtu.IsNull() || o.LeafInterfacesWwanMtu.IsUnknown()) {
		vyosData["mtu"] = o.LeafInterfacesWwanMtu.ValueString()
	}
	if !(o.LeafInterfacesWwanConnectOnDemand.IsNull() || o.LeafInterfacesWwanConnectOnDemand.IsUnknown()) {
		vyosData["connect-on-demand"] = o.LeafInterfacesWwanConnectOnDemand.ValueString()
	}
	if !(o.LeafInterfacesWwanRedirect.IsNull() || o.LeafInterfacesWwanRedirect.IsUnknown()) {
		vyosData["redirect"] = o.LeafInterfacesWwanRedirect.ValueString()
	}
	if !(o.LeafInterfacesWwanVrf.IsNull() || o.LeafInterfacesWwanVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafInterfacesWwanVrf.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeInterfacesWwanDhcpOptions.IsNull() || o.NodeInterfacesWwanDhcpOptions.IsUnknown()) {
		var subModel InterfacesWwanDhcpOptions
		diags.Append(o.NodeInterfacesWwanDhcpOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcp-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWwanDhcpvsixOptions.IsNull() || o.NodeInterfacesWwanDhcpvsixOptions.IsUnknown()) {
		var subModel InterfacesWwanDhcpvsixOptions
		diags.Append(o.NodeInterfacesWwanDhcpvsixOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcpv6-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWwanAuthentication.IsNull() || o.NodeInterfacesWwanAuthentication.IsUnknown()) {
		var subModel InterfacesWwanAuthentication
		diags.Append(o.NodeInterfacesWwanAuthentication.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["authentication"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWwanMirror.IsNull() || o.NodeInterfacesWwanMirror.IsUnknown()) {
		var subModel InterfacesWwanMirror
		diags.Append(o.NodeInterfacesWwanMirror.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["mirror"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWwanIP.IsNull() || o.NodeInterfacesWwanIP.IsUnknown()) {
		var subModel InterfacesWwanIP
		diags.Append(o.NodeInterfacesWwanIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesWwanIPvsix.IsNull() || o.NodeInterfacesWwanIPvsix.IsUnknown()) {
		var subModel InterfacesWwanIPvsix
		diags.Append(o.NodeInterfacesWwanIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesWwan) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "wwan"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesWwanAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["apn"]; ok {
		o.LeafInterfacesWwanApn = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanApn = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafInterfacesWwanDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesWwanDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-link-detect"]; ok {
		o.LeafInterfacesWwanDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanDisableLinkDetect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mtu"]; ok {
		o.LeafInterfacesWwanMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanMtu = basetypes.NewStringNull()
	}
	if value, ok := vyosData["connect-on-demand"]; ok {
		o.LeafInterfacesWwanConnectOnDemand = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanConnectOnDemand = basetypes.NewStringNull()
	}
	if value, ok := vyosData["redirect"]; ok {
		o.LeafInterfacesWwanRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanRedirect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafInterfacesWwanVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesWwanVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["dhcp-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanDhcpOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanDhcpOptions = data

	} else {
		o.NodeInterfacesWwanDhcpOptions = basetypes.NewObjectNull(InterfacesWwanDhcpOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["dhcpv6-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanDhcpvsixOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanDhcpvsixOptions = data

	} else {
		o.NodeInterfacesWwanDhcpvsixOptions = basetypes.NewObjectNull(InterfacesWwanDhcpvsixOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["authentication"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanAuthentication{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanAuthentication = data

	} else {
		o.NodeInterfacesWwanAuthentication = basetypes.NewObjectNull(InterfacesWwanAuthentication{}.AttributeTypes())
	}
	if value, ok := vyosData["mirror"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanMirror{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanMirror = data

	} else {
		o.NodeInterfacesWwanMirror = basetypes.NewObjectNull(InterfacesWwanMirror{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanIP = data

	} else {
		o.NodeInterfacesWwanIP = basetypes.NewObjectNull(InterfacesWwanIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesWwanIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesWwanIPvsix = data

	} else {
		o.NodeInterfacesWwanIPvsix = basetypes.NewObjectNull(InterfacesWwanIPvsix{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "wwan"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesWwan) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address":             types.StringType,
		"apn":                 types.StringType,
		"description":         types.StringType,
		"disable":             types.StringType,
		"disable_link_detect": types.StringType,
		"mtu":                 types.StringType,
		"connect_on_demand":   types.StringType,
		"redirect":            types.StringType,
		"vrf":                 types.StringType,

		// Tags

		// Nodes
		"dhcp_options":   types.ObjectType{AttrTypes: InterfacesWwanDhcpOptions{}.AttributeTypes()},
		"dhcpv6_options": types.ObjectType{AttrTypes: InterfacesWwanDhcpvsixOptions{}.AttributeTypes()},
		"authentication": types.ObjectType{AttrTypes: InterfacesWwanAuthentication{}.AttributeTypes()},
		"mirror":         types.ObjectType{AttrTypes: InterfacesWwanMirror{}.AttributeTypes()},
		"ip":             types.ObjectType{AttrTypes: InterfacesWwanIP{}.AttributeTypes()},
		"ipv6":           types.ObjectType{AttrTypes: InterfacesWwanIPvsix{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWwan) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Wireless Modem (WWAN) Interface

|  Format  |  Description  |
|----------|---------------|
|  wwanN  |  Wireless Wide Area Network interface name  |

`,
		},

		// LeafNodes

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

		"apn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access Point Name (APN)

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"disable": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"disable_link_detect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ignore link state changes

`,
		},

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-1500  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1430`),
			Computed: true,
		},

		"connect_on_demand": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

`,
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
			Attributes: InterfacesWwanDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesWwanIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},
	}
}