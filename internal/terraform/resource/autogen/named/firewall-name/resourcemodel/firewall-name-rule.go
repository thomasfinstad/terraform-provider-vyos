// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallNameRule describes the resource data model.
type FirewallNameRule struct {
	// LeafNodes
	LeafFirewallNameRuleAction              types.String `tfsdk:"action" json:"action,omitempty"`
	LeafFirewallNameRuleDescrIPtion         types.String `tfsdk:"description" json:"description,omitempty"`
	LeafFirewallNameRuleDisable             types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafFirewallNameRuleLog                 types.String `tfsdk:"log" json:"log,omitempty"`
	LeafFirewallNameRuleLogLevel            types.String `tfsdk:"log_level" json:"log-level,omitempty"`
	LeafFirewallNameRuleProtocol            types.String `tfsdk:"protocol" json:"protocol,omitempty"`
	LeafFirewallNameRuleDscp                types.String `tfsdk:"dscp" json:"dscp,omitempty"`
	LeafFirewallNameRuleDscpExclude         types.String `tfsdk:"dscp_exclude" json:"dscp-exclude,omitempty"`
	LeafFirewallNameRulePacketLength        types.String `tfsdk:"packet_length" json:"packet-length,omitempty"`
	LeafFirewallNameRulePacketLengthExclude types.String `tfsdk:"packet_length_exclude" json:"packet-length-exclude,omitempty"`
	LeafFirewallNameRulePacketType          types.String `tfsdk:"packet_type" json:"packet-type,omitempty"`
	LeafFirewallNameRuleConnectionMark      types.String `tfsdk:"connection_mark" json:"connection-mark,omitempty"`
	LeafFirewallNameRuleJumpTarget          types.String `tfsdk:"jump_target" json:"jump-target,omitempty"`
	LeafFirewallNameRuleQueue               types.String `tfsdk:"queue" json:"queue,omitempty"`
	LeafFirewallNameRuleQueueOptions        types.String `tfsdk:"queue_options" json:"queue-options,omitempty"`

	// TagNodes

	// Nodes
	NodeFirewallNameRuleDestination       *FirewallNameRuleDestination       `tfsdk:"destination" json:"destination,omitempty"`
	NodeFirewallNameRuleSource            *FirewallNameRuleSource            `tfsdk:"source" json:"source,omitempty"`
	NodeFirewallNameRuleFragment          *FirewallNameRuleFragment          `tfsdk:"fragment" json:"fragment,omitempty"`
	NodeFirewallNameRuleInboundInterface  *FirewallNameRuleInboundInterface  `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`
	NodeFirewallNameRuleOutboundInterface *FirewallNameRuleOutboundInterface `tfsdk:"outbound_interface" json:"outbound-interface,omitempty"`
	NodeFirewallNameRuleIPsec             *FirewallNameRuleIPsec             `tfsdk:"ipsec" json:"ipsec,omitempty"`
	NodeFirewallNameRuleLimit             *FirewallNameRuleLimit             `tfsdk:"limit" json:"limit,omitempty"`
	NodeFirewallNameRuleConnectionStatus  *FirewallNameRuleConnectionStatus  `tfsdk:"connection_status" json:"connection-status,omitempty"`
	NodeFirewallNameRuleRecent            *FirewallNameRuleRecent            `tfsdk:"recent" json:"recent,omitempty"`
	NodeFirewallNameRuleState             *FirewallNameRuleState             `tfsdk:"state" json:"state,omitempty"`
	NodeFirewallNameRuleTCP               *FirewallNameRuleTCP               `tfsdk:"tcp" json:"tcp,omitempty"`
	NodeFirewallNameRuleTime              *FirewallNameRuleTime              `tfsdk:"time" json:"time,omitempty"`
	NodeFirewallNameRuleIcmp              *FirewallNameRuleIcmp              `tfsdk:"icmp" json:"icmp,omitempty"`
	NodeFirewallNameRuleTTL               *FirewallNameRuleTTL               `tfsdk:"ttl" json:"ttl,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallNameRule) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: FirewallNameRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"inbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleInboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match inbound-interface

`,
		},

		"outbound_interface": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleOutboundInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match outbound-interface

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"connection_status": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleConnectionStatus{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Connection status

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleState{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleIcmp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: FirewallNameRuleTTL{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallNameRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafFirewallNameRuleAction.IsNull() && !o.LeafFirewallNameRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafFirewallNameRuleAction.ValueString()
	}

	if !o.LeafFirewallNameRuleDescrIPtion.IsNull() && !o.LeafFirewallNameRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafFirewallNameRuleDescrIPtion.ValueString()
	}

	if !o.LeafFirewallNameRuleDisable.IsNull() && !o.LeafFirewallNameRuleDisable.IsUnknown() {
		jsonData["disable"] = o.LeafFirewallNameRuleDisable.ValueString()
	}

	if !o.LeafFirewallNameRuleLog.IsNull() && !o.LeafFirewallNameRuleLog.IsUnknown() {
		jsonData["log"] = o.LeafFirewallNameRuleLog.ValueString()
	}

	if !o.LeafFirewallNameRuleLogLevel.IsNull() && !o.LeafFirewallNameRuleLogLevel.IsUnknown() {
		jsonData["log-level"] = o.LeafFirewallNameRuleLogLevel.ValueString()
	}

	if !o.LeafFirewallNameRuleProtocol.IsNull() && !o.LeafFirewallNameRuleProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafFirewallNameRuleProtocol.ValueString()
	}

	if !o.LeafFirewallNameRuleDscp.IsNull() && !o.LeafFirewallNameRuleDscp.IsUnknown() {
		jsonData["dscp"] = o.LeafFirewallNameRuleDscp.ValueString()
	}

	if !o.LeafFirewallNameRuleDscpExclude.IsNull() && !o.LeafFirewallNameRuleDscpExclude.IsUnknown() {
		jsonData["dscp-exclude"] = o.LeafFirewallNameRuleDscpExclude.ValueString()
	}

	if !o.LeafFirewallNameRulePacketLength.IsNull() && !o.LeafFirewallNameRulePacketLength.IsUnknown() {
		jsonData["packet-length"] = o.LeafFirewallNameRulePacketLength.ValueString()
	}

	if !o.LeafFirewallNameRulePacketLengthExclude.IsNull() && !o.LeafFirewallNameRulePacketLengthExclude.IsUnknown() {
		jsonData["packet-length-exclude"] = o.LeafFirewallNameRulePacketLengthExclude.ValueString()
	}

	if !o.LeafFirewallNameRulePacketType.IsNull() && !o.LeafFirewallNameRulePacketType.IsUnknown() {
		jsonData["packet-type"] = o.LeafFirewallNameRulePacketType.ValueString()
	}

	if !o.LeafFirewallNameRuleConnectionMark.IsNull() && !o.LeafFirewallNameRuleConnectionMark.IsUnknown() {
		jsonData["connection-mark"] = o.LeafFirewallNameRuleConnectionMark.ValueString()
	}

	if !o.LeafFirewallNameRuleJumpTarget.IsNull() && !o.LeafFirewallNameRuleJumpTarget.IsUnknown() {
		jsonData["jump-target"] = o.LeafFirewallNameRuleJumpTarget.ValueString()
	}

	if !o.LeafFirewallNameRuleQueue.IsNull() && !o.LeafFirewallNameRuleQueue.IsUnknown() {
		jsonData["queue"] = o.LeafFirewallNameRuleQueue.ValueString()
	}

	if !o.LeafFirewallNameRuleQueueOptions.IsNull() && !o.LeafFirewallNameRuleQueueOptions.IsUnknown() {
		jsonData["queue-options"] = o.LeafFirewallNameRuleQueueOptions.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeFirewallNameRuleDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleDestination)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["destination"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleSource)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["source"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleFragment).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleFragment)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["fragment"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleInboundInterface).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleInboundInterface)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["inbound-interface"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleOutboundInterface).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleOutboundInterface)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["outbound-interface"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleIPsec).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleIPsec)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ipsec"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleLimit).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleLimit)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["limit"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleConnectionStatus).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleConnectionStatus)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["connection-status"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleRecent).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleRecent)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["recent"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleState).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleState)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["state"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleTCP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleTCP)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["tcp"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleTime).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleTime)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["time"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleIcmp).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleIcmp)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["icmp"] = subData
	}

	if !reflect.ValueOf(o.NodeFirewallNameRuleTTL).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeFirewallNameRuleTTL)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["ttl"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallNameRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafFirewallNameRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafFirewallNameRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafFirewallNameRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["log"]; ok {
		o.LeafFirewallNameRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleLog = basetypes.NewStringNull()
	}

	if value, ok := jsonData["log-level"]; ok {
		o.LeafFirewallNameRuleLogLevel = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleLogLevel = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafFirewallNameRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleProtocol = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dscp"]; ok {
		o.LeafFirewallNameRuleDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDscp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dscp-exclude"]; ok {
		o.LeafFirewallNameRuleDscpExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleDscpExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-length"]; ok {
		o.LeafFirewallNameRulePacketLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRulePacketLength = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-length-exclude"]; ok {
		o.LeafFirewallNameRulePacketLengthExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRulePacketLengthExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-type"]; ok {
		o.LeafFirewallNameRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRulePacketType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["connection-mark"]; ok {
		o.LeafFirewallNameRuleConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleConnectionMark = basetypes.NewStringNull()
	}

	if value, ok := jsonData["jump-target"]; ok {
		o.LeafFirewallNameRuleJumpTarget = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleJumpTarget = basetypes.NewStringNull()
	}

	if value, ok := jsonData["queue"]; ok {
		o.LeafFirewallNameRuleQueue = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleQueue = basetypes.NewStringNull()
	}

	if value, ok := jsonData["queue-options"]; ok {
		o.LeafFirewallNameRuleQueueOptions = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafFirewallNameRuleQueueOptions = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleDestination = &FirewallNameRuleDestination{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleSource = &FirewallNameRuleSource{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleSource)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["fragment"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleFragment = &FirewallNameRuleFragment{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleFragment)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["inbound-interface"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleInboundInterface = &FirewallNameRuleInboundInterface{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleInboundInterface)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["outbound-interface"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleOutboundInterface = &FirewallNameRuleOutboundInterface{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleOutboundInterface)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipsec"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleIPsec = &FirewallNameRuleIPsec{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleIPsec)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["limit"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleLimit = &FirewallNameRuleLimit{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleLimit)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["connection-status"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleConnectionStatus = &FirewallNameRuleConnectionStatus{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleConnectionStatus)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["recent"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleRecent = &FirewallNameRuleRecent{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleRecent)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["state"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleState = &FirewallNameRuleState{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleState)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["tcp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleTCP = &FirewallNameRuleTCP{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleTCP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["time"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleTime = &FirewallNameRuleTime{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleTime)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["icmp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleIcmp = &FirewallNameRuleIcmp{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleIcmp)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ttl"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeFirewallNameRuleTTL = &FirewallNameRuleTTL{}

		err = json.Unmarshal(subJSONStr, o.NodeFirewallNameRuleTTL)
		if err != nil {
			return err
		}
	}

	return nil
}
