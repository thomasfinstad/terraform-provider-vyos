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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll         types.String `tfsdk:"addpath_tx_all"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs       types.String `tfsdk:"addpath_tx_per_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe           types.String `tfsdk:"as_override"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix        types.String `tfsdk:"maximum_prefix"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut     types.String `tfsdk:"maximum_prefix_out"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs      types.String `tfsdk:"remove_private_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient types.String `tfsdk:"route_reflector_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient    types.String `tfsdk:"route_server_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap        types.String `tfsdk:"unsuppress_map"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight               types.String `tfsdk:"weight"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList             types.Object `tfsdk:"prefix_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn              types.Object `tfsdk:"allowas_in"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity   types.Object `tfsdk:"disable_send_community"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList         types.Object `tfsdk:"distribute_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList             types.Object `tfsdk:"filter_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf            types.Object `tfsdk:"nexthop_self"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap               types.Object `tfsdk:"route_map"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.IsUnknown()) {
		vyosData["addpath-tx-all"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.IsUnknown()) {
		vyosData["addpath-tx-per-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.IsUnknown()) {
		vyosData["as-override"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.IsUnknown()) {
		vyosData["maximum-prefix"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.IsUnknown()) {
		vyosData["maximum-prefix-out"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.IsUnknown()) {
		vyosData["remove-private-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.IsUnknown()) {
		vyosData["route-reflector-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.IsUnknown()) {
		vyosData["route-server-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.IsUnknown()) {
		vyosData["unsuppress-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.IsUnknown()) {
		vyosData["weight"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["conditionally-advertise"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["allowas-in"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["attribute-unchanged"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["disable-send-community"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["distribute-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["filter-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["nexthop-self"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["route-map"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["soft-reconfiguration"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})

	// Leafs
	if value, ok := vyosData["addpath-tx-all"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxAll = basetypes.NewStringNull()
	}
	if value, ok := vyosData["addpath-tx-per-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAddpathTxPerAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["as-override"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAsOverrIDe = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefix = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix-out"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnMaximumPrefixOut = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remove-private-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRemovePrivateAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-reflector-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteReflectorClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-server-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteServerClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["unsuppress-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnUnsuppressMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weight"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes())
	}
	if value, ok := vyosData["conditionally-advertise"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes())
	}
	if value, ok := vyosData["allowas-in"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes())
	}
	if value, ok := vyosData["attribute-unchanged"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes())
	}
	if value, ok := vyosData["disable-send-community"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes())
	}
	if value, ok := vyosData["distribute-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes())
	}
	if value, ok := vyosData["filter-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes())
	}
	if value, ok := vyosData["nexthop-self"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes())
	}
	if value, ok := vyosData["route-map"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes())
	}
	if value, ok := vyosData["soft-reconfiguration"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-vpn"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn) AttributeTypes() map[string]attr.Type {
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
		"prefix_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.AttributeTypes()},
		"conditionally_advertise": types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.AttributeTypes()},
		"allowas_in":              types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.AttributeTypes()},
		"attribute_unchanged":     types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.AttributeTypes()},
		"disable_send_community":  types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.AttributeTypes()},
		"distribute_list":         types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.AttributeTypes()},
		"filter_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.AttributeTypes()},
		"nexthop_self":            types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.AttributeTypes()},
		"route_map":               types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.AttributeTypes()},
		"soft_reconfiguration":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourVpnSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}
