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

// VrfNameProtocolsBgpAddressFamily describes the resource data model.
type VrfNameProtocolsBgpAddressFamily struct {
	// LeafNodes

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast        types.Object `tfsdk:"ipv4_unicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast      types.Object `tfsdk:"ipv4_multicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast types.Object `tfsdk:"ipv4_labeled_unicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec       types.Object `tfsdk:"ipv4_flowspec"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn            types.Object `tfsdk:"ipv4_vpn"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast         types.Object `tfsdk:"ipv6_unicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast       types.Object `tfsdk:"ipv6_multicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast  types.Object `tfsdk:"ipv6_labeled_unicast"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec        types.Object `tfsdk:"ipv6_flowspec"`
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn             types.Object `tfsdk:"ipv6_vpn"`
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn           types.Object `tfsdk:"l2vpn_evpn"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpAddressFamily) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family"}})

	vyosData := make(map[string]interface{})

	// Leafs

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourUnicast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv4-unicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourMulticast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv4-multicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv4-labeled-unicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv4-flowspec"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvfourVpn
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv4-vpn"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixUnicast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6-unicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixMulticast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6-multicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6-labeled-unicast"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6-flowspec"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyIPvsixVpn
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipv6-vpn"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn.IsNull() || o.NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn
		diags.Append(o.NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["l2vpn-evpn"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpAddressFamily) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family"}})

	// Leafs

	// Tags

	// Nodes
	if value, ok := vyosData["ipv4-unicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv4-multicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourMulticast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv4-labeled-unicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv4-flowspec"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourFlowspec = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv4-vpn"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvfourVpn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvfourVpn = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvfourVpn{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6-unicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixUnicast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixUnicast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6-multicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6-labeled-unicast"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6-flowspec"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixFlowspec = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}.AttributeTypes())
	}
	if value, ok := vyosData["ipv6-vpn"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyIPvsixVpn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyIPvsixVpn = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyIPvsixVpn{}.AttributeTypes())
	}
	if value, ok := vyosData["l2vpn-evpn"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn = data

	} else {
		o.NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpn = basetypes.NewObjectNull(VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "address-family"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpAddressFamily) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs

		// Tags

		// Nodes
		"ipv4_unicast":         types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}.AttributeTypes()},
		"ipv4_multicast":       types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}.AttributeTypes()},
		"ipv4_labeled_unicast": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast{}.AttributeTypes()},
		"ipv4_flowspec":        types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec{}.AttributeTypes()},
		"ipv4_vpn":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvfourVpn{}.AttributeTypes()},
		"ipv6_unicast":         types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicast{}.AttributeTypes()},
		"ipv6_multicast":       types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}.AttributeTypes()},
		"ipv6_labeled_unicast": types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast{}.AttributeTypes()},
		"ipv6_flowspec":        types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}.AttributeTypes()},
		"ipv6_vpn":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyIPvsixVpn{}.AttributeTypes()},
		"l2vpn_evpn":           types.ObjectType{AttrTypes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamily) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"ipv4_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 BGP settings

`,
		},

		"ipv4_multicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourMulticast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Multicast IPv4 BGP settings

`,
		},

		"ipv4_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Labeled Unicast IPv4 BGP settings

`,
		},

		"ipv4_flowspec": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourFlowspec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Flowspec IPv4 BGP settings

`,
		},

		"ipv4_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Unicast VPN IPv4 BGP settings

`,
		},

		"ipv6_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 BGP settings

`,
		},

		"ipv6_multicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Multicast IPv6 BGP settings

`,
		},

		"ipv6_labeled_unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixLabeledUnicast{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Labeled Unicast IPv6 BGP settings

`,
		},

		"ipv6_flowspec": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixFlowspec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Flowspec IPv6 BGP settings

`,
		},

		"ipv6_vpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixVpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Unicast VPN IPv6 BGP settings

`,
		},

		"l2vpn_evpn": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `L2VPN EVPN BGP settings

`,
		},
	}
}
