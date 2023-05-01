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

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll         types.String `tfsdk:"addpath_tx_all"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs       types.String `tfsdk:"addpath_tx_per_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe           types.String `tfsdk:"as_override"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix        types.String `tfsdk:"maximum_prefix"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut     types.String `tfsdk:"maximum_prefix_out"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs      types.String `tfsdk:"remove_private_as"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient types.String `tfsdk:"route_reflector_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient    types.String `tfsdk:"route_server_client"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight               types.String `tfsdk:"weight"`

	// TagNodes

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability             types.Object `tfsdk:"capability"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList             types.Object `tfsdk:"prefix_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise types.Object `tfsdk:"conditionally_advertise"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn              types.Object `tfsdk:"allowas_in"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged     types.Object `tfsdk:"attribute_unchanged"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity   types.Object `tfsdk:"disable_send_community"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList         types.Object `tfsdk:"distribute_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList             types.Object `tfsdk:"filter_list"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf            types.Object `tfsdk:"nexthop_self"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap               types.Object `tfsdk:"route_map"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration    types.Object `tfsdk:"soft_reconfiguration"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate       types.Object `tfsdk:"default_originate"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll.IsUnknown()) {
		vyosData["addpath-tx-all"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs.IsUnknown()) {
		vyosData["addpath-tx-per-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe.IsUnknown()) {
		vyosData["as-override"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix.IsUnknown()) {
		vyosData["maximum-prefix"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut.IsUnknown()) {
		vyosData["maximum-prefix-out"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs.IsUnknown()) {
		vyosData["remove-private-as"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient.IsUnknown()) {
		vyosData["route-reflector-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient.IsUnknown()) {
		vyosData["route-server-client"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap.IsUnknown()) {
		vyosData["unsuppress-map"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap.ValueString()
	}
	if !(o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight.IsNull() || o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight.IsUnknown()) {
		vyosData["weight"] = o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["capability"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["prefix-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["conditionally-advertise"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["allowas-in"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["attribute-unchanged"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["disable-send-community"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["distribute-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["filter-list"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["nexthop-self"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["route-map"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["soft-reconfiguration"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate.IsNull() || o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate.IsUnknown()) {
		var subModel VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate
		diags.Append(o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["default-originate"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast"}})

	// Leafs
	if value, ok := vyosData["addpath-tx-all"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxAll = basetypes.NewStringNull()
	}
	if value, ok := vyosData["addpath-tx-per-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAddpathTxPerAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["as-override"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAsOverrIDe = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefix = basetypes.NewStringNull()
	}
	if value, ok := vyosData["maximum-prefix-out"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastMaximumPrefixOut = basetypes.NewStringNull()
	}
	if value, ok := vyosData["remove-private-as"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRemovePrivateAs = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-reflector-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteReflectorClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["route-server-client"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteServerClient = basetypes.NewStringNull()
	}
	if value, ok := vyosData["unsuppress-map"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastUnsuppressMap = basetypes.NewStringNull()
	}
	if value, ok := vyosData["weight"]; ok {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastWeight = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["capability"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability{}.AttributeTypes())
	}
	if value, ok := vyosData["prefix-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList{}.AttributeTypes())
	}
	if value, ok := vyosData["conditionally-advertise"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise{}.AttributeTypes())
	}
	if value, ok := vyosData["allowas-in"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn{}.AttributeTypes())
	}
	if value, ok := vyosData["attribute-unchanged"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged{}.AttributeTypes())
	}
	if value, ok := vyosData["disable-send-community"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity{}.AttributeTypes())
	}
	if value, ok := vyosData["distribute-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList{}.AttributeTypes())
	}
	if value, ok := vyosData["filter-list"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList{}.AttributeTypes())
	}
	if value, ok := vyosData["nexthop-self"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf{}.AttributeTypes())
	}
	if value, ok := vyosData["route-map"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap{}.AttributeTypes())
	}
	if value, ok := vyosData["soft-reconfiguration"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration{}.AttributeTypes())
	}
	if value, ok := vyosData["default-originate"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate = data

	} else {
		o.NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate = basetypes.NewObjectNull(VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"vrf", "name", "protocols", "bgp", "neighbor", "address-family", "ipv4-unicast"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast) AttributeTypes() map[string]attr.Type {
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
		"capability":              types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability{}.AttributeTypes()},
		"prefix_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList{}.AttributeTypes()},
		"conditionally_advertise": types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise{}.AttributeTypes()},
		"allowas_in":              types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn{}.AttributeTypes()},
		"attribute_unchanged":     types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged{}.AttributeTypes()},
		"disable_send_community":  types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity{}.AttributeTypes()},
		"distribute_list":         types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList{}.AttributeTypes()},
		"filter_list":             types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList{}.AttributeTypes()},
		"nexthop_self":            types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf{}.AttributeTypes()},
		"route_map":               types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap{}.AttributeTypes()},
		"soft_reconfiguration":    types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration{}.AttributeTypes()},
		"default_originate":       types.ObjectType{AttrTypes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicast) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAllowasIn{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastAttributeUnchanged{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDisableSendCommunity{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDistributeList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastNexthopSelf{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourUnicastDefaultOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
		},
	}
}