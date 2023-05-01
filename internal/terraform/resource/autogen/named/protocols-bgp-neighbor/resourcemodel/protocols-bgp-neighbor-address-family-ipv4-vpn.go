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

// ProtocolsBgpNeighborAddressFamilyIPvfourVpn describes the resource data model.
type ProtocolsBgpNeighborAddressFamilyIPvfourVpn struct {
	// LeafNodes
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll         types.String `tfsdk:"addpath_tx_all"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs       types.String `tfsdk:"addpath_tx_per_as"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe           types.String `tfsdk:"as_override"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix        types.String `tfsdk:"maximum_prefix"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut     types.String `tfsdk:"maximum_prefix_out"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs      types.String `tfsdk:"remove_private_as"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient types.String `tfsdk:"route_reflector_client"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient    types.String `tfsdk:"route_server_client"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap        types.String `tfsdk:"unsuppress_map"`
	LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight               types.String `tfsdk:"weight"`

	// TagNodes

	// Nodes
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList             types.Object `tfsdk:"prefix_list"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn              types.Object `tfsdk:"allowas_in"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity   types.Object `tfsdk:"disable_send_community"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList         types.Object `tfsdk:"distribute_list"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList             types.Object `tfsdk:"filter_list"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf            types.Object `tfsdk:"nexthop_self"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap               types.Object `tfsdk:"route_map"`
	NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.IsUnknown()) {
		vyosData["addpath-tx-all"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.IsUnknown()) {
		vyosData["addpath-tx-per-as"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.IsUnknown()) {
		vyosData["as-override"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.IsUnknown()) {
		vyosData["maximum-prefix"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.IsUnknown()) {
		vyosData["maximum-prefix-out"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.IsUnknown()) {
		vyosData["remove-private-as"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.IsUnknown()) {
		vyosData["route-reflector-client"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.IsUnknown()) {
		vyosData["route-server-client"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.IsUnknown()) {
		vyosData["unsuppress-map"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.ValueString()
	}
	if !(o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.IsNull() || o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.IsUnknown()) {
		vyosData["weight"] = o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["conditionally-advertise"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["allowas-in"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["attribute-unchanged"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["disable-send-community"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["distribute-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["filter-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["nexthop-self"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["route-map"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.IsNull() || o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.IsUnknown()) {
		var subModel ProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration
		diags.Append(o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["soft-reconfiguration"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *ProtocolsBgpNeighborAddressFamilyIPvfourVpn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})

	// Leafs
	if value, ok := vyosData["addpath-tx-all"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll = basetypes.NewStringNull()
	}
	if value, ok := vyosData["addpath-tx-per-as"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["as-override"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix-out"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remove-private-as"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-reflector-client"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-server-client"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["unsuppress-map"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weight"]; ok {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes())
	}
	if value, ok := vyosData["conditionally-advertise"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes())
	}
	if value, ok := vyosData["allowas-in"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes())
	}
	if value, ok := vyosData["attribute-unchanged"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes())
	}
	if value, ok := vyosData["disable-send-community"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes())
	}
	if value, ok := vyosData["distribute-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes())
	}
	if value, ok := vyosData["filter-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes())
	}
	if value, ok := vyosData["nexthop-self"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes())
	}
	if value, ok := vyosData["route-map"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes())
	}
	if value, ok := vyosData["soft-reconfiguration"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, ProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration = data

	} else {
		o.NodeProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration = basetypes.NewObjectNull(ProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpn) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"addpath_tx_all":         types.StringType,
		"addpath_tx_per_as":      types.StringType,
		"as_override":            types.StringType,
		"maximum_prefix":         types.StringType,
		"maximum_prefix_out":     types.StringType,
		"remove_private_as":      types.StringType,
		"route_reflector_client": types.StringType,
		"route_server_client":    types.StringType,
		"unsuppress_map":         types.StringType,
		"weight":                 types.StringType,

		// Tags

		// Nodes
		"prefix_list":             types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes()},
		"conditionally_advertise": types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes()},
		"allowas_in":              types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes()},
		"attribute_unchanged":     types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes()},
		"disable_send_community":  types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes()},
		"distribute_list":         types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes()},
		"filter_list":             types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes()},
		"nexthop_self":            types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes()},
		"route_map":               types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes()},
		"soft_reconfiguration":    types.ObjectType{AttrTypes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborAddressFamilyIPvfourVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
		},

		"addpath_tx_per_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
		},

		"as_override": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
		},

		"maximum_prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"maximum_prefix_out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Prefix limit  |

`,
		},

		"remove_private_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_reflector_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
		},

		"route_server_client": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
		},

		"unsuppress_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |

`,
		},

		"weight": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Default weight  |

`,
		},

		// TagNodes

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}
