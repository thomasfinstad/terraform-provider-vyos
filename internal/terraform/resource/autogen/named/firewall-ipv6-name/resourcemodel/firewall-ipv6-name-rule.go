// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallIPvsixNameRule describes the resource data model.
type FirewallIPvsixNameRule struct {
	// LeafNodes
	FirewallIPvsixNameRuleAction              customtypes.CustomStringValue `tfsdk:"action" json:"action,omitempty"`
	FirewallIPvsixNameRuleDescrIPtion         customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	FirewallIPvsixNameRuleDisable             customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	FirewallIPvsixNameRuleLog                 customtypes.CustomStringValue `tfsdk:"log" json:"log,omitempty"`
	FirewallIPvsixNameRuleLogLevel            customtypes.CustomStringValue `tfsdk:"log_level" json:"log-level,omitempty"`
	FirewallIPvsixNameRuleProtocol            customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	FirewallIPvsixNameRuleDscp                customtypes.CustomStringValue `tfsdk:"dscp" json:"dscp,omitempty"`
	FirewallIPvsixNameRuleDscpExclude         customtypes.CustomStringValue `tfsdk:"dscp_exclude" json:"dscp-exclude,omitempty"`
	FirewallIPvsixNameRulePacketLength        customtypes.CustomStringValue `tfsdk:"packet_length" json:"packet-length,omitempty"`
	FirewallIPvsixNameRulePacketLengthExclude customtypes.CustomStringValue `tfsdk:"packet_length_exclude" json:"packet-length-exclude,omitempty"`
	FirewallIPvsixNameRulePacketType          customtypes.CustomStringValue `tfsdk:"packet_type" json:"packet-type,omitempty"`
	FirewallIPvsixNameRuleConnectionMark      customtypes.CustomStringValue `tfsdk:"connection_mark" json:"connection-mark,omitempty"`
	FirewallIPvsixNameRuleJumpTarget          customtypes.CustomStringValue `tfsdk:"jump_target" json:"jump-target,omitempty"`
	FirewallIPvsixNameRuleQueue               customtypes.CustomStringValue `tfsdk:"queue" json:"queue,omitempty"`
	FirewallIPvsixNameRuleQueueOptions        customtypes.CustomStringValue `tfsdk:"queue_options" json:"queue-options,omitempty"`

	// TagNodes

	// Nodes
	FirewallIPvsixNameRuleDestination       types.Object `tfsdk:"destination" json:"destination,omitempty"`
	FirewallIPvsixNameRuleSource            types.Object `tfsdk:"source" json:"source,omitempty"`
	FirewallIPvsixNameRuleFragment          types.Object `tfsdk:"fragment" json:"fragment,omitempty"`
	FirewallIPvsixNameRuleInboundInterface  types.Object `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`
	FirewallIPvsixNameRuleOutboundInterface types.Object `tfsdk:"outbound_interface" json:"outbound-interface,omitempty"`
	FirewallIPvsixNameRuleIPsec             types.Object `tfsdk:"ipsec" json:"ipsec,omitempty"`
	FirewallIPvsixNameRuleLimit             types.Object `tfsdk:"limit" json:"limit,omitempty"`
	FirewallIPvsixNameRuleConnectionStatus  types.Object `tfsdk:"connection_status" json:"connection-status,omitempty"`
	FirewallIPvsixNameRuleRecent            types.Object `tfsdk:"recent" json:"recent,omitempty"`
	FirewallIPvsixNameRuleState             types.Object `tfsdk:"state" json:"state,omitempty"`
	FirewallIPvsixNameRuleTCP               types.Object `tfsdk:"tcp" json:"tcp,omitempty"`
	FirewallIPvsixNameRuleTime              types.Object `tfsdk:"time" json:"time,omitempty"`
	FirewallIPvsixNameRuleHopLimit          types.Object `tfsdk:"hop_limit" json:"hop-limit,omitempty"`
	FirewallIPvsixNameRuleIcmpvsix          types.Object `tfsdk:"icmpv6" json:"icmpv6,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallIPvsixNameRule) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  jump  |  Jump to another chain  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the
                      last chain  |
|  drop  |  Drop matching entries  |
|  queue  |  Enqueue packet to userspace  |
`,
		},

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Option to disable firewall rule

`,
		},

		"log": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Option to log packets matching rule

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable log  |
|  disable  |  Disable log  |
`,
		},

		"log_level": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
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
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
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
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DSCP value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value to match  |
|  <start-end>  |  DSCP range to match  |
`,
		},

		"dscp_exclude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `DSCP value not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value not to match  |
|  <start-end>  |  DSCP range not to match  |
`,
		},

		"packet_length": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length to match  |
|  <start-end>  |  Packet length range to match  |
`,
		},

		"packet_length_exclude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length not to match  |
|  <start-end>  |  Packet length range not to match  |
`,
		},

		"packet_type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
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
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Connection mark

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection-mark to match  |
`,
		},

		"jump_target": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"queue": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Queue target  |
`,
		},

		"queue_options": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
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
			Attributes: FirewallIPvsixNameRuleDestination{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleSource{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleFragment{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleInboundInterface{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleOutboundInterface{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleIPsec{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleLimit{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleConnectionStatus{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleRecent{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleState{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTime{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"hop_limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleHopLimit{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Hop limit

`,
		},

		"icmpv6": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleIcmpvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMPv6 type and code information

`,
		},
	}
}
