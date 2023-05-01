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

// PolicyRoutesixRule describes the resource data model.
type PolicyRoutesixRule struct {
	// LeafNodes
	LeafPolicyRoutesixRuleAction              types.String `tfsdk:"action"`
	LeafPolicyRoutesixRuleDescrIPtion         types.String `tfsdk:"description"`
	LeafPolicyRoutesixRuleDisable             types.String `tfsdk:"disable"`
	LeafPolicyRoutesixRuleLog                 types.String `tfsdk:"log"`
	LeafPolicyRoutesixRuleProtocol            types.String `tfsdk:"protocol"`
	LeafPolicyRoutesixRuleDscp                types.String `tfsdk:"dscp"`
	LeafPolicyRoutesixRuleDscpExclude         types.String `tfsdk:"dscp_exclude"`
	LeafPolicyRoutesixRulePacketLength        types.String `tfsdk:"packet_length"`
	LeafPolicyRoutesixRulePacketLengthExclude types.String `tfsdk:"packet_length_exclude"`
	LeafPolicyRoutesixRulePacketType          types.String `tfsdk:"packet_type"`
	LeafPolicyRoutesixRuleConnectionMark      types.String `tfsdk:"connection_mark"`

	// TagNodes

	// Nodes
	NodePolicyRoutesixRuleDestination types.Object `tfsdk:"destination"`
	NodePolicyRoutesixRuleSource      types.Object `tfsdk:"source"`
	NodePolicyRoutesixRuleFragment    types.Object `tfsdk:"fragment"`
	NodePolicyRoutesixRuleIPsec       types.Object `tfsdk:"ipsec"`
	NodePolicyRoutesixRuleLimit       types.Object `tfsdk:"limit"`
	NodePolicyRoutesixRuleRecent      types.Object `tfsdk:"recent"`
	NodePolicyRoutesixRuleSet         types.Object `tfsdk:"set"`
	NodePolicyRoutesixRuleState       types.Object `tfsdk:"state"`
	NodePolicyRoutesixRuleTCP         types.Object `tfsdk:"tcp"`
	NodePolicyRoutesixRuleTime        types.Object `tfsdk:"time"`
	NodePolicyRoutesixRuleIcmpvsix    types.Object `tfsdk:"icmpv6"`
	NodePolicyRoutesixRuleHopLimit    types.Object `tfsdk:"hop_limit"`
}

// TerraformToVyos converts terraform data to vyos data
func (o *PolicyRoutesixRule) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"policy", "route6", "rule"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafPolicyRoutesixRuleAction.IsNull() || o.LeafPolicyRoutesixRuleAction.IsUnknown()) {
		vyosData["action"] = o.LeafPolicyRoutesixRuleAction.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleDescrIPtion.IsNull() || o.LeafPolicyRoutesixRuleDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafPolicyRoutesixRuleDescrIPtion.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleDisable.IsNull() || o.LeafPolicyRoutesixRuleDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafPolicyRoutesixRuleDisable.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleLog.IsNull() || o.LeafPolicyRoutesixRuleLog.IsUnknown()) {
		vyosData["log"] = o.LeafPolicyRoutesixRuleLog.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleProtocol.IsNull() || o.LeafPolicyRoutesixRuleProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafPolicyRoutesixRuleProtocol.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleDscp.IsNull() || o.LeafPolicyRoutesixRuleDscp.IsUnknown()) {
		vyosData["dscp"] = o.LeafPolicyRoutesixRuleDscp.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleDscpExclude.IsNull() || o.LeafPolicyRoutesixRuleDscpExclude.IsUnknown()) {
		vyosData["dscp-exclude"] = o.LeafPolicyRoutesixRuleDscpExclude.ValueString()
	}
	if !(o.LeafPolicyRoutesixRulePacketLength.IsNull() || o.LeafPolicyRoutesixRulePacketLength.IsUnknown()) {
		vyosData["packet-length"] = o.LeafPolicyRoutesixRulePacketLength.ValueString()
	}
	if !(o.LeafPolicyRoutesixRulePacketLengthExclude.IsNull() || o.LeafPolicyRoutesixRulePacketLengthExclude.IsUnknown()) {
		vyosData["packet-length-exclude"] = o.LeafPolicyRoutesixRulePacketLengthExclude.ValueString()
	}
	if !(o.LeafPolicyRoutesixRulePacketType.IsNull() || o.LeafPolicyRoutesixRulePacketType.IsUnknown()) {
		vyosData["packet-type"] = o.LeafPolicyRoutesixRulePacketType.ValueString()
	}
	if !(o.LeafPolicyRoutesixRuleConnectionMark.IsNull() || o.LeafPolicyRoutesixRuleConnectionMark.IsUnknown()) {
		vyosData["connection-mark"] = o.LeafPolicyRoutesixRuleConnectionMark.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodePolicyRoutesixRuleDestination.IsNull() || o.NodePolicyRoutesixRuleDestination.IsUnknown()) {
		var subModel PolicyRoutesixRuleDestination
		diags.Append(o.NodePolicyRoutesixRuleDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleSource.IsNull() || o.NodePolicyRoutesixRuleSource.IsUnknown()) {
		var subModel PolicyRoutesixRuleSource
		diags.Append(o.NodePolicyRoutesixRuleSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleFragment.IsNull() || o.NodePolicyRoutesixRuleFragment.IsUnknown()) {
		var subModel PolicyRoutesixRuleFragment
		diags.Append(o.NodePolicyRoutesixRuleFragment.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["fragment"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleIPsec.IsNull() || o.NodePolicyRoutesixRuleIPsec.IsUnknown()) {
		var subModel PolicyRoutesixRuleIPsec
		diags.Append(o.NodePolicyRoutesixRuleIPsec.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["ipsec"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleLimit.IsNull() || o.NodePolicyRoutesixRuleLimit.IsUnknown()) {
		var subModel PolicyRoutesixRuleLimit
		diags.Append(o.NodePolicyRoutesixRuleLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["limit"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleRecent.IsNull() || o.NodePolicyRoutesixRuleRecent.IsUnknown()) {
		var subModel PolicyRoutesixRuleRecent
		diags.Append(o.NodePolicyRoutesixRuleRecent.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["recent"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleSet.IsNull() || o.NodePolicyRoutesixRuleSet.IsUnknown()) {
		var subModel PolicyRoutesixRuleSet
		diags.Append(o.NodePolicyRoutesixRuleSet.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["set"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleState.IsNull() || o.NodePolicyRoutesixRuleState.IsUnknown()) {
		var subModel PolicyRoutesixRuleState
		diags.Append(o.NodePolicyRoutesixRuleState.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["state"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleTCP.IsNull() || o.NodePolicyRoutesixRuleTCP.IsUnknown()) {
		var subModel PolicyRoutesixRuleTCP
		diags.Append(o.NodePolicyRoutesixRuleTCP.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["tcp"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleTime.IsNull() || o.NodePolicyRoutesixRuleTime.IsUnknown()) {
		var subModel PolicyRoutesixRuleTime
		diags.Append(o.NodePolicyRoutesixRuleTime.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["time"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleIcmpvsix.IsNull() || o.NodePolicyRoutesixRuleIcmpvsix.IsUnknown()) {
		var subModel PolicyRoutesixRuleIcmpvsix
		diags.Append(o.NodePolicyRoutesixRuleIcmpvsix.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["icmpv6"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodePolicyRoutesixRuleHopLimit.IsNull() || o.NodePolicyRoutesixRuleHopLimit.IsUnknown()) {
		var subModel PolicyRoutesixRuleHopLimit
		diags.Append(o.NodePolicyRoutesixRuleHopLimit.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["hop-limit"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *PolicyRoutesixRule) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"policy", "route6", "rule"}})

	// Leafs
	if value, ok := vyosData["action"]; ok {
		o.LeafPolicyRoutesixRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleAction = basetypes.NewStringNull()
	}
	if value, ok := vyosData["description"]; ok {
		o.LeafPolicyRoutesixRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafPolicyRoutesixRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["log"]; ok {
		o.LeafPolicyRoutesixRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleLog = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafPolicyRoutesixRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleProtocol = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp"]; ok {
		o.LeafPolicyRoutesixRuleDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDscp = basetypes.NewStringNull()
	}
	if value, ok := vyosData["dscp-exclude"]; ok {
		o.LeafPolicyRoutesixRuleDscpExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleDscpExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-length"]; ok {
		o.LeafPolicyRoutesixRulePacketLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRulePacketLength = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-length-exclude"]; ok {
		o.LeafPolicyRoutesixRulePacketLengthExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRulePacketLengthExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-type"]; ok {
		o.LeafPolicyRoutesixRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRulePacketType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["connection-mark"]; ok {
		o.LeafPolicyRoutesixRuleConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRoutesixRuleConnectionMark = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleDestination = data

	} else {
		o.NodePolicyRoutesixRuleDestination = basetypes.NewObjectNull(PolicyRoutesixRuleDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleSource = data

	} else {
		o.NodePolicyRoutesixRuleSource = basetypes.NewObjectNull(PolicyRoutesixRuleSource{}.AttributeTypes())
	}
	if value, ok := vyosData["fragment"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleFragment{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleFragment = data

	} else {
		o.NodePolicyRoutesixRuleFragment = basetypes.NewObjectNull(PolicyRoutesixRuleFragment{}.AttributeTypes())
	}
	if value, ok := vyosData["ipsec"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleIPsec{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleIPsec = data

	} else {
		o.NodePolicyRoutesixRuleIPsec = basetypes.NewObjectNull(PolicyRoutesixRuleIPsec{}.AttributeTypes())
	}
	if value, ok := vyosData["limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleLimit = data

	} else {
		o.NodePolicyRoutesixRuleLimit = basetypes.NewObjectNull(PolicyRoutesixRuleLimit{}.AttributeTypes())
	}
	if value, ok := vyosData["recent"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleRecent{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleRecent = data

	} else {
		o.NodePolicyRoutesixRuleRecent = basetypes.NewObjectNull(PolicyRoutesixRuleRecent{}.AttributeTypes())
	}
	if value, ok := vyosData["set"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleSet{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleSet = data

	} else {
		o.NodePolicyRoutesixRuleSet = basetypes.NewObjectNull(PolicyRoutesixRuleSet{}.AttributeTypes())
	}
	if value, ok := vyosData["state"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleState{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleState = data

	} else {
		o.NodePolicyRoutesixRuleState = basetypes.NewObjectNull(PolicyRoutesixRuleState{}.AttributeTypes())
	}
	if value, ok := vyosData["tcp"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleTCP{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleTCP = data

	} else {
		o.NodePolicyRoutesixRuleTCP = basetypes.NewObjectNull(PolicyRoutesixRuleTCP{}.AttributeTypes())
	}
	if value, ok := vyosData["time"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleTime{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleTime = data

	} else {
		o.NodePolicyRoutesixRuleTime = basetypes.NewObjectNull(PolicyRoutesixRuleTime{}.AttributeTypes())
	}
	if value, ok := vyosData["icmpv6"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleIcmpvsix{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleIcmpvsix = data

	} else {
		o.NodePolicyRoutesixRuleIcmpvsix = basetypes.NewObjectNull(PolicyRoutesixRuleIcmpvsix{}.AttributeTypes())
	}
	if value, ok := vyosData["hop-limit"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, PolicyRoutesixRuleHopLimit{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodePolicyRoutesixRuleHopLimit = data

	} else {
		o.NodePolicyRoutesixRuleHopLimit = basetypes.NewObjectNull(PolicyRoutesixRuleHopLimit{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"policy", "route6", "rule"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o PolicyRoutesixRule) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"action":                types.StringType,
		"description":           types.StringType,
		"disable":               types.StringType,
		"log":                   types.StringType,
		"protocol":              types.StringType,
		"dscp":                  types.StringType,
		"dscp_exclude":          types.StringType,
		"packet_length":         types.StringType,
		"packet_length_exclude": types.StringType,
		"packet_type":           types.StringType,
		"connection_mark":       types.StringType,

		// Tags

		// Nodes
		"destination": types.ObjectType{AttrTypes: PolicyRoutesixRuleDestination{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: PolicyRoutesixRuleSource{}.AttributeTypes()},
		"fragment":    types.ObjectType{AttrTypes: PolicyRoutesixRuleFragment{}.AttributeTypes()},
		"ipsec":       types.ObjectType{AttrTypes: PolicyRoutesixRuleIPsec{}.AttributeTypes()},
		"limit":       types.ObjectType{AttrTypes: PolicyRoutesixRuleLimit{}.AttributeTypes()},
		"recent":      types.ObjectType{AttrTypes: PolicyRoutesixRuleRecent{}.AttributeTypes()},
		"set":         types.ObjectType{AttrTypes: PolicyRoutesixRuleSet{}.AttributeTypes()},
		"state":       types.ObjectType{AttrTypes: PolicyRoutesixRuleState{}.AttributeTypes()},
		"tcp":         types.ObjectType{AttrTypes: PolicyRoutesixRuleTCP{}.AttributeTypes()},
		"time":        types.ObjectType{AttrTypes: PolicyRoutesixRuleTime{}.AttributeTypes()},
		"icmpv6":      types.ObjectType{AttrTypes: PolicyRoutesixRuleIcmpvsix{}.AttributeTypes()},
		"hop_limit":   types.ObjectType{AttrTypes: PolicyRoutesixRuleHopLimit{}.AttributeTypes()},
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the last chain  |
|  drop  |  Drop matching entries  |

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

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  tcp_udp  |  Both TCP and UDP  |
|  0-255  |  IP protocol number  |
|  !<protocol>  |  IP protocol number  |

`,

			// Default:          stringdefault.StaticString(`all`),
			Computed: true,
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

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleState{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleIcmpvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: PolicyRoutesixRuleHopLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
		},
	}
}
