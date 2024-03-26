// Package resourcemodel code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallIPvfourForwardFilterRule{}

// FirewallIPvfourForwardFilterRule describes the resource data model.
type FirewallIPvfourForwardFilterRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleAction              types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleDscp                types.List   `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleDscpExclude         types.List   `tfsdk:"dscp_exclude" vyos:"dscp-exclude,omitempty"`
	LeafFirewallIPvfourForwardFilterRulePacketLength        types.List   `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafFirewallIPvfourForwardFilterRulePacketLengthExclude types.List   `tfsdk:"packet_length_exclude" vyos:"packet-length-exclude,omitempty"`
	LeafFirewallIPvfourForwardFilterRulePacketType          types.String `tfsdk:"packet_type" vyos:"packet-type,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleMark                types.String `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleConnectionMark      types.List   `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleConntrackHelper     types.List   `tfsdk:"conntrack_helper" vyos:"conntrack-helper,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleQueue               types.Number `tfsdk:"queue" vyos:"queue,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleQueueOptions        types.List   `tfsdk:"queue_options" vyos:"queue-options,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleLog                 types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleProtocol            types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleState               types.List   `tfsdk:"state" vyos:"state,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleJumpTarget          types.String `tfsdk:"jump_target" vyos:"jump-target,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleOffloadTarget       types.String `tfsdk:"offload_target" vyos:"offload-target,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourForwardFilterRuleFragment          *FirewallIPvfourForwardFilterRuleFragment          `tfsdk:"fragment" vyos:"fragment,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleLimit             *FirewallIPvfourForwardFilterRuleLimit             `tfsdk:"limit" vyos:"limit,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleLogOptions        *FirewallIPvfourForwardFilterRuleLogOptions        `tfsdk:"log_options" vyos:"log-options,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleConnectionStatus  *FirewallIPvfourForwardFilterRuleConnectionStatus  `tfsdk:"connection_status" vyos:"connection-status,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleRecent            *FirewallIPvfourForwardFilterRuleRecent            `tfsdk:"recent" vyos:"recent,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleSynproxy          *FirewallIPvfourForwardFilterRuleSynproxy          `tfsdk:"synproxy" vyos:"synproxy,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleTCP               *FirewallIPvfourForwardFilterRuleTCP               `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleTime              *FirewallIPvfourForwardFilterRuleTime              `tfsdk:"time" vyos:"time,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleTTL               *FirewallIPvfourForwardFilterRuleTTL               `tfsdk:"ttl" vyos:"ttl,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleAddAddressToGroup *FirewallIPvfourForwardFilterRuleAddAddressToGroup `tfsdk:"add_address_to_group" vyos:"add-address-to-group,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleDestination       *FirewallIPvfourForwardFilterRuleDestination       `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleIcmp              *FirewallIPvfourForwardFilterRuleIcmp              `tfsdk:"icmp" vyos:"icmp,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleSource            *FirewallIPvfourForwardFilterRuleSource            `tfsdk:"source" vyos:"source,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleInboundInterface  *FirewallIPvfourForwardFilterRuleInboundInterface  `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleIPsec             *FirewallIPvfourForwardFilterRuleIPsec             `tfsdk:"ipsec" vyos:"ipsec,omitempty"`
	NodeFirewallIPvfourForwardFilterRuleOutboundInterface *FirewallIPvfourForwardFilterRuleOutboundInterface `tfsdk:"outbound_interface" vyos:"outbound-interface,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallIPvfourForwardFilterRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallIPvfourForwardFilterRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallIPvfourForwardFilterRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallIPvfourForwardFilterRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallIPvfourForwardFilterRule) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"ipv4",

		"forward",

		"filter",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallIPvfourForwardFilterRule) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"rule_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `IPv4 Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			Description: `IPv4 Firewall forward filter rule number

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  1-999999  |  Number for this firewall rule  |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  reject    |  Reject matching entries                                                        |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
    |  offload   |  Offload packet via flowtable                                                   |
    |  synproxy  |  Synproxy connections                                                           |
`,
			Description: `Rule action

    |  Format    |  Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    |  accept    |  Accept matching entries                                                        |
    |  continue  |  Continue parsing next rule                                                     |
    |  jump      |  Jump to another chain                                                          |
    |  reject    |  Reject matching entries                                                        |
    |  return    |  Return from the current chain and continue at the next rule of the last chain  |
    |  drop      |  Drop matching entries                                                          |
    |  queue     |  Enqueue packet to userspace                                                    |
    |  offload   |  Offload packet via flowtable                                                   |
    |  synproxy  |  Synproxy connections                                                           |
`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"dscp": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value

    |  Format       |  Description          |
    |---------------|-----------------------|
    |  0-63         |  DSCP value to match  |
    |  <start-end>  |  DSCP range to match  |
`,
			Description: `DSCP value

    |  Format       |  Description          |
    |---------------|-----------------------|
    |  0-63         |  DSCP value to match  |
    |  <start-end>  |  DSCP range to match  |
`,
		},

		"dscp_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DSCP value not to match

    |  Format       |  Description              |
    |---------------|---------------------------|
    |  0-63         |  DSCP value not to match  |
    |  <start-end>  |  DSCP range not to match  |
`,
			Description: `DSCP value not to match

    |  Format       |  Description              |
    |---------------|---------------------------|
    |  0-63         |  DSCP value not to match  |
    |  <start-end>  |  DSCP range not to match  |
`,
		},

		"packet_length": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data to match

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  1-65535      |  Packet length to match        |
    |  <start-end>  |  Packet length range to match  |
`,
			Description: `Payload size in bytes, including header and data to match

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  1-65535      |  Packet length to match        |
    |  <start-end>  |  Packet length range to match  |
`,
		},

		"packet_length_exclude": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Payload size in bytes, including header and data not to match

    |  Format       |  Description                       |
    |---------------|------------------------------------|
    |  1-65535      |  Packet length not to match        |
    |  <start-end>  |  Packet length range not to match  |
`,
			Description: `Payload size in bytes, including header and data not to match

    |  Format       |  Description                       |
    |---------------|------------------------------------|
    |  1-65535      |  Packet length not to match        |
    |  <start-end>  |  Packet length range not to match  |
`,
		},

		"packet_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Packet type

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  broadcast  |  Match broadcast packet type                      |
    |  host       |  Match host packet type, addressed to local host  |
    |  multicast  |  Match multicast packet type                      |
    |  other      |  Match packet addressed to another host           |
`,
			Description: `Packet type

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  broadcast  |  Match broadcast packet type                      |
    |  host       |  Match host packet type, addressed to local host  |
    |  multicast  |  Match multicast packet type                      |
    |  other      |  Match packet addressed to another host           |
`,
		},

		"mark": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Firewall mark

    |  Format         |  Description                            |
    |-----------------|-----------------------------------------|
    |  0-2147483647   |  Firewall mark to match                 |
    |  !0-2147483647  |  Inverted Firewall mark to match        |
    |  <start-end>    |  Firewall mark range to match           |
    |  !<start-end>   |  Firewall mark inverted range to match  |
`,
			Description: `Firewall mark

    |  Format         |  Description                            |
    |-----------------|-----------------------------------------|
    |  0-2147483647   |  Firewall mark to match                 |
    |  !0-2147483647  |  Inverted Firewall mark to match        |
    |  <start-end>    |  Firewall mark range to match           |
    |  !<start-end>   |  Firewall mark inverted range to match  |
`,
		},

		"connection_mark": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Connection mark

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  0-2147483647  |  Connection-mark to match  |
`,
			Description: `Connection mark

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  0-2147483647  |  Connection-mark to match  |
`,
		},

		"conntrack_helper": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Match related traffic from conntrack helpers

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ftp     |  Related traffic from FTP helper     |
    |  h323    |  Related traffic from H.323 helper   |
    |  pptp    |  Related traffic from PPTP helper    |
    |  nfs     |  Related traffic from NFS helper     |
    |  rtsp    |  Related traffic from RTSP helper    |
    |  sip     |  Related traffic from SIP helper     |
    |  tftp    |  Related traffic from TFTP helper    |
    |  sqlnet  |  Related traffic from SQLNet helper  |
`,
			Description: `Match related traffic from conntrack helpers

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ftp     |  Related traffic from FTP helper     |
    |  h323    |  Related traffic from H.323 helper   |
    |  pptp    |  Related traffic from PPTP helper    |
    |  nfs     |  Related traffic from NFS helper     |
    |  rtsp    |  Related traffic from RTSP helper    |
    |  sip     |  Related traffic from SIP helper     |
    |  tftp    |  Related traffic from TFTP helper    |
    |  sqlnet  |  Related traffic from SQLNet helper  |
`,
		},

		"queue": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

    |  Format   |  Description   |
    |-----------|----------------|
    |  0-65535  |  Queue target  |
`,
			Description: `Queue target to use. Action queue must be defined to use this setting

    |  Format   |  Description   |
    |-----------|----------------|
    |  0-65535  |  Queue target  |
`,
		},

		"queue_options": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Options used for queue target. Action queue must be defined to use this setting

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  bypass  |  Let packets go through if userspace application cannot back off  |
    |  fanout  |  Distribute packets between several queues                        |
`,
			Description: `Options used for queue target. Action queue must be defined to use this setting

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  bypass  |  Let packets go through if userspace application cannot back off  |
    |  fanout  |  Distribute packets between several queues                        |
`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Description: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name, number, or "all")

    |  Format       |  Description         |
    |---------------|----------------------|
    |  all          |  All IP protocols    |
    |  tcp_udp      |  Both TCP and UDP    |
    |  0-255        |  IP protocol number  |
    |  <protocol>   |  IP protocol name    |
    |  !<protocol>  |  IP protocol name    |
`,
			Description: `Protocol to match (protocol name, number, or "all")

    |  Format       |  Description         |
    |---------------|----------------------|
    |  all          |  All IP protocols    |
    |  tcp_udp      |  Both TCP and UDP    |
    |  0-255        |  IP protocol number  |
    |  <protocol>   |  IP protocol name    |
    |  !<protocol>  |  IP protocol name    |
`,
		},

		"state": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Session state

    |  Format       |  Description        |
    |---------------|---------------------|
    |  established  |  Established state  |
    |  invalid      |  Invalid state      |
    |  new          |  New state          |
    |  related      |  Related state      |
`,
			Description: `Session state

    |  Format       |  Description        |
    |---------------|---------------------|
    |  established  |  Established state  |
    |  invalid      |  Invalid state      |
    |  new          |  New state          |
    |  related      |  Related state      |
`,
		},

		"jump_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
			Description: `Set jump target. Action jump must be defined to use this setting

`,
		},

		"offload_target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set flowtable offload target. Action offload must be defined to use this setting

`,
			Description: `Set flowtable offload target. Action offload must be defined to use this setting

`,
		},

		// Nodes

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleFragment{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
			Description: `IP fragment match

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
			Description: `Rate limit using a token bucket filter

`,
		},

		"log_options": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleLogOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log options

`,
			Description: `Log options

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleConnectionStatus{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
			Description: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleRecent{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
			Description: `Parameters for matching recently seen sources

`,
		},

		"synproxy": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleSynproxy{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Synproxy options

`,
			Description: `Synproxy options

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP options to match

`,
			Description: `TCP options to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleTime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
			Description: `Time to match rule

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleTTL{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
			Description: `Time to live limit

`,
		},

		"add_address_to_group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleAddAddressToGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add ip address to dynamic address-group

`,
			Description: `Add ip address to dynamic address-group

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleIcmp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
			Description: `ICMP type and code information

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleInboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
			Description: `Match inbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleIPsec{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
			Description: `Inbound IPsec packets

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourForwardFilterRuleOutboundInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
			Description: `Match outbound-interface

`,
		},
	}
}
