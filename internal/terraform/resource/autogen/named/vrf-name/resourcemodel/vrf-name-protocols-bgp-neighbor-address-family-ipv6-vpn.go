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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll         types.String `tfsdk:"addpath_tx_all"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs       types.String `tfsdk:"addpath_tx_per_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe           types.String `tfsdk:"as_override"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix        types.String `tfsdk:"maximum_prefix"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut     types.String `tfsdk:"maximum_prefix_out"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs      types.String `tfsdk:"remove_private_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient types.String `tfsdk:"route_reflector_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient    types.String `tfsdk:"route_server_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap        types.String `tfsdk:"unsuppress_map"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight               types.String `tfsdk:"weight"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal           types.Object `tfsdk:"nexthop_local"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList             types.Object `tfsdk:"prefix_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn              types.Object `tfsdk:"allowas_in"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity   types.Object `tfsdk:"disable_send_community"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList         types.Object `tfsdk:"distribute_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList             types.Object `tfsdk:"filter_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf            types.Object `tfsdk:"nexthop_self"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap               types.Object `tfsdk:"route_map"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-vpn"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll.IsUnknown()) {
		vyosData["addpath-tx-all"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs.IsUnknown()) {
		vyosData["addpath-tx-per-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe.IsUnknown()) {
		vyosData["as-override"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix.IsUnknown()) {
		vyosData["maximum-prefix"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut.IsUnknown()) {
		vyosData["maximum-prefix-out"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs.IsUnknown()) {
		vyosData["remove-private-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient.IsUnknown()) {
		vyosData["route-reflector-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient.IsUnknown()) {
		vyosData["route-server-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap.IsUnknown()) {
		vyosData["unsuppress-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight.IsUnknown()) {
		vyosData["weight"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["nexthop-local"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["conditionally-advertise"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["allowas-in"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["attribute-unchanged"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["disable-send-community"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["distribute-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["filter-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["nexthop-self"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["route-map"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["soft-reconfiguration"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-vpn"}})

	// Leafs
	if value, ok := vyosData["addpath-tx-all"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxAll = basetypes.NewStringNull()
	}
	if value, ok := vyosData["addpath-tx-per-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAddpathTxPerAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["as-override"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAsOverrIDe = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefix = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix-out"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnMaximumPrefixOut = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remove-private-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRemovePrivateAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-reflector-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteReflectorClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-server-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteServerClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["unsuppress-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnUnsuppressMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weight"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["nexthop-local"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal{}.AttributeTypes())
	}
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList{}.AttributeTypes())
	}
	if value, ok := vyosData["conditionally-advertise"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise{}.AttributeTypes())
	}
	if value, ok := vyosData["allowas-in"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}.AttributeTypes())
	}
	if value, ok := vyosData["attribute-unchanged"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}.AttributeTypes())
	}
	if value, ok := vyosData["disable-send-community"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity{}.AttributeTypes())
	}
	if value, ok := vyosData["distribute-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList{}.AttributeTypes())
	}
	if value, ok := vyosData["filter-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList{}.AttributeTypes())
	}
	if value, ok := vyosData["nexthop-self"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf{}.AttributeTypes())
	}
	if value, ok := vyosData["route-map"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap{}.AttributeTypes())
	}
	if value, ok := vyosData["soft-reconfiguration"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv6-vpn"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn) AttributeTypes() map[string]attr.Type {
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
		"nexthop_local":           types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal{}.AttributeTypes()},
		"prefix_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList{}.AttributeTypes()},
		"conditionally_advertise": types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise{}.AttributeTypes()},
		"allowas_in":              types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}.AttributeTypes()},
		"attribute_unchanged":     types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}.AttributeTypes()},
		"disable_send_community":  types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity{}.AttributeTypes()},
		"distribute_list":         types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList{}.AttributeTypes()},
		"filter_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList{}.AttributeTypes()},
		"nexthop_self":            types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf{}.AttributeTypes()},
		"route_map":               types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap{}.AttributeTypes()},
		"soft_reconfiguration":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpn) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopLocal{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvsixVpnSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}
