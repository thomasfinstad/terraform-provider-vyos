// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// NatDestinationRule describes the resource data model.
type NatDestinationRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafNatDestinationRuleDescrIPtion      types.String `tfsdk:"description" json:"description,omitempty"`
	LeafNatDestinationRuleDisable          types.String `tfsdk:"disable" json:"disable,omitempty"`
	LeafNatDestinationRuleExclude          types.String `tfsdk:"exclude" json:"exclude,omitempty"`
	LeafNatDestinationRuleLog              types.String `tfsdk:"log" json:"log,omitempty"`
	LeafNatDestinationRulePacketType       types.String `tfsdk:"packet_type" json:"packet-type,omitempty"`
	LeafNatDestinationRuleProtocol         types.String `tfsdk:"protocol" json:"protocol,omitempty"`
	LeafNatDestinationRuleInboundInterface types.String `tfsdk:"inbound_interface" json:"inbound-interface,omitempty"`

	// TagNodes

	// Nodes
	NodeNatDestinationRuleDestination *NatDestinationRuleDestination `tfsdk:"destination" json:"destination,omitempty"`
	NodeNatDestinationRuleSource      *NatDestinationRuleSource      `tfsdk:"source" json:"source,omitempty"`
	NodeNatDestinationRuleTranSLAtion *NatDestinationRuleTranSLAtion `tfsdk:"translation" json:"translation,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NatDestinationRule) GetVyosPath() []string {
	return []string{
		"nat",
		"destination",
		"rule",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Rule number for NAT

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of NAT rule  |

`,
		},

		// LeafNodes

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
			MarkdownDescription: `Disable instance

`,
		},

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Exclude packets matching this rule from NAT

`,
		},

		"log": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAT rule logging

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

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to NAT

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  ip  |  Internet Protocol, pseudo protocol number  |
|  hopopt  |  IPv6 Hop-by-Hop Option [RFC1883]  |
|  icmp  |  internet control message protocol  |
|  igmp  |  Internet Group Management  |
|  ggp  |  gateway-gateway protocol  |
|  ipencap  |  IP encapsulated in IP (officially IP)  |
|  st  |  ST datagram mode  |
|  tcp  |  transmission control protocol  |
|  egp  |  exterior gateway protocol  |
|  igp  |  any private interior gateway (Cisco)  |
|  pup  |  PARC universal packet protocol  |
|  udp  |  user datagram protocol  |
|  tcp_udp  |  Both TCP and UDP  |
|  hmp  |  host monitoring protocol  |
|  xns-idp  |  Xerox NS IDP  |
|  rdp  |  "reliable datagram" protocol  |
|  iso-tp4  |  ISO Transport Protocol class 4 [RFC905]  |
|  dccp  |  Datagram Congestion Control Prot. [RFC4340]  |
|  xtp  |  Xpress Transfer Protocol  |
|  ddp  |  Datagram Delivery Protocol  |
|  idpr-cmtp  |  IDPR Control Message Transport  |
|  Ipv6  |  Internet Protocol, version 6  |
|  ipv6-route  |  Routing Header for IPv6  |
|  ipv6-frag  |  Fragment Header for IPv6  |
|  idrp  |  Inter-Domain Routing Protocol  |
|  rsvp  |  Reservation Protocol  |
|  gre  |  General Routing Encapsulation  |
|  esp  |  Encap Security Payload [RFC2406]  |
|  ah  |  Authentication Header [RFC2402]  |
|  skip  |  SKIP  |
|  ipv6-icmp  |  ICMP for IPv6  |
|  ipv6-nonxt  |  No Next Header for IPv6  |
|  ipv6-opts  |  Destination Options for IPv6  |
|  rspf  |  Radio Shortest Path First (officially CPHB)  |
|  vmtp  |  Versatile Message Transport  |
|  eigrp  |  Enhanced Interior Routing Protocol (Cisco)  |
|  ospf  |  Open Shortest Path First IGP  |
|  ax.25  |  AX.25 frames  |
|  ipip  |  IP-within-IP Encapsulation Protocol  |
|  etherip  |  Ethernet-within-IP Encapsulation [RFC3378]  |
|  encap  |  Yet Another IP encapsulation [RFC1241]  |
|  99  |  Any private encryption scheme  |
|  pim  |  Protocol Independent Multicast  |
|  ipcomp  |  IP Payload Compression Protocol  |
|  vrrp  |  Virtual Router Redundancy Protocol [RFC5798]  |
|  l2tp  |  Layer Two Tunneling Protocol [RFC2661]  |
|  isis  |  IS-IS over IPv4  |
|  sctp  |  Stream Control Transmission Protocol  |
|  fc  |  Fibre Channel  |
|  mobility-header  |  Mobility Support for IPv6 [RFC3775]  |
|  udplite  |  UDP-Lite [RFC3828]  |
|  mpls-in-ip  |  MPLS-in-IP [RFC4023]  |
|  manet  |  MANET Protocols [RFC5498]  |
|  hip  |  Host Identity Protocol  |
|  shim6  |  Shim6 Protocol  |
|  wesp  |  Wrapped Encapsulating Security Payload  |
|  rohc  |  Robust Header Compression  |
|  u32:0-255  |  IP protocol number  |

`,

			// Default:          stringdefault.StaticString(`all`),
			Computed: true,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound interface of NAT traffic

`,
		},

		// TagNodes

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `NAT destination parameters

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `NAT source parameters

`,
		},

		"translation": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleTranSLAtion{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Inside NAT IP (destination NAT only)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NatDestinationRule) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafNatDestinationRuleDescrIPtion.IsNull() && !o.LeafNatDestinationRuleDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafNatDestinationRuleDescrIPtion.ValueString()
	}

	if !o.LeafNatDestinationRuleDisable.IsNull() && !o.LeafNatDestinationRuleDisable.IsUnknown() {
		jsonData["disable"] = o.LeafNatDestinationRuleDisable.ValueString()
	}

	if !o.LeafNatDestinationRuleExclude.IsNull() && !o.LeafNatDestinationRuleExclude.IsUnknown() {
		jsonData["exclude"] = o.LeafNatDestinationRuleExclude.ValueString()
	}

	if !o.LeafNatDestinationRuleLog.IsNull() && !o.LeafNatDestinationRuleLog.IsUnknown() {
		jsonData["log"] = o.LeafNatDestinationRuleLog.ValueString()
	}

	if !o.LeafNatDestinationRulePacketType.IsNull() && !o.LeafNatDestinationRulePacketType.IsUnknown() {
		jsonData["packet-type"] = o.LeafNatDestinationRulePacketType.ValueString()
	}

	if !o.LeafNatDestinationRuleProtocol.IsNull() && !o.LeafNatDestinationRuleProtocol.IsUnknown() {
		jsonData["protocol"] = o.LeafNatDestinationRuleProtocol.ValueString()
	}

	if !o.LeafNatDestinationRuleInboundInterface.IsNull() && !o.LeafNatDestinationRuleInboundInterface.IsUnknown() {
		jsonData["inbound-interface"] = o.LeafNatDestinationRuleInboundInterface.ValueString()
	}

	// Tags

	// Nodes

	if !reflect.ValueOf(o.NodeNatDestinationRuleDestination).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeNatDestinationRuleDestination)
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

	if !reflect.ValueOf(o.NodeNatDestinationRuleSource).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeNatDestinationRuleSource)
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

	if !reflect.ValueOf(o.NodeNatDestinationRuleTranSLAtion).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeNatDestinationRuleTranSLAtion)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["translation"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NatDestinationRule) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["description"]; ok {
		o.LeafNatDestinationRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["disable"]; ok {
		o.LeafNatDestinationRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleDisable = basetypes.NewStringNull()
	}

	if value, ok := jsonData["exclude"]; ok {
		o.LeafNatDestinationRuleExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["log"]; ok {
		o.LeafNatDestinationRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleLog = basetypes.NewStringNull()
	}

	if value, ok := jsonData["packet-type"]; ok {
		o.LeafNatDestinationRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRulePacketType = basetypes.NewStringNull()
	}

	if value, ok := jsonData["protocol"]; ok {
		o.LeafNatDestinationRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleProtocol = basetypes.NewStringNull()
	}

	if value, ok := jsonData["inbound-interface"]; ok {
		o.LeafNatDestinationRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleInboundInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := jsonData["destination"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeNatDestinationRuleDestination = &NatDestinationRuleDestination{}

		err = json.Unmarshal(subJSONStr, o.NodeNatDestinationRuleDestination)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["source"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeNatDestinationRuleSource = &NatDestinationRuleSource{}

		err = json.Unmarshal(subJSONStr, o.NodeNatDestinationRuleSource)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["translation"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeNatDestinationRuleTranSLAtion = &NatDestinationRuleTranSLAtion{}

		err = json.Unmarshal(subJSONStr, o.NodeNatDestinationRuleTranSLAtion)
		if err != nil {
			return err
		}
	}

	return nil
}
