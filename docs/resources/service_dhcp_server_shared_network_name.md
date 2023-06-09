---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_dhcp_server_shared_network_name Resource - vyos"
subcategory: ""
description: |-
  Dynamic Host Configuration Protocol (DHCP) for DHCP server
  Name of DHCP shared network
---

# vyos_service_dhcp_server_shared_network_name (Resource)

Dynamic Host Configuration Protocol (DHCP) for DHCP server

Name of DHCP shared network



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Name of DHCP shared network

### Optional

- `authoritative` (String) Option to make DHCP server authoritative for this physical network
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `disable` (String) Disable instance
- `domain_name` (String) Client Domain Name
- `domain_search` (String) Client Domain Name search list
- `name_server` (String) Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
- `ntp_server` (String) IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |
- `ping_check` (String) Sends ICMP Echo request to the address being assigned
- `shared_network_parameters` (String) Additional shared-network parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.
- `subnet` (Attributes Map) DHCP subnet for shared network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  | (see [below for nested schema](#nestedatt--subnet))

<a id="nestedatt--subnet"></a>
### Nested Schema for `subnet`

Optional:

- `bootfile_name` (String) Bootstrap file name
- `bootfile_server` (String) Server from which the initial boot file is to be loaded

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Bootfile server IPv4 address  |
|  hostname  |  Bootfile server FQDN  |
- `bootfile_size` (String) Bootstrap file size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16  |  Bootstrap file size in 512 byte blocks  |
- `client_prefix_length` (String) Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  DHCP client prefix length must be 0 to 32  |
- `default_router` (String) IP address of default router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default router IPv4 address  |
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `domain_name` (String) Client Domain Name
- `domain_search` (String) Client Domain Name search list
- `enable_failover` (String) Enable DHCP failover support for this subnet
- `exclude` (String) IP address to exclude from DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to exclude from lease range  |
- `ip_forwarding` (String) Enable IP forwarding on client
- `ipv6_only_preferred` (String) Disable IPv4 on IPv6 only hosts (RFC 8925)

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Seconds  |
- `lease` (String) Lease timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32  |  DHCP lease time in seconds  |
- `name_server` (String) Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
- `ntp_server` (String) IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |
- `ping_check` (String) Sends ICMP Echo request to the address being assigned
- `pop_server` (String) IP address of POP3 server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  POP3 server IPv4 address  |
- `range` (Attributes Map) DHCP lease range (see [below for nested schema](#nestedatt--subnet--range))
- `server_identifier` (String) Address for DHCP server identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  DHCP server identifier IPv4 address  |
- `smtp_server` (String) IP address of SMTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  SMTP server IPv4 address  |
- `static_mapping` (Attributes Map) Name of static mapping (see [below for nested schema](#nestedatt--subnet--static_mapping))
- `static_route` (Attributes Map) Classless static route destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  | (see [below for nested schema](#nestedatt--subnet--static_route))
- `subnet_parameters` (String) Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.
- `tftp_server_name` (String) TFTP server name

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  TFTP server IPv4 address  |
|  hostname  |  TFTP server FQDN  |
- `time_offset` (String) Client subnet offset in seconds from Coordinated Universal Time (UTC)

|  Format  |  Description  |
|----------|---------------|
|  [-]N  |  Time offset (number, may be negative)  |
- `time_server` (String) IP address of time server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Time server IPv4 address  |
- `vendor_option` (Attributes) Vendor Specific Options (see [below for nested schema](#nestedatt--subnet--vendor_option))
- `wins_server` (String) IP address for Windows Internet Name Service (WINS) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  WINS server IPv4 address  |
- `wpad_url` (String) Web Proxy Autodiscovery (WPAD) URL

<a id="nestedatt--subnet--range"></a>
### Nested Schema for `subnet.range`

Optional:

- `start` (String) First IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 start address of pool  |
- `stop` (String) Last IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 end address of pool  |


<a id="nestedatt--subnet--static_mapping"></a>
### Nested Schema for `subnet.static_mapping`

Optional:

- `disable` (String) Disable instance
- `ip_address` (String) Fixed IP address of static mapping

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address used in static mapping  |
- `mac_address` (String) Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |
- `static_mapping_parameters` (String) Additional static-mapping parameters for DHCP server. Will be placed inside the "host" block of the mapping. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.


<a id="nestedatt--subnet--static_route"></a>
### Nested Schema for `subnet.static_route`

Optional:

- `next_hop` (String) IP address of router to be used to reach the destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of router  |


<a id="nestedatt--subnet--vendor_option"></a>
### Nested Schema for `subnet.vendor_option`

Optional:

- `ubiquiti` (Attributes) Ubiquiti specific parameters (see [below for nested schema](#nestedatt--subnet--vendor_option--ubiquiti))

<a id="nestedatt--subnet--vendor_option--ubiquiti"></a>
### Nested Schema for `subnet.vendor_option.ubiquiti`

Optional:

- `unifi_controller` (String) Address of UniFi controller

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of UniFi controller  |
