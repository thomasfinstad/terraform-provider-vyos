// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PolicyRouteRule describes the resource data model.
type PolicyRouteRule struct {
	// LeafNodes
	LeafPolicyRouteRuleAction              types.String `tfsdk:"action" json:"action,omitempty"`
	LeafPolicyRouteRuleDescrIPtion         types.String `tfsdk:"description" json:"description,omitempty"`
	LeafPolicyRouteRuleDisable             types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafPolicyRouteRuleLog                 types.String `tfsdk:"log" json:"log,omitempty"`
	LeafPolicyRouteRuleProtocol            types.String `tfsdk:"protocol" json:"protocol,omitempty"`
	LeafPolicyRouteRuleDscp                types.String `tfsdk:"dscp" json:"dscp,omitempty"`
	LeafPolicyRouteRuleDscpExclude         types.String `tfsdk:"dscp_exclude" json:"dscp-exclude,omitempty"`
	LeafPolicyRouteRulePacketLength        types.String `tfsdk:"packet_length" json:"packet-length,omitempty"`
	LeafPolicyRouteRulePacketLengthExclude types.String `tfsdk:"packet_length_exclude" json:"packet-length-exclude,omitempty"`
	LeafPolicyRouteRulePacketType          types.String `tfsdk:"packet_type" json:"packet-type,omitempty"`
	LeafPolicyRouteRuleConnectionMark      types.String `tfsdk:"connection_mark" json:"connection-mark,omitempty"`

	// TagNodes

	// Nodes
	NodePolicyRouteRuleDestination *PolicyRouteRuleDestination `tfsdk:"destination" json:"destination,omitempty"`
	NodePolicyRouteRuleSource      *PolicyRouteRuleSource      `tfsdk:"source" json:"source,omitempty"`
	NodePolicyRouteRuleFragment    *PolicyRouteRuleFragment    `tfsdk:"fragment" json:"fragment,omitempty"`
	NodePolicyRouteRuleIPsec       *PolicyRouteRuleIPsec       `tfsdk:"ipsec" json:"ipsec,omitempty"`
	NodePolicyRouteRuleLimit       *PolicyRouteRuleLimit       `tfsdk:"limit" json:"limit,omitempty"`
	NodePolicyRouteRuleRecent      *PolicyRouteRuleRecent      `tfsdk:"recent" json:"recent,omitempty"`
	NodePolicyRouteRuleSet         *PolicyRouteRuleSet         `tfsdk:"set" json:"set,omitempty"`
	NodePolicyRouteRuleState       *PolicyRouteRuleState       `tfsdk:"state" json:"state,omitempty"`
	NodePolicyRouteRuleTCP         *PolicyRouteRuleTCP         `tfsdk:"tcp" json:"tcp,omitempty"`
	NodePolicyRouteRuleTime        *PolicyRouteRuleTime        `tfsdk:"time" json:"time,omitempty"`
	NodePolicyRouteRuleIcmp        *PolicyRouteRuleIcmp        `tfsdk:"icmp" json:"icmp,omitempty"`
	NodePolicyRouteRuleTTL         *PolicyRouteRuleTTL         `tfsdk:"ttl" json:"ttl,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteRule) ResourceSchemaAttributes() map[string]schema.Attribute {
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
			Attributes: PolicyRouteRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
		},

		"fragment": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleFragment{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IP fragment match

`,
		},

		"ipsec": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleIPsec{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inbound IPsec packets

`,
		},

		"limit": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Rate limit using a token bucket filter

`,
		},

		"recent": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleRecent{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Parameters for matching recently seen sources

`,
		},

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleSet{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
		},

		"state": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleState{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Session state

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
		},

		"time": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTime{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to match rule

`,
		},

		"icmp": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleIcmp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ICMP type and code information

`,
		},

		"ttl": schema.SingleNestedAttribute{
			Attributes: PolicyRouteRuleTTL{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Time to live limit

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PolicyRouteRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafPolicyRouteRuleAction.IsNull() && !o.LeafPolicyRouteRuleAction.IsUnknown() {
		jsonData["action"] = o.LeafPolicyRouteRuleAction.ValueString()
	}

	if !o.LeafPolicyRouteRuleDescrIPtion.IsNull() && !o.LeafPolicyRouteRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafPolicyRouteRuleDescrIPtion.ValueString()
	}

	if !o.LeafPolicyRouteRuleDisable.IsNull() && !o.LeafPolicyRouteRuleDisable.IsUnknown() {
		jsonData["disable"] = o.LeafPolicyRouteRuleDisable.ValueString()
	}

	if !o.LeafPolicyRouteRuleLog.IsNull() && !o.LeafPolicyRouteRuleLog.IsUnknown() {
		jsonData["log"] = o.LeafPolicyRouteRuleLog.ValueString()
	}

	if !o.LeafPolicyRouteRuleProtocol.IsNull() && !o.LeafPolicyRouteRuleProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafPolicyRouteRuleProtocol.ValueString()
	}

	if !o.LeafPolicyRouteRuleDscp.IsNull() && !o.LeafPolicyRouteRuleDscp.IsUnknown() {
		jsonData["dscp"] = o.LeafPolicyRouteRuleDscp.ValueString()
	}

	if !o.LeafPolicyRouteRuleDscpExclude.IsNull() && !o.LeafPolicyRouteRuleDscpExclude.IsUnknown() {
		jsonData["dscp-exclude"] = o.LeafPolicyRouteRuleDscpExclude.ValueString()
	}

	if !o.LeafPolicyRouteRulePacketLength.IsNull() && !o.LeafPolicyRouteRulePacketLength.IsUnknown() {
		jsonData["packet-length"] = o.LeafPolicyRouteRulePacketLength.ValueString()
	}

	if !o.LeafPolicyRouteRulePacketLengthExclude.IsNull() && !o.LeafPolicyRouteRulePacketLengthExclude.IsUnknown() {
		jsonData["packet-length-exclude"] = o.LeafPolicyRouteRulePacketLengthExclude.ValueString()
	}

	if !o.LeafPolicyRouteRulePacketType.IsNull() && !o.LeafPolicyRouteRulePacketType.IsUnknown() {
		jsonData["packet-type"] = o.LeafPolicyRouteRulePacketType.ValueString()
	}

	if !o.LeafPolicyRouteRuleConnectionMark.IsNull() && !o.LeafPolicyRouteRuleConnectionMark.IsUnknown() {
		jsonData["connection-mark"] = o.LeafPolicyRouteRuleConnectionMark.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodePolicyRouteRuleDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleDestination)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleSource)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleFragment).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleFragment)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleIPsec).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleIPsec)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleLimit).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleLimit)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleRecent).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleRecent)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleSet).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleSet)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["set"] = subData
	}

	if !reflect.ValueOf(o.NodePolicyRouteRuleState).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleState)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleTCP).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleTCP)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleTime).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleTime)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleIcmp).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleIcmp)
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

	if !reflect.ValueOf(o.NodePolicyRouteRuleTTL).IsZero() {
		subJSONStr, err := json.Marshal(o.NodePolicyRouteRuleTTL)
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
func (o *PolicyRouteRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["action"]; ok {
		o.LeafPolicyRouteRuleAction = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleAction = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafPolicyRouteRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafPolicyRouteRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["log"]; ok {
		o.LeafPolicyRouteRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleLog = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafPolicyRouteRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleProtocol = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dscp"]; ok {
		o.LeafPolicyRouteRuleDscp = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDscp = basetypes.NewStringNull()
	}

	if value, ok := jsonData["dscp-exclude"]; ok {
		o.LeafPolicyRouteRuleDscpExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleDscpExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-length"]; ok {
		o.LeafPolicyRouteRulePacketLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRulePacketLength = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-length-exclude"]; ok {
		o.LeafPolicyRouteRulePacketLengthExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRulePacketLengthExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-type"]; ok {
		o.LeafPolicyRouteRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRulePacketType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["connection-mark"]; ok {
		o.LeafPolicyRouteRuleConnectionMark = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafPolicyRouteRuleConnectionMark = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleDestination = &PolicyRouteRuleDestination{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleSource = &PolicyRouteRuleSource{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleSource)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["fragment"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleFragment = &PolicyRouteRuleFragment{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleFragment)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ipsec"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleIPsec = &PolicyRouteRuleIPsec{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleIPsec)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["limit"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleLimit = &PolicyRouteRuleLimit{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleLimit)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["recent"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleRecent = &PolicyRouteRuleRecent{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleRecent)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["set"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleSet = &PolicyRouteRuleSet{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleSet)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["state"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleState = &PolicyRouteRuleState{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleState)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["tcp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleTCP = &PolicyRouteRuleTCP{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleTCP)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["time"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleTime = &PolicyRouteRuleTime{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleTime)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["icmp"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleIcmp = &PolicyRouteRuleIcmp{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleIcmp)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["ttl"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodePolicyRouteRuleTTL = &PolicyRouteRuleTTL{}

		err = json.Unmarshal(subJSONStr, o.NodePolicyRouteRuleTTL)
		if err != nil {
			return err
		}
	}

	return nil
}
