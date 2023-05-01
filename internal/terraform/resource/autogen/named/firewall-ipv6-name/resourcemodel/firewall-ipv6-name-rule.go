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

// FirewallIPvsixNameRule describes the resource data model.
type FirewallIPvsixNameRule struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleAction              types.String `tfsdk:"action"`
	LeafFirewallIPvsixNameRuleDescrIPtion         types.String `tfsdk:"description"`
	LeafFirewallIPvsixNameRuleDisable             types.String `tfsdk:"disable"`
	LeafFirewallIPvsixNameRuleLog                 types.String `tfsdk:"log"`
	LeafFirewallIPvsixNameRuleLogLevel            types.String `tfsdk:"log_level"`
	LeafFirewallIPvsixNameRuleProtocol            types.String `tfsdk:"protocol"`
	LeafFirewallIPvsixNameRuleDscp                types.String `tfsdk:"dscp"`
	LeafFirewallIPvsixNameRuleDscpExclude         types.String `tfsdk:"dscp_exclude"`
	LeafFirewallIPvsixNameRulePacketLength        types.String `tfsdk:"packet_length"`
	LeafFirewallIPvsixNameRulePacketLengthExclude types.String `tfsdk:"packet_length_exclude"`
	LeafFirewallIPvsixNameRulePacketType          types.String `tfsdk:"packet_type"`
	LeafFirewallIPvsixNameRuleConnectionMark      types.String `tfsdk:"connection_mark"`
	LeafFirewallIPvsixNameRuleJumpTarget          types.String `tfsdk:"jump_target"`
	LeafFirewallIPvsixNameRuleQueue               types.String `tfsdk:"queue"`
	LeafFirewallIPvsixNameRuleQueueOptions        types.String `tfsdk:"queue_options"`

	// TagNodes

	// Nodes
	NodeFirewallIPvsixNameRuleDestination       types.Object `tfsdk:"destination"`
	NodeFirewallIPvsixNameRuleSource            types.Object `tfsdk:"source"`
	NodeFirewallIPvsixNameRuleFragment          types.Object `tfsdk:"fragment"`
	NodeFirewallIPvsixNameRuleInboundInterface  types.Object `tfsdk:"inbound_interface"`
	NodeFirewallIPvsixNameRuleOutboundInterface types.Object `tfsdk:"outbound_interface"`
	NodeFirewallIPvsixNameRuleIPsec             types.Object `tfsdk:"ipsec"`
	NodeFirewallIPvsixNameRuleLimit             types.Object `tfsdk:"limit"`
	NodeFirewallIPvsixNameRuleConnectionStatus  types.Object `tfsdk:"connection_status"`
	NodeFirewallIPvsixNameRuleRecent            types.Object `tfsdk:"recent"`
	NodeFirewallIPvsixNameRuleState             types.Object `tfsdk:"state"`
	NodeFirewallIPvsixNameRuleTCP               types.Object `tfsdk:"tcp"`
	NodeFirewallIPvsixNameRuleTime              types.Object `tfsdk:"time"`
	NodeFirewallIPvsixNameRuleHopLimit          types.Object `tfsdk:"hop_limit"`
	NodeFirewallIPvsixNameRuleIcmpvsix          types.Object `tfsdk:"icmpv6"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *FirewallIPvsixNameRule) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafFirewallIPvsixNameRuleAction.IsNull() || o.LeafFirewallIPvsixNameRuleAction.IsUnknown()) {
		vyosData["action"] = o.LeafFirewallIPvsixNameRuleAction.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleDescrIPtion.IsNull() || o.LeafFirewallIPvsixNameRuleDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafFirewallIPvsixNameRuleDescrIPtion.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleDisable.IsNull() || o.LeafFirewallIPvsixNameRuleDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafFirewallIPvsixNameRuleDisable.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleLog.IsNull() || o.LeafFirewallIPvsixNameRuleLog.IsUnknown()) {
		vyosData["log"] = o.LeafFirewallIPvsixNameRuleLog.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleLogLevel.IsNull() || o.LeafFirewallIPvsixNameRuleLogLevel.IsUnknown()) {
		vyosData["log-level"] = o.LeafFirewallIPvsixNameRuleLogLevel.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleProtocol.IsNull() || o.LeafFirewallIPvsixNameRuleProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafFirewallIPvsixNameRuleProtocol.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleDscp.IsNull() || o.LeafFirewallIPvsixNameRuleDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafFirewallIPvsixNameRuleDscp.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleDscpExclude.IsNull() || o.LeafFirewallIPvsixNameRuleDscpExclude.IsUnknown()) {
		vyosData["dscp-exclude"] = o.LeafFirewallIPvsixNameRuleDscpExclude.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRulePacketLength.IsNull() || o.LeafFirewallIPvsixNameRulePacketLength.IsUnknown()) {
		vyosData["packet-length"] = o.LeafFirewallIPvsixNameRulePacketLength.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRulePacketLengthExclude.IsNull() || o.LeafFirewallIPvsixNameRulePacketLengthExclude.IsUnknown()) {
		vyosData["packet-length-exclude"] = o.LeafFirewallIPvsixNameRulePacketLengthExclude.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRulePacketType.IsNull() || o.LeafFirewallIPvsixNameRulePacketType.IsUnknown()) {
		vyosData["packet-type"] = o.LeafFirewallIPvsixNameRulePacketType.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleConnectionMark.IsNull() || o.LeafFirewallIPvsixNameRuleConnectionMark.IsUnknown()) {
		vyosData["connection-mark"] = o.LeafFirewallIPvsixNameRuleConnectionMark.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleJumpTarget.IsNull() || o.LeafFirewallIPvsixNameRuleJumpTarget.IsUnknown()) {
		vyosData["jump-target"] = o.LeafFirewallIPvsixNameRuleJumpTarget.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleQueue.IsNull() || o.LeafFirewallIPvsixNameRuleQueue.IsUnknown()) {
		vyosData["queue"] = o.LeafFirewallIPvsixNameRuleQueue.ValueString()
	}
	if !(o.LeafFirewallIPvsixNameRuleQueueOptions.IsNull() || o.LeafFirewallIPvsixNameRuleQueueOptions.IsUnknown()) {
		vyosData["queue-options"] = o.LeafFirewallIPvsixNameRuleQueueOptions.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeFirewallIPvsixNameRuleDestination.IsNull() || o.NodeFirewallIPvsixNameRuleDestination.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleDestination
		diags.Append(o.NodeFirewallIPvsixNameRuleDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleSource.IsNull() || o.NodeFirewallIPvsixNameRuleSource.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleSource
		diags.Append(o.NodeFirewallIPvsixNameRuleSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleFragment.IsNull() || o.NodeFirewallIPvsixNameRuleFragment.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleFragment
		diags.Append(o.NodeFirewallIPvsixNameRuleFragment.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["fragment"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleInboundInterface.IsNull() || o.NodeFirewallIPvsixNameRuleInboundInterface.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleInboundInterface
		diags.Append(o.NodeFirewallIPvsixNameRuleInboundInterface.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["inbound-interface"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleOutboundInterface.IsNull() || o.NodeFirewallIPvsixNameRuleOutboundInterface.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleOutboundInterface
		diags.Append(o.NodeFirewallIPvsixNameRuleOutboundInterface.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["outbound-interface"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleIPsec.IsNull() || o.NodeFirewallIPvsixNameRuleIPsec.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleIPsec
		diags.Append(o.NodeFirewallIPvsixNameRuleIPsec.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipsec"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleLimit.IsNull() || o.NodeFirewallIPvsixNameRuleLimit.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleLimit
		diags.Append(o.NodeFirewallIPvsixNameRuleLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["limit"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleConnectionStatus.IsNull() || o.NodeFirewallIPvsixNameRuleConnectionStatus.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleConnectionStatus
		diags.Append(o.NodeFirewallIPvsixNameRuleConnectionStatus.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["connection-status"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleRecent.IsNull() || o.NodeFirewallIPvsixNameRuleRecent.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleRecent
		diags.Append(o.NodeFirewallIPvsixNameRuleRecent.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["recent"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleState.IsNull() || o.NodeFirewallIPvsixNameRuleState.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleState
		diags.Append(o.NodeFirewallIPvsixNameRuleState.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["state"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleTCP.IsNull() || o.NodeFirewallIPvsixNameRuleTCP.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleTCP
		diags.Append(o.NodeFirewallIPvsixNameRuleTCP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["tcp"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleTime.IsNull() || o.NodeFirewallIPvsixNameRuleTime.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleTime
		diags.Append(o.NodeFirewallIPvsixNameRuleTime.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["time"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleHopLimit.IsNull() || o.NodeFirewallIPvsixNameRuleHopLimit.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleHopLimit
		diags.Append(o.NodeFirewallIPvsixNameRuleHopLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["hop-limit"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeFirewallIPvsixNameRuleIcmpvsix.IsNull() || o.NodeFirewallIPvsixNameRuleIcmpvsix.IsUnknown()) {
		var subModel FirewallIPvsixNameRuleIcmpvsix
		diags.Append(o.NodeFirewallIPvsixNameRuleIcmpvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["icmpv6"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *FirewallIPvsixNameRule) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule"}})

	// Leafs
	if value, ok := vyosData["action"]; ok {
		o.LeafFirewallIPvsixNameRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleAction = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafFirewallIPvsixNameRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafFirewallIPvsixNameRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["log"]; ok {
		o.LeafFirewallIPvsixNameRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleLog = basetypes.NewStringNull()
	}
	if value, ok := vyosData["log-level"]; ok {
		o.LeafFirewallIPvsixNameRuleLogLevel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleLogLevel = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafFirewallIPvsixNameRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleProtocol = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp"]; ok {
		o.LeafFirewallIPvsixNameRuleDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp-exclude"]; ok {
		o.LeafFirewallIPvsixNameRuleDscpExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleDscpExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-length"]; ok {
		o.LeafFirewallIPvsixNameRulePacketLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRulePacketLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-length-exclude"]; ok {
		o.LeafFirewallIPvsixNameRulePacketLengthExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRulePacketLengthExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-type"]; ok {
		o.LeafFirewallIPvsixNameRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRulePacketType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["connection-mark"]; ok {
		o.LeafFirewallIPvsixNameRuleConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleConnectionMark = basetypes.NewStringNull()
	}
	if value, ok := vyosData["jump-target"]; ok {
		o.LeafFirewallIPvsixNameRuleJumpTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleJumpTarget = basetypes.NewStringNull()
	}
	if value, ok := vyosData["queue"]; ok {
		o.LeafFirewallIPvsixNameRuleQueue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleQueue = basetypes.NewStringNull()
	}
	if value, ok := vyosData["queue-options"]; ok {
		o.LeafFirewallIPvsixNameRuleQueueOptions = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallIPvsixNameRuleQueueOptions = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleDestination = data

	} else {
		o.NodeFirewallIPvsixNameRuleDestination = basetypes.NewObjectNull(FirewallIPvsixNameRuleDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleSource = data

	} else {
		o.NodeFirewallIPvsixNameRuleSource = basetypes.NewObjectNull(FirewallIPvsixNameRuleSource{}.AttributeTypes())
	}
	if value, ok := vyosData["fragment"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleFragment{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleFragment = data

	} else {
		o.NodeFirewallIPvsixNameRuleFragment = basetypes.NewObjectNull(FirewallIPvsixNameRuleFragment{}.AttributeTypes())
	}
	if value, ok := vyosData["inbound-interface"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleInboundInterface{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleInboundInterface = data

	} else {
		o.NodeFirewallIPvsixNameRuleInboundInterface = basetypes.NewObjectNull(FirewallIPvsixNameRuleInboundInterface{}.AttributeTypes())
	}
	if value, ok := vyosData["outbound-interface"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleOutboundInterface{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleOutboundInterface = data

	} else {
		o.NodeFirewallIPvsixNameRuleOutboundInterface = basetypes.NewObjectNull(FirewallIPvsixNameRuleOutboundInterface{}.AttributeTypes())
	}
	if value, ok := vyosData["ipsec"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleIPsec{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleIPsec = data

	} else {
		o.NodeFirewallIPvsixNameRuleIPsec = basetypes.NewObjectNull(FirewallIPvsixNameRuleIPsec{}.AttributeTypes())
	}
	if value, ok := vyosData["limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleLimit = data

	} else {
		o.NodeFirewallIPvsixNameRuleLimit = basetypes.NewObjectNull(FirewallIPvsixNameRuleLimit{}.AttributeTypes())
	}
	if value, ok := vyosData["connection-status"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleConnectionStatus{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleConnectionStatus = data

	} else {
		o.NodeFirewallIPvsixNameRuleConnectionStatus = basetypes.NewObjectNull(FirewallIPvsixNameRuleConnectionStatus{}.AttributeTypes())
	}
	if value, ok := vyosData["recent"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleRecent{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleRecent = data

	} else {
		o.NodeFirewallIPvsixNameRuleRecent = basetypes.NewObjectNull(FirewallIPvsixNameRuleRecent{}.AttributeTypes())
	}
	if value, ok := vyosData["state"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleState{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleState = data

	} else {
		o.NodeFirewallIPvsixNameRuleState = basetypes.NewObjectNull(FirewallIPvsixNameRuleState{}.AttributeTypes())
	}
	if value, ok := vyosData["tcp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleTCP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleTCP = data

	} else {
		o.NodeFirewallIPvsixNameRuleTCP = basetypes.NewObjectNull(FirewallIPvsixNameRuleTCP{}.AttributeTypes())
	}
	if value, ok := vyosData["time"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleTime{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleTime = data

	} else {
		o.NodeFirewallIPvsixNameRuleTime = basetypes.NewObjectNull(FirewallIPvsixNameRuleTime{}.AttributeTypes())
	}
	if value, ok := vyosData["hop-limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleHopLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleHopLimit = data

	} else {
		o.NodeFirewallIPvsixNameRuleHopLimit = basetypes.NewObjectNull(FirewallIPvsixNameRuleHopLimit{}.AttributeTypes())
	}
	if value, ok := vyosData["icmpv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, FirewallIPvsixNameRuleIcmpvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeFirewallIPvsixNameRuleIcmpvsix = data

	} else {
		o.NodeFirewallIPvsixNameRuleIcmpvsix = basetypes.NewObjectNull(FirewallIPvsixNameRuleIcmpvsix{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"firewall", "ipv6-name", "rule"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o FirewallIPvsixNameRule) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"action":                types.StringType,
		"description":           types.StringType,
		"disable":               types.StringType,
		"log":                   types.StringType,
		"log_level":             types.StringType,
		"protocol":              types.StringType,
		"dscp":                  types.StringType,
		"dscp_exclude":          types.StringType,
		"packet_length":         types.StringType,
		"packet_length_exclude": types.StringType,
		"packet_type":           types.StringType,
		"connection_mark":       types.StringType,
		"jump_target":           types.StringType,
		"queue":                 types.StringType,
		"queue_options":         types.StringType,

		// Tags

		// Nodes
		"destination":        types.ObjectType{AttrTypes: FirewallIPvsixNameRuleDestination{}.AttributeTypes()},
		"source":             types.ObjectType{AttrTypes: FirewallIPvsixNameRuleSource{}.AttributeTypes()},
		"fragment":           types.ObjectType{AttrTypes: FirewallIPvsixNameRuleFragment{}.AttributeTypes()},
		"inbound_interface":  types.ObjectType{AttrTypes: FirewallIPvsixNameRuleInboundInterface{}.AttributeTypes()},
		"outbound_interface": types.ObjectType{AttrTypes: FirewallIPvsixNameRuleOutboundInterface{}.AttributeTypes()},
		"ipsec":              types.ObjectType{AttrTypes: FirewallIPvsixNameRuleIPsec{}.AttributeTypes()},
		"limit":              types.ObjectType{AttrTypes: FirewallIPvsixNameRuleLimit{}.AttributeTypes()},
		"connection_status":  types.ObjectType{AttrTypes: FirewallIPvsixNameRuleConnectionStatus{}.AttributeTypes()},
		"recent":             types.ObjectType{AttrTypes: FirewallIPvsixNameRuleRecent{}.AttributeTypes()},
		"state":              types.ObjectType{AttrTypes: FirewallIPvsixNameRuleState{}.AttributeTypes()},
		"tcp":                types.ObjectType{AttrTypes: FirewallIPvsixNameRuleTCP{}.AttributeTypes()},
		"time":               types.ObjectType{AttrTypes: FirewallIPvsixNameRuleTime{}.AttributeTypes()},
		"hop_limit":          types.ObjectType{AttrTypes: FirewallIPvsixNameRuleHopLimit{}.AttributeTypes()},
		"icmpv6":             types.ObjectType{AttrTypes: FirewallIPvsixNameRuleIcmpvsix{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  jump  |  Jump to another chain  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the                       last chain  |
|  drop  |  Drop matching entries  |
|  queue  |  Enqueue packet to userspace  |

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
			MarkdownDescription: `Option to disable firewall rule

`,
		},

		"log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Option to log packets matching rule

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable log  |
|  disable  |  Disable log  |

`,
		},

		"log_level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level. Log must be enable.

|  Format  |  Description  |
|----------|---------------|
|  emerg  |  Emerg log level  |
|  alert  |  Alert log level  |
|  crit  |  Critical log level  |
|  err  |  Error log level  |
|  warn  |  Warning log level  |
|  notice  |  Notice log level  |
|  info  |  Info log level  |
|  debug  |  Debug log level  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  tcp_udp  |  Both TCP and UDP  |
|  u32:0-255  |  IP protocol number  |
|  <protocol>  |  IP protocol name  |
|  !<protocol>  |  IP protocol name  |

`,
		},

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DSCP value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value to match  |
|  <start-end>  |  DSCP range to match  |

`,
		},

		"dscp_exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DSCP value not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value not to match  |
|  <start-end>  |  DSCP range not to match  |

`,
		},

		"packet_length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length to match  |
|  <start-end>  |  Packet length range to match  |

`,
		},

		"packet_length_exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length not to match  |
|  <start-end>  |  Packet length range not to match  |

`,
		},

		"packet_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Match broadcast packet type  |
|  host  |  Match host packet type, addressed to local host  |
|  multicast  |  Match multicast packet type  |
|  other  |  Match packet addressed to another host  |

`,
		},

		"connection_mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Connection mark

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection-mark to match  |

`,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"queue": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Queue target  |

`,
		},

		"queue_options": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Options used for queue target. Action queue must be defined to use this
                    setting

|  Format  |  Description  |
|----------|---------------|
|  bypass  |  Let packets go through if userspace application cannot back off  |
|  fanout  |  Distribute packets between several queues  |

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleOutboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleConnectionStatus{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleState{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleHopLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleIcmpvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
		},
	}
}
