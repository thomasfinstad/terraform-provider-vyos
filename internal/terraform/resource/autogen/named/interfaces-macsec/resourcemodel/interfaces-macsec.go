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

// InterfacesMacsec describes the resource data model.
type InterfacesMacsec struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesMacsecAddress         types.String `tfsdk:"address"`
	LeafInterfacesMacsecDescrIPtion     types.String `tfsdk:"description"`
	LeafInterfacesMacsecDisable         types.String `tfsdk:"disable"`
	LeafInterfacesMacsecMtu             types.String `tfsdk:"mtu"`
	LeafInterfacesMacsecSourceInterface types.String `tfsdk:"source_interface"`
	LeafInterfacesMacsecRedirect        types.String `tfsdk:"redirect"`
	LeafInterfacesMacsecVrf             types.String `tfsdk:"vrf"`

	// TagNodes

	// Nodes
	NodeInterfacesMacsecDhcpOptions     types.Object `tfsdk:"dhcp_options"`
	NodeInterfacesMacsecDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options"`
	NodeInterfacesMacsecIP              types.Object `tfsdk:"ip"`
	NodeInterfacesMacsecIPvsix          types.Object `tfsdk:"ipv6"`
	NodeInterfacesMacsecMirror          types.Object `tfsdk:"mirror"`
	NodeInterfacesMacsecSecURIty        types.Object `tfsdk:"security"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesMacsec) GetVyosPath() []string {
	return []string{
		"interfaces",
		"macsec",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesMacsec) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "macsec"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesMacsecAddress.IsNull() || o.LeafInterfacesMacsecAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesMacsecAddress.ValueString()
	}
	if !(o.LeafInterfacesMacsecDescrIPtion.IsNull() || o.LeafInterfacesMacsecDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafInterfacesMacsecDescrIPtion.ValueString()
	}
	if !(o.LeafInterfacesMacsecDisable.IsNull() || o.LeafInterfacesMacsecDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesMacsecDisable.ValueString()
	}
	if !(o.LeafInterfacesMacsecMtu.IsNull() || o.LeafInterfacesMacsecMtu.IsUnknown()) {
		vyosData["mtu"] = o.LeafInterfacesMacsecMtu.ValueString()
	}
	if !(o.LeafInterfacesMacsecSourceInterface.IsNull() || o.LeafInterfacesMacsecSourceInterface.IsUnknown()) {
		vyosData["source-interface"] = o.LeafInterfacesMacsecSourceInterface.ValueString()
	}
	if !(o.LeafInterfacesMacsecRedirect.IsNull() || o.LeafInterfacesMacsecRedirect.IsUnknown()) {
		vyosData["redirect"] = o.LeafInterfacesMacsecRedirect.ValueString()
	}
	if !(o.LeafInterfacesMacsecVrf.IsNull() || o.LeafInterfacesMacsecVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafInterfacesMacsecVrf.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeInterfacesMacsecDhcpOptions.IsNull() || o.NodeInterfacesMacsecDhcpOptions.IsUnknown()) {
		var subModel InterfacesMacsecDhcpOptions
		diags.Append(o.NodeInterfacesMacsecDhcpOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcp-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesMacsecDhcpvsixOptions.IsNull() || o.NodeInterfacesMacsecDhcpvsixOptions.IsUnknown()) {
		var subModel InterfacesMacsecDhcpvsixOptions
		diags.Append(o.NodeInterfacesMacsecDhcpvsixOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcpv6-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesMacsecIP.IsNull() || o.NodeInterfacesMacsecIP.IsUnknown()) {
		var subModel InterfacesMacsecIP
		diags.Append(o.NodeInterfacesMacsecIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesMacsecIPvsix.IsNull() || o.NodeInterfacesMacsecIPvsix.IsUnknown()) {
		var subModel InterfacesMacsecIPvsix
		diags.Append(o.NodeInterfacesMacsecIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesMacsecMirror.IsNull() || o.NodeInterfacesMacsecMirror.IsUnknown()) {
		var subModel InterfacesMacsecMirror
		diags.Append(o.NodeInterfacesMacsecMirror.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["mirror"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesMacsecSecURIty.IsNull() || o.NodeInterfacesMacsecSecURIty.IsUnknown()) {
		var subModel InterfacesMacsecSecURIty
		diags.Append(o.NodeInterfacesMacsecSecURIty.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["security"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesMacsec) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "macsec"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesMacsecAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafInterfacesMacsecDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesMacsecDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mtu"]; ok {
		o.LeafInterfacesMacsecMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecMtu = basetypes.NewStringNull()
	}
	if value, ok := vyosData["source-interface"]; ok {
		o.LeafInterfacesMacsecSourceInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecSourceInterface = basetypes.NewStringNull()
	}
	if value, ok := vyosData["redirect"]; ok {
		o.LeafInterfacesMacsecRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecRedirect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafInterfacesMacsecVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesMacsecVrf = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["dhcp-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecDhcpOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecDhcpOptions = data

	} else {
		o.NodeInterfacesMacsecDhcpOptions = basetypes.NewObjectNull(InterfacesMacsecDhcpOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["dhcpv6-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecDhcpvsixOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecDhcpvsixOptions = data

	} else {
		o.NodeInterfacesMacsecDhcpvsixOptions = basetypes.NewObjectNull(InterfacesMacsecDhcpvsixOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecIP = data

	} else {
		o.NodeInterfacesMacsecIP = basetypes.NewObjectNull(InterfacesMacsecIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecIPvsix = data

	} else {
		o.NodeInterfacesMacsecIPvsix = basetypes.NewObjectNull(InterfacesMacsecIPvsix{}.AttributeTypes())
	}
	if value, ok := vyosData["mirror"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecMirror{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecMirror = data

	} else {
		o.NodeInterfacesMacsecMirror = basetypes.NewObjectNull(InterfacesMacsecMirror{}.AttributeTypes())
	}
	if value, ok := vyosData["security"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesMacsecSecURIty{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesMacsecSecURIty = data

	} else {
		o.NodeInterfacesMacsecSecURIty = basetypes.NewObjectNull(InterfacesMacsecSecURIty{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "macsec"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesMacsec) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address":          types.StringType,
		"description":      types.StringType,
		"disable":          types.StringType,
		"mtu":              types.StringType,
		"source_interface": types.StringType,
		"redirect":         types.StringType,
		"vrf":              types.StringType,

		// Tags

		// Nodes
		"dhcp_options":   types.ObjectType{AttrTypes: InterfacesMacsecDhcpOptions{}.AttributeTypes()},
		"dhcpv6_options": types.ObjectType{AttrTypes: InterfacesMacsecDhcpvsixOptions{}.AttributeTypes()},
		"ip":             types.ObjectType{AttrTypes: InterfacesMacsecIP{}.AttributeTypes()},
		"ipv6":           types.ObjectType{AttrTypes: InterfacesMacsecIPvsix{}.AttributeTypes()},
		"mirror":         types.ObjectType{AttrTypes: InterfacesMacsecMirror{}.AttributeTypes()},
		"security":       types.ObjectType{AttrTypes: InterfacesMacsecSecURIty{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `MACsec Interface (802.1ae)

|  Format  |  Description  |
|----------|---------------|
|  macsecN  |  MACsec interface name  |

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

		"mtu": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1460`),
			Computed: true,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical interface the traffic will go through

|  Format  |  Description  |
|----------|---------------|
|  interface  |  Physical interface used for traffic forwarding  |

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
			Attributes: InterfacesMacsecDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"security": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Security/Encryption Settings

`,
		},
	}
}
