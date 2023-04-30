// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesOpenvpn describes the resource data model.
type InterfacesOpenvpn struct {
	// LeafNodes
	InterfacesOpenvpnDescrIPtion       customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	InterfacesOpenvpnDeviceType        customtypes.CustomStringValue `tfsdk:"device_type" json:"device-type,omitempty"`
	InterfacesOpenvpnDisable           customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	InterfacesOpenvpnHash              customtypes.CustomStringValue `tfsdk:"hash" json:"hash,omitempty"`
	InterfacesOpenvpnLocalHost         customtypes.CustomStringValue `tfsdk:"local_host" json:"local-host,omitempty"`
	InterfacesOpenvpnLocalPort         customtypes.CustomStringValue `tfsdk:"local_port" json:"local-port,omitempty"`
	InterfacesOpenvpnMode              customtypes.CustomStringValue `tfsdk:"mode" json:"mode,omitempty"`
	InterfacesOpenvpnOpenvpnOption     customtypes.CustomStringValue `tfsdk:"openvpn_option" json:"openvpn-option,omitempty"`
	InterfacesOpenvpnPersistentTunnel  customtypes.CustomStringValue `tfsdk:"persistent_tunnel" json:"persistent-tunnel,omitempty"`
	InterfacesOpenvpnProtocol          customtypes.CustomStringValue `tfsdk:"protocol" json:"protocol,omitempty"`
	InterfacesOpenvpnRemoteAddress     customtypes.CustomStringValue `tfsdk:"remote_address" json:"remote-address,omitempty"`
	InterfacesOpenvpnRemoteHost        customtypes.CustomStringValue `tfsdk:"remote_host" json:"remote-host,omitempty"`
	InterfacesOpenvpnRemotePort        customtypes.CustomStringValue `tfsdk:"remote_port" json:"remote-port,omitempty"`
	InterfacesOpenvpnSharedSecretKey   customtypes.CustomStringValue `tfsdk:"shared_secret_key" json:"shared-secret-key,omitempty"`
	InterfacesOpenvpnUseLzoCompression customtypes.CustomStringValue `tfsdk:"use_lzo_compression" json:"use-lzo-compression,omitempty"`
	InterfacesOpenvpnRedirect          customtypes.CustomStringValue `tfsdk:"redirect" json:"redirect,omitempty"`
	InterfacesOpenvpnVrf               customtypes.CustomStringValue `tfsdk:"vrf" json:"vrf,omitempty"`

	// TagNodes
	InterfacesOpenvpnLocalAddress types.Map `tfsdk:"local_address" json:"local-address,omitempty"`

	// Nodes
	InterfacesOpenvpnAuthentication      types.Object `tfsdk:"authentication" json:"authentication,omitempty"`
	InterfacesOpenvpnEncryption          types.Object `tfsdk:"encryption" json:"encryption,omitempty"`
	InterfacesOpenvpnIP                  types.Object `tfsdk:"ip" json:"ip,omitempty"`
	InterfacesOpenvpnIPvsix              types.Object `tfsdk:"ipv6" json:"ipv6,omitempty"`
	InterfacesOpenvpnMirror              types.Object `tfsdk:"mirror" json:"mirror,omitempty"`
	InterfacesOpenvpnKeepAlive           types.Object `tfsdk:"keep_alive" json:"keep-alive,omitempty"`
	InterfacesOpenvpnReplaceDefaultRoute types.Object `tfsdk:"replace_default_route" json:"replace-default-route,omitempty"`
	InterfacesOpenvpnServer              types.Object `tfsdk:"server" json:"server,omitempty"`
	InterfacesOpenvpnTLS                 types.Object `tfsdk:"tls" json:"tls,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesOpenvpn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
`,
		},

		"device_type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `OpenVPN interface device-type

|  Format  |  Description  |
|----------|---------------|
|  tun  |  TUN device, required for OSI layer 3  |
|  tap  |  TAP device, required for OSI layer 2  |
`,

			// Default:          stringdefault.StaticString(`tun`),
			Computed: true,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administratively disable interface

`,
		},

		"hash": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Hashing Algorithm

|  Format  |  Description  |
|----------|---------------|
|  md5  |  MD5 algorithm  |
|  sha1  |  SHA-1 algorithm  |
|  sha256  |  SHA-256 algorithm  |
|  sha384  |  SHA-384 algorithm  |
|  sha512  |  SHA-512 algorithm  |
`,
		},

		"local_host": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Local IP address to accept connections (all if not set)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Local IPv4 address  |
|  ipv6  |  Local IPv6 address  |
`,
		},

		"local_port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Local port number to accept connections

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
`,
		},

		"mode": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `OpenVPN mode of operation

|  Format  |  Description  |
|----------|---------------|
|  site-to-site  |  Site-to-site mode  |
|  client  |  Client in client-server mode  |
|  server  |  Server in client-server mode  |
`,
		},

		"openvpn_option": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Additional OpenVPN options. You must use the syntax of openvpn.conf in this text-field. Using this without proper knowledge may result in a crashed OpenVPN server. Check system log to look for errors.

`,
		},

		"persistent_tunnel": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Do not close and reopen interface (TUN/TAP device) on client restarts

`,
		},

		"protocol": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `OpenVPN communication protocol

|  Format  |  Description  |
|----------|---------------|
|  udp  |  UDP  |
|  tcp-passive  |  TCP and accepts connections passively  |
|  tcp-active  |  TCP and initiates connections actively  |
`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		"remote_address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of remote end of tunnel

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Remote end IPv4 address  |
|  ipv6  |  Remote end IPv6 address  |
`,
		},

		"remote_host": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remote host to connect to (dynamic if not set)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of remote host  |
|  ipv6  |  IPv6 address of remote host  |
|  txt  |  Hostname of remote host  |
`,
		},

		"remote_port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Remote port number to connect to

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
`,
		},

		"shared_secret_key": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Secret key shared with remote end of tunnel

`,
		},

		"use_lzo_compression": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Use fast LZO compression on this TUN/TAP interface

`,
		},

		"redirect": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
		},

		"vrf": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |
`,
		},

		// TagNodes

		"local_address": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesOpenvpnLocalAddress{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Local IP address of tunnel (IPv4 or IPv6)

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnAuthentication{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},

		"encryption": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnEncryption{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Data Encryption settings

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnIP{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnIPvsix{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnMirror{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},

		"keep_alive": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnKeepAlive{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Keepalive helper options

`,
		},

		"replace_default_route": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnReplaceDefaultRoute{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `OpenVPN tunnel to be used as the default route

`,
		},

		"server": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServer{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Server-mode options

`,
		},

		"tls": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnTLS{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Transport Layer Security (TLS) options

`,
		},
	}
}
