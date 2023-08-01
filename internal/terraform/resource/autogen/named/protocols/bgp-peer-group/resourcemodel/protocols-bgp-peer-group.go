// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpPeerGroup describes the resource data model.
type ProtocolsBgpPeerGroup struct {
	SelfIdentifier types.String `tfsdk:"peer_group_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpPeerGroupDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafProtocolsBgpPeerGroupDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafProtocolsBgpPeerGroupDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafProtocolsBgpPeerGroupEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafProtocolsBgpPeerGroupGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafProtocolsBgpPeerGroupOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafProtocolsBgpPeerGroupPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBgpPeerGroupPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafProtocolsBgpPeerGroupRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpPeerGroupShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBgpPeerGroupUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpPeerGroupLocalAs   bool `tfsdk:"local_as" vyos:"local-as,child"`
	ExistsTagProtocolsBgpPeerGroupLocalRole bool `tfsdk:"local_role" vyos:"local-role,child"`

	// Nodes
	NodeProtocolsBgpPeerGroupAddressFamily *ProtocolsBgpPeerGroupAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeProtocolsBgpPeerGroupBfd           *ProtocolsBgpPeerGroupBfd           `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeProtocolsBgpPeerGroupCapability    *ProtocolsBgpPeerGroupCapability    `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeProtocolsBgpPeerGroupTTLSecURIty   *ProtocolsBgpPeerGroupTTLSecURIty   `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpPeerGroup) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"peer-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `peer_group_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"peer_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of peer-group

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable_capability_negotiation": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable capability negotiation with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_connected_check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable check to see if eBGP peer address is a connected route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_multihop": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Number of hops  |

`,
		},

		"graceful_restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable BGP graceful restart at peer level  |
    |  disable  &emsp; |  Disable BGP graceful restart at peer level  |
    |  restart-helper  &emsp; |  Enable BGP graceful restart helper only functionality  |

`,
		},

		"override_capability": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate a session with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Neighbor AS number  |
    |  external  &emsp; |  Any AS different from the local AS  |
    |  internal  &emsp; |  Neighbor AS number  |

`,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update_source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP of routing updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address of route source  |
    |  ipv6  &emsp; |  IPv6 address of route source  |
    |  txt  &emsp; |  Interface as route source  |

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupTTLSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroup) UnmarshalJSON(_ []byte) error {
	return nil
}
