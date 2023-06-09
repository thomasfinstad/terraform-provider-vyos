// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDhcpServerSharedNetworkNameSubnet describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnet struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName        types.String `tfsdk:"bootfile_name" json:"bootfile-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer      types.String `tfsdk:"bootfile_server" json:"bootfile-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize        types.String `tfsdk:"bootfile_size" json:"bootfile-size,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength  types.String `tfsdk:"client_prefix_length" json:"client-prefix-length,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter       types.String `tfsdk:"default_router" json:"default-router,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDomainName          types.String `tfsdk:"domain_name" json:"domain-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch        types.String `tfsdk:"domain_search" json:"domain-search,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion         types.String `tfsdk:"description" json:"description,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetNameServer          types.String `tfsdk:"name_server" json:"name-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover      types.String `tfsdk:"enable_failover" json:"enable-failover,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetExclude             types.String `tfsdk:"exclude" json:"exclude,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding        types.String `tfsdk:"ip_forwarding" json:"ip-forwarding,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetLease               types.String `tfsdk:"lease" json:"lease,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer           types.String `tfsdk:"ntp_server" json:"ntp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck           types.String `tfsdk:"ping_check" json:"ping-check,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetPopServer           types.String `tfsdk:"pop_server" json:"pop-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier    types.String `tfsdk:"server_identifier" json:"server-identifier,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer          types.String `tfsdk:"smtp_server" json:"smtp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred types.String `tfsdk:"ipv6_only_preferred" json:"ipv6-only-preferred,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters    types.String `tfsdk:"subnet_parameters" json:"subnet-parameters,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName      types.String `tfsdk:"tftp_server_name" json:"tftp-server-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset          types.String `tfsdk:"time_offset" json:"time-offset,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer          types.String `tfsdk:"time_server" json:"time-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer          types.String `tfsdk:"wins_server" json:"wins-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL             types.String `tfsdk:"wpad_url" json:"wpad-url,omitempty"`

	// TagNodes
	TagServiceDhcpServerSharedNetworkNameSubnetRange         *map[string]ServiceDhcpServerSharedNetworkNameSubnetRange         `tfsdk:"range" json:"range,omitempty"`
	TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping *map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticMapping `tfsdk:"static_mapping" json:"static-mapping,omitempty"`
	TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute   *map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticRoute   `tfsdk:"static_route" json:"static-route,omitempty"`

	// Nodes
	NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption *ServiceDhcpServerSharedNetworkNameSubnetVendorOption `tfsdk:"vendor_option" json:"vendor-option,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnet) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"bootfile_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file name

`,
		},

		"bootfile_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server from which the initial boot file is to be loaded

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Bootfile server IPv4 address  |
|  hostname  |  Bootfile server FQDN  |

`,
		},

		"bootfile_size": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16  |  Bootstrap file size in 512 byte blocks  |

`,
		},

		"client_prefix_length": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  DHCP client prefix length must be 0 to 32  |

`,
		},

		"default_router": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of default router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default router IPv4 address  |

`,
		},

		"domain_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name

`,
		},

		"domain_search": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name search list

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

		"name_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |

`,
		},

		"enable_failover": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable DHCP failover support for this subnet

`,
		},

		"exclude": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address to exclude from DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to exclude from lease range  |

`,
		},

		"ip_forwarding": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enable IP forwarding on client

`,
		},

		"lease": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Lease timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32  |  DHCP lease time in seconds  |

`,

			// Default:          stringdefault.StaticString(`86400`),
			Computed: true,
		},

		"ntp_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |

`,
		},

		"ping_check": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
		},

		"pop_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of POP3 server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  POP3 server IPv4 address  |

`,
		},

		"server_identifier": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address for DHCP server identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  DHCP server identifier IPv4 address  |

`,
		},

		"smtp_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of SMTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  SMTP server IPv4 address  |

`,
		},

		"ipv6_only_preferred": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Seconds  |

`,
		},

		"subnet_parameters": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
		},

		"tftp_server_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TFTP server name

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  TFTP server IPv4 address  |
|  hostname  |  TFTP server FQDN  |

`,
		},

		"time_offset": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

|  Format  |  Description  |
|----------|---------------|
|  [-]N  |  Time offset (number, may be negative)  |

`,
		},

		"time_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of time server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Time server IPv4 address  |

`,
		},

		"wins_server": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  WINS server IPv4 address  |

`,
		},

		"wpad_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
		},

		// TagNodes

		"range": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetRange{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `DHCP lease range

`,
		},

		"static_mapping": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}.ResourceSchemaAttributes(),
			},
			Optional: true,
			MarkdownDescription: `Name of static mapping

`,
		},

		"static_route": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}.ResourceSchemaAttributes(),
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
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDhcpServerSharedNetworkNameSubnet) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	// Leafs

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.IsUnknown() {
		jsonData["bootfile-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.IsUnknown() {
		jsonData["bootfile-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.IsUnknown() {
		jsonData["bootfile-size"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.IsUnknown() {
		jsonData["client-prefix-length"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.IsUnknown() {
		jsonData["default-router"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.IsUnknown() {
		jsonData["domain-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.IsUnknown() {
		jsonData["domain-search"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.IsUnknown() {
		jsonData["description"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.IsUnknown() {
		jsonData["name-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.IsUnknown() {
		jsonData["enable-failover"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.IsUnknown() {
		jsonData["exclude"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.IsUnknown() {
		jsonData["ip-forwarding"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.IsUnknown() {
		jsonData["lease"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetLease.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.IsUnknown() {
		jsonData["ntp-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.IsUnknown() {
		jsonData["ping-check"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.IsUnknown() {
		jsonData["pop-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.IsUnknown() {
		jsonData["server-identifier"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.IsUnknown() {
		jsonData["smtp-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.IsUnknown() {
		jsonData["ipv6-only-preferred"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.IsUnknown() {
		jsonData["subnet-parameters"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.IsUnknown() {
		jsonData["tftp-server-name"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.IsUnknown() {
		jsonData["time-offset"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.IsUnknown() {
		jsonData["time-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.IsUnknown() {
		jsonData["wins-server"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer.ValueString()
	}

	if !o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.IsNull() && !o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.IsUnknown() {
		jsonData["wpad-url"] = o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL.ValueString()
	}

	// Tags

	if !reflect.ValueOf(o.TagServiceDhcpServerSharedNetworkNameSubnetRange).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceDhcpServerSharedNetworkNameSubnetRange)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["range"] = subData
	}

	if !reflect.ValueOf(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["static-mapping"] = subData
	}

	if !reflect.ValueOf(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute).IsZero() {
		subJSONStr, err := json.Marshal(o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["static-route"] = subData
	}

	// Nodes

	if !reflect.ValueOf(o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption).IsZero() {
		subJSONStr, err := json.Marshal(o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption)
		if err != nil {
			return nil, err
		}

		subData := make(map[string]interface{})
		err = json.Unmarshal(subJSONStr, &subData)
		if err != nil {
			return nil, err
		}
		jsonData["vendor-option"] = subData
	}

	// Return compiled data
	ret, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDhcpServerSharedNetworkNameSubnet) UnmarshalJSON(jsonStr []byte) error {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &jsonData)
	if err != nil {
		return err
	}

	// Leafs

	if value, ok := jsonData["bootfile-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["bootfile-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["bootfile-size"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetBootfileSize = basetypes.NewStringNull()
	}

	if value, ok := jsonData["client-prefix-length"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetClientPrefixLength = basetypes.NewStringNull()
	}

	if value, ok := jsonData["default-router"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDefaultRouter = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["domain-search"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDomainSearch = basetypes.NewStringNull()
	}

	if value, ok := jsonData["description"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetDescrIPtion = basetypes.NewStringNull()
	}

	if value, ok := jsonData["name-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNameServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["enable-failover"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetEnableFailover = basetypes.NewStringNull()
	}

	if value, ok := jsonData["exclude"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetExclude = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ip-forwarding"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPForwarding = basetypes.NewStringNull()
	}

	if value, ok := jsonData["lease"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetLease = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetLease = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ntp-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetNtpServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ping-check"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPingCheck = basetypes.NewStringNull()
	}

	if value, ok := jsonData["pop-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetPopServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["server-identifier"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetServerIDentifier = basetypes.NewStringNull()
	}

	if value, ok := jsonData["smtp-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSMTPServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["ipv6-only-preferred"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetIPvsixOnlyPreferred = basetypes.NewStringNull()
	}

	if value, ok := jsonData["subnet-parameters"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetSubnetParameters = basetypes.NewStringNull()
	}

	if value, ok := jsonData["tftp-server-name"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTftpServerName = basetypes.NewStringNull()
	}

	if value, ok := jsonData["time-offset"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeOffset = basetypes.NewStringNull()
	}

	if value, ok := jsonData["time-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetTimeServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["wins-server"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWinsServer = basetypes.NewStringNull()
	}

	if value, ok := jsonData["wpad-url"]; ok {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL = basetypes.NewStringValue(value.(string))
	} else {
		o.LeafServiceDhcpServerSharedNetworkNameSubnetWpadURL = basetypes.NewStringNull()
	}

	// Tags
	if value, ok := jsonData["range"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceDhcpServerSharedNetworkNameSubnetRange = &map[string]ServiceDhcpServerSharedNetworkNameSubnetRange{}

		err = json.Unmarshal(subJSONStr, o.TagServiceDhcpServerSharedNetworkNameSubnetRange)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["static-mapping"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping = &map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}

		err = json.Unmarshal(subJSONStr, o.TagServiceDhcpServerSharedNetworkNameSubnetStaticMapping)
		if err != nil {
			return err
		}
	}
	if value, ok := jsonData["static-route"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute = &map[string]ServiceDhcpServerSharedNetworkNameSubnetStaticRoute{}

		err = json.Unmarshal(subJSONStr, o.TagServiceDhcpServerSharedNetworkNameSubnetStaticRoute)
		if err != nil {
			return err
		}
	}

	// Nodes
	if value, ok := jsonData["vendor-option"]; ok {
		subJSONStr, err := json.Marshal(value)
		if err != nil {
			return err
		}

		o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption = &ServiceDhcpServerSharedNetworkNameSubnetVendorOption{}

		err = json.Unmarshal(subJSONStr, o.NodeServiceDhcpServerSharedNetworkNameSubnetVendorOption)
		if err != nil {
			return err
		}
	}

	return nil
}
