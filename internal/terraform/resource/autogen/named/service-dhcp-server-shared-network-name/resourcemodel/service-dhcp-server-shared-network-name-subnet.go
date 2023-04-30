// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDhcpServerSharedNetworkNameSubnet describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnet struct {
	// LeafNodes
	ServiceDhcpServerSharedNetworkNameSubnetBootfileName        customtypes.CustomStringValue `tfsdk:"bootfile_name" json:"bootfile-name,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetBootfileServer      customtypes.CustomStringValue `tfsdk:"bootfile_server" json:"bootfile-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetBootfileSize        customtypes.CustomStringValue `tfsdk:"bootfile_size" json:"bootfile-size,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength  customtypes.CustomStringValue `tfsdk:"client_prefix_length" json:"client-prefix-length,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetDefaultRouter       customtypes.CustomStringValue `tfsdk:"default_router" json:"default-router,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetDomainName          customtypes.CustomStringValue `tfsdk:"domain_name" json:"domain-name,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetDomainSearch        customtypes.CustomStringValue `tfsdk:"domain_search" json:"domain-search,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetDescrIPtion         customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetNameServer          customtypes.CustomStringValue `tfsdk:"name_server" json:"name-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetEnableFailover      customtypes.CustomStringValue `tfsdk:"enable_failover" json:"enable-failover,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetExclude             customtypes.CustomStringValue `tfsdk:"exclude" json:"exclude,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetIPForwarding        customtypes.CustomStringValue `tfsdk:"ip_forwarding" json:"ip-forwarding,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetLease               customtypes.CustomStringValue `tfsdk:"lease" json:"lease,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetNtpServer           customtypes.CustomStringValue `tfsdk:"ntp_server" json:"ntp-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetPingCheck           customtypes.CustomStringValue `tfsdk:"ping_check" json:"ping-check,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetPopServer           customtypes.CustomStringValue `tfsdk:"pop_server" json:"pop-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetServerIDentifier    customtypes.CustomStringValue `tfsdk:"server_identifier" json:"server-identifier,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetSMTPServer          customtypes.CustomStringValue `tfsdk:"smtp_server" json:"smtp-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred customtypes.CustomStringValue `tfsdk:"ipv6_only_preferred" json:"ipv6-only-preferred,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetSubnetParameters    customtypes.CustomStringValue `tfsdk:"subnet_parameters" json:"subnet-parameters,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetTftpServerName      customtypes.CustomStringValue `tfsdk:"tftp_server_name" json:"tftp-server-name,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetTimeOffset          customtypes.CustomStringValue `tfsdk:"time_offset" json:"time-offset,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetTimeServer          customtypes.CustomStringValue `tfsdk:"time_server" json:"time-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetWinsServer          customtypes.CustomStringValue `tfsdk:"wins_server" json:"wins-server,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetWpadURL             customtypes.CustomStringValue `tfsdk:"wpad_url" json:"wpad-url,omitempty"`

	// TagNodes
	ServiceDhcpServerSharedNetworkNameSubnetRange         types.Map `tfsdk:"range" json:"range,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetStaticMapping types.Map `tfsdk:"static_mapping" json:"static-mapping,omitempty"`
	ServiceDhcpServerSharedNetworkNameSubnetStaticRoute   types.Map `tfsdk:"static_route" json:"static-route,omitempty"`

	// Nodes
	ServiceDhcpServerSharedNetworkNameSubnetVendorOption types.Object `tfsdk:"vendor_option" json:"vendor-option,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnet) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bootfile_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Bootstrap file name

`,
		},

		"bootfile_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Server from which the initial boot file is to be loaded

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Bootfile server IPv4 address  |
|  hostname  |  Bootfile server FQDN  |
`,
		},

		"bootfile_size": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Bootstrap file size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16  |  Bootstrap file size in 512 byte blocks  |
`,
		},

		"client_prefix_length": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  DHCP client prefix length must be 0 to 32  |
`,
		},

		"default_router": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of default router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default router IPv4 address  |
`,
		},

		"domain_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client Domain Name

`,
		},

		"domain_search": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client Domain Name search list

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

		"name_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
`,
		},

		"enable_failover": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable DHCP failover support for this subnet

`,
		},

		"exclude": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address to exclude from DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to exclude from lease range  |
`,
		},

		"ip_forwarding": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable IP forwarding on client

`,
		},

		"lease": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Lease timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32  |  DHCP lease time in seconds  |
`,

			// Default:          stringdefault.StaticString(`86400`),
			Computed: true,
		},

		"ntp_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |
`,
		},

		"ping_check": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
		},

		"pop_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of POP3 server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  POP3 server IPv4 address  |
`,
		},

		"server_identifier": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Address for DHCP server identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  DHCP server identifier IPv4 address  |
`,
		},

		"smtp_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of SMTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  SMTP server IPv4 address  |
`,
		},

		"ipv6_only_preferred": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Seconds  |
`,
		},

		"subnet_parameters": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		"tftp_server_name": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `TFTP server name

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  TFTP server IPv4 address  |
|  hostname  |  TFTP server FQDN  |
`,
		},

		"time_offset": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

|  Format  |  Description  |
|----------|---------------|
|  [-]N  |  Time offset (number, may be negative)  |
`,
		},

		"time_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address of time server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Time server IPv4 address  |
`,
		},

		"wins_server": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  WINS server IPv4 address  |
`,
		},

		"wpad_url": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetRange{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCP lease range

`,
		},

		"static_mapping": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Name of static mapping

`,
		},

		"static_route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Classless static route destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
`,
		},

		// Nodes

		"vendor_option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
		},
	}
}
