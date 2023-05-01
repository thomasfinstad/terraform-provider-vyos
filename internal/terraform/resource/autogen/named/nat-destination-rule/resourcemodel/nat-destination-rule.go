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

// NatDestinationRule describes the resource data model.
type NatDestinationRule struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	LeafNatDestinationRuleDescrIPtion      types.String `tfsdk:"description"`
	LeafNatDestinationRuleDisable          types.String `tfsdk:"disable"`
	LeafNatDestinationRuleExclude          types.String `tfsdk:"exclude"`
	LeafNatDestinationRuleLog              types.String `tfsdk:"log"`
	LeafNatDestinationRulePacketType       types.String `tfsdk:"packet_type"`
	LeafNatDestinationRuleProtocol         types.String `tfsdk:"protocol"`
	LeafNatDestinationRuleInboundInterface types.String `tfsdk:"inbound_interface"`

	// TagNodes

	// Nodes
	NodeNatDestinationRuleDestination types.Object `tfsdk:"destination"`
	NodeNatDestinationRuleSource      types.Object `tfsdk:"source"`
	NodeNatDestinationRuleTranSLAtion types.Object `tfsdk:"translation"`
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

// TerraformToVyos converts terraform data to vyos data
func (o *NatDestinationRule) TerraformToVyos(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	tflog.Error(ctx, "TerraformToVyos", map[string]interface{}{"Path": []string{"nat", "destination", "rule"}})

	vyosData := make(map[string]interface{})

	// Leafs
	if !(o.LeafNatDestinationRuleDescrIPtion.IsNull() || o.LeafNatDestinationRuleDescrIPtion.IsUnknown()) {
		vyosData["description"] = o.LeafNatDestinationRuleDescrIPtion.ValueString()
	}
	if !(o.LeafNatDestinationRuleDisable.IsNull() || o.LeafNatDestinationRuleDisable.IsUnknown()) {
		vyosData["disable"] = o.LeafNatDestinationRuleDisable.ValueString()
	}
	if !(o.LeafNatDestinationRuleExclude.IsNull() || o.LeafNatDestinationRuleExclude.IsUnknown()) {
		vyosData["exclude"] = o.LeafNatDestinationRuleExclude.ValueString()
	}
	if !(o.LeafNatDestinationRuleLog.IsNull() || o.LeafNatDestinationRuleLog.IsUnknown()) {
		vyosData["log"] = o.LeafNatDestinationRuleLog.ValueString()
	}
	if !(o.LeafNatDestinationRulePacketType.IsNull() || o.LeafNatDestinationRulePacketType.IsUnknown()) {
		vyosData["packet-type"] = o.LeafNatDestinationRulePacketType.ValueString()
	}
	if !(o.LeafNatDestinationRuleProtocol.IsNull() || o.LeafNatDestinationRuleProtocol.IsUnknown()) {
		vyosData["protocol"] = o.LeafNatDestinationRuleProtocol.ValueString()
	}
	if !(o.LeafNatDestinationRuleInboundInterface.IsNull() || o.LeafNatDestinationRuleInboundInterface.IsUnknown()) {
		vyosData["inbound-interface"] = o.LeafNatDestinationRuleInboundInterface.ValueString()
	}

	// Tags

	// Nodes
	if !(o.NodeNatDestinationRuleDestination.IsNull() || o.NodeNatDestinationRuleDestination.IsUnknown()) {
		var subModel NatDestinationRuleDestination
		diags.Append(o.NodeNatDestinationRuleDestination.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["destination"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeNatDestinationRuleSource.IsNull() || o.NodeNatDestinationRuleSource.IsUnknown()) {
		var subModel NatDestinationRuleSource
		diags.Append(o.NodeNatDestinationRuleSource.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["source"] = subModel.TerraformToVyos(ctx, diags)
	}
	if !(o.NodeNatDestinationRuleTranSLAtion.IsNull() || o.NodeNatDestinationRuleTranSLAtion.IsUnknown()) {
		var subModel NatDestinationRuleTranSLAtion
		diags.Append(o.NodeNatDestinationRuleTranSLAtion.As(ctx, &subModel, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		vyosData["translation"] = subModel.TerraformToVyos(ctx, diags)
	}

	// Return compiled data
	return vyosData
}

// VyosToTerraform converts vyos data to terraform data
func (o *NatDestinationRule) VyosToTerraform(ctx context.Context, diags *diag.Diagnostics, vyosData map[string]interface{}) {
	tflog.Error(ctx, "VyosToTerraform begin", map[string]interface{}{"Path": []string{"nat", "destination", "rule"}})

	// Leafs
	if value, ok := vyosData["description"]; ok {
		o.LeafNatDestinationRuleDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleDescrIPtion = basetypes.NewStringNull()
	}
	if value, ok := vyosData["disable"]; ok {
		o.LeafNatDestinationRuleDisable = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleDisable = basetypes.NewStringNull()
	}
	if value, ok := vyosData["exclude"]; ok {
		o.LeafNatDestinationRuleExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleExclude = basetypes.NewStringNull()
	}
	if value, ok := vyosData["log"]; ok {
		o.LeafNatDestinationRuleLog = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleLog = basetypes.NewStringNull()
	}
	if value, ok := vyosData["packet-type"]; ok {
		o.LeafNatDestinationRulePacketType = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRulePacketType = basetypes.NewStringNull()
	}
	if value, ok := vyosData["protocol"]; ok {
		o.LeafNatDestinationRuleProtocol = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleProtocol = basetypes.NewStringNull()
	}
	if value, ok := vyosData["inbound-interface"]; ok {
		o.LeafNatDestinationRuleInboundInterface = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafNatDestinationRuleInboundInterface = basetypes.NewStringNull()
	}

	// Tags

	// Nodes
	if value, ok := vyosData["destination"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, NatDestinationRuleDestination{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeNatDestinationRuleDestination = data

	} else {
		o.NodeNatDestinationRuleDestination = basetypes.NewObjectNull(NatDestinationRuleDestination{}.AttributeTypes())
	}
	if value, ok := vyosData["source"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, NatDestinationRuleSource{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeNatDestinationRuleSource = data

	} else {
		o.NodeNatDestinationRuleSource = basetypes.NewObjectNull(NatDestinationRuleSource{}.AttributeTypes())
	}
	if value, ok := vyosData["translation"]; ok {
		data, d := basetypes.NewObjectValueFrom(ctx, NatDestinationRuleTranSLAtion{}.AttributeTypes(), value.(map[string]interface{}))
		diags.Append(d...)
		o.NodeNatDestinationRuleTranSLAtion = data

	} else {
		o.NodeNatDestinationRuleTranSLAtion = basetypes.NewObjectNull(NatDestinationRuleTranSLAtion{}.AttributeTypes())
	}

	tflog.Error(ctx, "VyosToTerraform end", map[string]interface{}{"Path": []string{"nat", "destination", "rule"}})
}

// AttributeTypes generates the attribute types for the resource at this level
func (o NatDestinationRule) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		// Leafs
		"description":       types.StringType,
		"disable":           types.StringType,
		"exclude":           types.StringType,
		"log":               types.StringType,
		"packet_type":       types.StringType,
		"protocol":          types.StringType,
		"inbound_interface": types.StringType,

		// Tags

		// Nodes
		"destination": types.ObjectType{AttrTypes: NatDestinationRuleDestination{}.AttributeTypes()},
		"source":      types.ObjectType{AttrTypes: NatDestinationRuleSource{}.AttributeTypes()},
		"translation": types.ObjectType{AttrTypes: NatDestinationRuleTranSLAtion{}.AttributeTypes()},
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
