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

// InterfacesBonding describes the resource data model.
type InterfacesBonding struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafInterfacesBondingAddress           types.String `tfsdk:"address"`
	LeafInterfacesBondingDescrIPtion       types.String `tfsdk:"description"`
	LeafInterfacesBondingDisableLinkDetect types.String `tfsdk:"disable_link_detect"`
	LeafInterfacesBondingDisable           types.String `tfsdk:"disable"`
	LeafInterfacesBondingVrf               types.String `tfsdk:"vrf"`
	LeafInterfacesBondingHashPolicy        types.String `tfsdk:"hash_policy"`
	LeafInterfacesBondingMac               types.String `tfsdk:"mac"`
	LeafInterfacesBondingMiiMonInterval    types.String `tfsdk:"mii_mon_interval"`
	LeafInterfacesBondingMinLinks          types.String `tfsdk:"min_links"`
	LeafInterfacesBondingLacpRate          types.String `tfsdk:"lacp_rate"`
	LeafInterfacesBondingMode              types.String `tfsdk:"mode"`
	LeafInterfacesBondingMtu               types.String `tfsdk:"mtu"`
	LeafInterfacesBondingPrimary           types.String `tfsdk:"primary"`
	LeafInterfacesBondingRedirect          types.String `tfsdk:"redirect"`
	LeafInterfacesBondingXdp               types.String `tfsdk:"xdp"`

	// TagNodes
	TagInterfacesBondingVifS types.Map `tfsdk:"vif_s"`
	TagInterfacesBondingVif  types.Map `tfsdk:"vif"`

	// Nodes
	NodeInterfacesBondingArpMonitor      types.Object `tfsdk:"arp_monitor"`
	NodeInterfacesBondingDhcpOptions     types.Object `tfsdk:"dhcp_options"`
	NodeInterfacesBondingDhcpvsixOptions types.Object `tfsdk:"dhcpv6_options"`
	NodeInterfacesBondingMirror          types.Object `tfsdk:"mirror"`
	NodeInterfacesBondingIP              types.Object `tfsdk:"ip"`
	NodeInterfacesBondingIPvsix          types.Object `tfsdk:"ipv6"`
	NodeInterfacesBondingMember          types.Object `tfsdk:"member"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBonding) GetVyosPath() []string {
	return []string{
		"interfaces",
		"bonding",
		o.ID.ValueString(),
	}
}

// TerraformToVyos converts terraform data to vyos data
func (o *InterfacesBonding) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"interfaces", "bonding"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafInterfacesBondingAddress.IsNull() || o.LeafInterfacesBondingAddress.IsUnknown()) {
		vyosData["address"] = o.LeafInterfacesBondingAddress.ValueString()
	}
	if !(o.LeafInterfacesBondingDescrIPtion.IsNull() || o.LeafInterfacesBondingDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafInterfacesBondingDescrIPtion.ValueString()
	}
	if !(o.LeafInterfacesBondingDisableLinkDetect.IsNull() || o.LeafInterfacesBondingDisableLinkDetect.IsUnknown()) {
		vyosData["disable-link-detect"] = o.LeafInterfacesBondingDisableLinkDetect.ValueString()
	}
	if !(o.LeafInterfacesBondingDisable.IsNull() || o.LeafInterfacesBondingDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafInterfacesBondingDisable.ValueString()
	}
	if !(o.LeafInterfacesBondingVrf.IsNull() || o.LeafInterfacesBondingVrf.IsUnknown()) {
		vyosData["vrf"] = o.LeafInterfacesBondingVrf.ValueString()
	}
	if !(o.LeafInterfacesBondingHashPolicy.IsNull() || o.LeafInterfacesBondingHashPolicy.IsUnknown()) {
		vyosData["hash-policy"] = o.LeafInterfacesBondingHashPolicy.ValueString()
	}
	if !(o.LeafInterfacesBondingMac.IsNull() || o.LeafInterfacesBondingMac.IsUnknown()) {
		vyosData["mac"] = o.LeafInterfacesBondingMac.ValueString()
	}
	if !(o.LeafInterfacesBondingMiiMonInterval.IsNull() || o.LeafInterfacesBondingMiiMonInterval.IsUnknown()) {
		vyosData["mii-mon-interval"] = o.LeafInterfacesBondingMiiMonInterval.ValueString()
	}
	if !(o.LeafInterfacesBondingMinLinks.IsNull() || o.LeafInterfacesBondingMinLinks.IsUnknown()) {
		vyosData["min-links"] = o.LeafInterfacesBondingMinLinks.ValueString()
	}
	if !(o.LeafInterfacesBondingLacpRate.IsNull() || o.LeafInterfacesBondingLacpRate.IsUnknown()) {
		vyosData["lacp-rate"] = o.LeafInterfacesBondingLacpRate.ValueString()
	}
	if !(o.LeafInterfacesBondingMode.IsNull() || o.LeafInterfacesBondingMode.IsUnknown()) {
		vyosData["mode"] = o.LeafInterfacesBondingMode.ValueString()
	}
	if !(o.LeafInterfacesBondingMtu.IsNull() || o.LeafInterfacesBondingMtu.IsUnknown()) {
		vyosData["mtu"] = o.LeafInterfacesBondingMtu.ValueString()
	}
	if !(o.LeafInterfacesBondingPrimary.IsNull() || o.LeafInterfacesBondingPrimary.IsUnknown()) {
		vyosData["primary"] = o.LeafInterfacesBondingPrimary.ValueString()
	}
	if !(o.LeafInterfacesBondingRedirect.IsNull() || o.LeafInterfacesBondingRedirect.IsUnknown()) {
		vyosData["redirect"] = o.LeafInterfacesBondingRedirect.ValueString()
	}
	if !(o.LeafInterfacesBondingXdp.IsNull() || o.LeafInterfacesBondingXdp.IsUnknown()) {
		vyosData["xdp"] = o.LeafInterfacesBondingXdp.ValueString()
	}

	// Tags
	if !(o.TagInterfacesBondingVifS.IsNull() || o.TagInterfacesBondingVifS.IsUnknown()) {
		subModel := make(map[string]InterfacesBondingVifS)
		diags.Append(o.TagInterfacesBondingVifS.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["vif-s"] = subData
	}
	if !(o.TagInterfacesBondingVif.IsNull() || o.TagInterfacesBondingVif.IsUnknown()) {
		subModel := make(map[string]InterfacesBondingVif)
		diags.Append(o.TagInterfacesBondingVif.ElementsAs(ctx, &subModel, false)...)

		subData := make(map[string]interface{})
		for k, v := range subModel {
			subData[k] = v.TerraformToVyos(ctx, diags)
		}
		vyosData["vif"] = subData
	}

	// Nodes
	if !(o.NodeInterfacesBondingArpMonitor.IsNull() || o.NodeInterfacesBondingArpMonitor.IsUnknown()) {
		var subModel InterfacesBondingArpMonitor
		diags.Append(o.NodeInterfacesBondingArpMonitor.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["arp-monitor"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingDhcpOptions.IsNull() || o.NodeInterfacesBondingDhcpOptions.IsUnknown()) {
		var subModel InterfacesBondingDhcpOptions
		diags.Append(o.NodeInterfacesBondingDhcpOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcp-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingDhcpvsixOptions.IsNull() || o.NodeInterfacesBondingDhcpvsixOptions.IsUnknown()) {
		var subModel InterfacesBondingDhcpvsixOptions
		diags.Append(o.NodeInterfacesBondingDhcpvsixOptions.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["dhcpv6-options"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingMirror.IsNull() || o.NodeInterfacesBondingMirror.IsUnknown()) {
		var subModel InterfacesBondingMirror
		diags.Append(o.NodeInterfacesBondingMirror.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["mirror"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingIP.IsNull() || o.NodeInterfacesBondingIP.IsUnknown()) {
		var subModel InterfacesBondingIP
		diags.Append(o.NodeInterfacesBondingIP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ip"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingIPvsix.IsNull() || o.NodeInterfacesBondingIPvsix.IsUnknown()) {
		var subModel InterfacesBondingIPvsix
		diags.Append(o.NodeInterfacesBondingIPvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeInterfacesBondingMember.IsNull() || o.NodeInterfacesBondingMember.IsUnknown()) {
		var subModel InterfacesBondingMember
		diags.Append(o.NodeInterfacesBondingMember.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["member"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *InterfacesBonding) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"interfaces", "bonding"}})

	// Leafs
	if value, ok := vyosData["address"]; ok {
		o.LeafInterfacesBondingAddress = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingAddress = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafInterfacesBondingDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable-link-detect"]; ok {
		o.LeafInterfacesBondingDisableLinkDetect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDisableLinkDetect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafInterfacesBondingDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["vrf"]; ok {
		o.LeafInterfacesBondingVrf = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingVrf = basetypes.NewStringNull()
	}
	if value, ok := vyosData["hash-policy"]; ok {
		o.LeafInterfacesBondingHashPolicy = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingHashPolicy = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mac"]; ok {
		o.LeafInterfacesBondingMac = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingMac = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mii-mon-interval"]; ok {
		o.LeafInterfacesBondingMiiMonInterval = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingMiiMonInterval = basetypes.NewStringNull()
	}
	if value, ok := vyosData["min-links"]; ok {
		o.LeafInterfacesBondingMinLinks = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingMinLinks = basetypes.NewStringNull()
	}
	if value, ok := vyosData["lacp-rate"]; ok {
		o.LeafInterfacesBondingLacpRate = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingLacpRate = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mode"]; ok {
		o.LeafInterfacesBondingMode = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingMode = basetypes.NewStringNull()
	}
	if value, ok := vyosData["mtu"]; ok {
		o.LeafInterfacesBondingMtu = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingMtu = basetypes.NewStringNull()
	}
	if value, ok := vyosData["primary"]; ok {
		o.LeafInterfacesBondingPrimary = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingPrimary = basetypes.NewStringNull()
	}
	if value, ok := vyosData["redirect"]; ok {
		o.LeafInterfacesBondingRedirect = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingRedirect = basetypes.NewStringNull()
	}
	if value, ok := vyosData["xdp"]; ok {
		o.LeafInterfacesBondingXdp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafInterfacesBondingXdp = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := vyosData["vif-s"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesBondingVifS{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesBondingVifS = data
	} else {
		o.TagInterfacesBondingVifS = basetypes.NewMapNull(types.ObjectType{})
	}
	if value, ok := vyosData["vif"]; ok {
		data, d := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: InterfacesBondingVif{}.AttributeTypes()}, value.(map[string]interface{}))
		diags.Append(d...)
		o.TagInterfacesBondingVif = data
	} else {
		o.TagInterfacesBondingVif = basetypes.NewMapNull(types.ObjectType{})
	}

	// Nodes
	if value, ok := vyosData["arp-monitor"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingArpMonitor{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingArpMonitor = data

	} else {
		o.NodeInterfacesBondingArpMonitor = basetypes.NewObjectNull(InterfacesBondingArpMonitor{}.AttributeTypes())
	}
	if value, ok := vyosData["dhcp-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingDhcpOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingDhcpOptions = data

	} else {
		o.NodeInterfacesBondingDhcpOptions = basetypes.NewObjectNull(InterfacesBondingDhcpOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["dhcpv6-options"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingDhcpvsixOptions{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingDhcpvsixOptions = data

	} else {
		o.NodeInterfacesBondingDhcpvsixOptions = basetypes.NewObjectNull(InterfacesBondingDhcpvsixOptions{}.AttributeTypes())
	}
	if value, ok := vyosData["mirror"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingMirror{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingMirror = data

	} else {
		o.NodeInterfacesBondingMirror = basetypes.NewObjectNull(InterfacesBondingMirror{}.AttributeTypes())
	}
	if value, ok := vyosData["ip"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingIP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingIP = data

	} else {
		o.NodeInterfacesBondingIP = basetypes.NewObjectNull(InterfacesBondingIP{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingIPvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingIPvsix = data

	} else {
		o.NodeInterfacesBondingIPvsix = basetypes.NewObjectNull(InterfacesBondingIPvsix{}.AttributeTypes())
	}
	if value, ok := vyosData["member"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, InterfacesBondingMember{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeInterfacesBondingMember = data

	} else {
		o.NodeInterfacesBondingMember = basetypes.NewObjectNull(InterfacesBondingMember{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"interfaces", "bonding"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o InterfacesBonding) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"address":             types.StringType,
		"description":         types.StringType,
		"disable_link_detect": types.StringType,
		"disable":             types.StringType,
		"vrf":                 types.StringType,
		"hash_policy":         types.StringType,
		"mac":                 types.StringType,
		"mii_mon_interval":    types.StringType,
		"min_links":           types.StringType,
		"lacp_rate":           types.StringType,
		"mode":                types.StringType,
		"mtu":                 types.StringType,
		"primary":             types.StringType,
		"redirect":            types.StringType,
		"xdp":                 types.StringType,

		// Tags
		"vif_s": types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesBondingVifS{}.AttributeTypes()}},
		"vif":   types.MapType{ElemType: types.ObjectType{AttrTypes: InterfacesBondingVif{}.AttributeTypes()}},

		// Nodes
		"arp_monitor":    types.ObjectType{AttrTypes: InterfacesBondingArpMonitor{}.AttributeTypes()},
		"dhcp_options":   types.ObjectType{AttrTypes: InterfacesBondingDhcpOptions{}.AttributeTypes()},
		"dhcpv6_options": types.ObjectType{AttrTypes: InterfacesBondingDhcpvsixOptions{}.AttributeTypes()},
		"mirror":         types.ObjectType{AttrTypes: InterfacesBondingMirror{}.AttributeTypes()},
		"ip":             types.ObjectType{AttrTypes: InterfacesBondingIP{}.AttributeTypes()},
		"ipv6":           types.ObjectType{AttrTypes: InterfacesBondingIPvsix{}.AttributeTypes()},
		"member":         types.ObjectType{AttrTypes: InterfacesBondingMember{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBonding) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Bonding Interface/Link Aggregation

|  Format  |  Description  |
|----------|---------------|
|  bondN  |  Bonding interface name  |

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

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

`,
		},

		"hash_policy": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bonding transmit hash policy

|  Format  |  Description  |
|----------|---------------|
|  layer2  |  use MAC addresses to generate the hash  |
|  layer2+3  |  combine MAC address and IP address to make hash  |
|  layer3+4  |  combine IP address and port to make hash  |
|  encap2+3  |  combine encapsulated MAC address and IP address to make hash  |
|  encap3+4  |  combine encapsulated IP address and port to make hash  |

`,

			// Default:          stringdefault.StaticString(`layer2`),
			Computed: true,
		},

		"mac": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		"mii_mon_interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the MII link monitoring frequency in milliseconds

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable MII link monitoring  |
|  u32:50-1000  |  MII link monitoring frequency in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"min_links": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Minimum number of member interfaces required up before enabling bond

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16  |  Minimum number of member interfaces required up before enabling bond  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"lacp_rate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rate in which we will ask our link partner to transmit LACPDU packets

|  Format  |  Description  |
|----------|---------------|
|  slow  |  Request partner to transmit LACPDUs every 30 seconds  |
|  fast  |  Request partner to transmit LACPDUs every 1 second  |

`,

			// Default:          stringdefault.StaticString(`slow`),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bonding mode

|  Format  |  Description  |
|----------|---------------|
|  802.3ad  |  IEEE 802.3ad Dynamic link aggregation  |
|  active-backup  |  Fault tolerant: only one slave in the bond is active  |
|  broadcast  |  Fault tolerant: transmits everything on all slave interfaces  |
|  round-robin  |  Load balance: transmit packets in sequential order  |
|  transmit-load-balance  |  Load balance: adapts based on transmit load and speed  |
|  adaptive-load-balance  |  Load balance: adapts based on transmit and receive plus ARP  |
|  xor-hash  |  Distribute based on MAC address  |

`,

			// Default:          stringdefault.StaticString(`802.3ad`),
			Computed: true,
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

		"primary": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Primary device interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

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

		"xdp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable eXpress Data Path

`,
		},

		// TagNodes

		"vif_s": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesBondingVifS{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |

`,
		},

		"vif": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesBondingVif{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4094  |  Virtual Local Area Network (VLAN) ID  |

`,
		},

		// Nodes

		"arp_monitor": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingArpMonitor{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ARP link monitoring parameters

`,
		},

		"dhcp_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingDhcpOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCP client settings/options

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingDhcpvsixOptions{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"member": schema.SingleNestedAttribute{
			Attributes: InterfacesBondingMember{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Bridge member interfaces

`,
		},
	}
}