// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsNhrpTunnel describes the resource data model.
type ProtocolsNhrpTunnel struct {
	// LeafNodes
	ProtocolsNhrpTunnelCiscoAuthentication customtypes.CustomStringValue `tfsdk:"cisco_authentication" json:"cisco-authentication,omitempty"`
	ProtocolsNhrpTunnelHoldingTime         customtypes.CustomStringValue `tfsdk:"holding_time" json:"holding-time,omitempty"`
	ProtocolsNhrpTunnelMulticast           customtypes.CustomStringValue `tfsdk:"multicast" json:"multicast,omitempty"`
	ProtocolsNhrpTunnelNonCaching          customtypes.CustomStringValue `tfsdk:"non_caching" json:"non-caching,omitempty"`
	ProtocolsNhrpTunnelRedirect            customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	ProtocolsNhrpTunnelShortcutDestination customtypes.CustomStringValue `tfsdk:"shortcut_destination" json:"shortcut-destination,omitempty"`
	ProtocolsNhrpTunnelShortcut            customtypes.CustomStringValue `tfsdk:"shortcut" json:"shortcut,omitempty"`

	// TagNodes
	ProtocolsNhrpTunnelDynamicMap     types.Map `tfsdk:"dynamic_map" json:"dynamic-map,omitempty"`
	ProtocolsNhrpTunnelMap            types.Map `tfsdk:"map" json:"map,omitempty"`
	ProtocolsNhrpTunnelShortcutTarget types.Map `tfsdk:"shortcut_target" json:"shortcut-target,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsNhrpTunnel) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"cisco_authentication": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Pass phrase for cisco authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Pass phrase for cisco authentication  |
`,
		},

		"holding_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Holding time in seconds

`,
		},

		"multicast": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Set multicast for NHRP

`,
		},

		"non_caching": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `This can be used to reduce memory consumption on big NBMA subnets

`,
		},

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable sending of Cisco style NHRP Traffic Indication packets

`,
		},

		"shortcut_destination": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `This instructs opennhrp to reply with authorative answers on NHRP Resolution Requests destined to addresses in this interface

`,
		},

		"shortcut": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route

`,
		},

		// TagNodes

		"dynamic_map": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsNhrpTunnelDynamicMap{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Set an HUB tunnel address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Set the IP address and prefix length  |
`,
		},

		"map": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsNhrpTunnelMap{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Set an HUB tunnel address

`,
		},

		"shortcut_target": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ProtocolsNhrpTunnelShortcutTarget{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Defines an off-NBMA network prefix for which the GRE interface will act as a gateway

`,
		},

		// Nodes

	}
}
