// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceDNSForwarding describes the resource data model.
type ServiceDNSForwarding struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceDNSForwardingCacheSize                 types.Number `tfsdk:"cache_size" vyos:"cache-size,omitempty"`
	LeafServiceDNSForwardingDhcp                      types.List   `tfsdk:"dhcp" vyos:"dhcp,omitempty"`
	LeafServiceDNSForwardingDNSsixfourPrefix          types.String `tfsdk:"dns64_prefix" vyos:"dns64-prefix,omitempty"`
	LeafServiceDNSForwardingDNSsec                    types.String `tfsdk:"dnssec" vyos:"dnssec,omitempty"`
	LeafServiceDNSForwardingIgnoreHostsFile           types.Bool   `tfsdk:"ignore_hosts_file" vyos:"ignore-hosts-file,omitempty"`
	LeafServiceDNSForwardingNoServeRfconenineoneeight types.Bool   `tfsdk:"no_serve_rfc1918" vyos:"no-serve-rfc1918,omitempty"`
	LeafServiceDNSForwardingAllowFrom                 types.List   `tfsdk:"allow_from" vyos:"allow-from,omitempty"`
	LeafServiceDNSForwardingListenAddress             types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceDNSForwardingPort                      types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceDNSForwardingNegativeTTL               types.Number `tfsdk:"negative_ttl" vyos:"negative-ttl,omitempty"`
	LeafServiceDNSForwardingTimeout                   types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceDNSForwardingNameServer                types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafServiceDNSForwardingSourceAddress             types.List   `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafServiceDNSForwardingSystem                    types.Bool   `tfsdk:"system" vyos:"system,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceDNSForwardingDomain              bool `tfsdk:"-" vyos:"domain,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomain bool `tfsdk:"-" vyos:"authoritative-domain,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceDNSForwarding) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwarding) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"dns",

		"forwarding",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwarding) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"cache_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `DNS forwarding cache size

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  DNS forwarding cache size  |

`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		"dhcp": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interfaces whose DHCP client nameservers to forward requests to

`,
		},

		"dns64_prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Help to communicate between IPv6-only client and IPv4-only server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and /96 only prefix length  |

`,
		},

		"dnssec": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNSSEC mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  off  &emsp; |  No DNSSEC processing whatsoever!  |
    |  process-no-validate  &emsp; |  Respond with DNSSEC records to clients that ask for it. No validation done at all!  |
    |  process  &emsp; |  Respond with DNSSEC records to clients that ask for it. Validation for clients that request it.  |
    |  log-fail  &emsp; |  Similar behaviour to process, but validate RRSIGs on responses and log bogus responses.  |
    |  validate  &emsp; |  Full blown DNSSEC validation. Send SERVFAIL to clients on bogus responses.  |

`,

			// Default:          stringdefault.StaticString(`process-no-validate`),
			Computed: true,
		},

		"ignore_hosts_file": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not use local /etc/hosts file in name resolution

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_serve_rfc1918": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Makes the server authoritatively not aware of RFC1918 addresses

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"allow_from": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Networks allowed to query this server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IP address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		"listen_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to listen for incoming connections  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming connections  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`53`),
			Computed: true,
		},

		"negative_ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum amount of time negative entries are cached

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-7200  &emsp; |  Seconds to cache NXDOMAIN entries  |

`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of milliseconds to wait for a remote authoritative server to respond

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 10-60000  &emsp; |  Network timeout in milliseconds  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"name_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6  &emsp; |  Domain Name Server (DNS) IPv6 address  |

`,
		},

		"source_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local addresses from which to send DNS queries

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address from which to send traffic  |
    |  ipv6  &emsp; |  IPv6 address from which to send traffic  |

`,

			// Default:          stringdefault.StaticString(`0.0.0.0 ::`),
			Computed: true,
		},

		"system": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use system name servers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
