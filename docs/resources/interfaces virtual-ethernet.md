---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces virtual-ethernet Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_interfaces virtual-ethernet (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (String) IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
Format: dhcp
Dynamic Host Configuration Protocol
Format: dhcpv6
Dynamic Host Configuration Protocol for IPv6
- `description` (String) Description

Format: txt
Description
- `dhcp_options` (Attributes) DHCP client settings/options (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcpv6_options` (Attributes) DHCPv6 client settings/options (see [below for nested schema](#nestedatt--dhcpv6_options))
- `disable` (String) Administratively disable interface
- `peer_name` (String) Virtual ethernet peer interface name

Format: txt
Name of peer interface
- `vrf` (String) VRF instance name

Format: txt
VRF instance name

<a id="nestedatt--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Optional:

- `client_id` (String) Identifier used by client to identify itself to the DHCP server
- `default_route_distance` (String) Distance for installed default route

Format: u32:1-255
Distance for the default route from DHCP server
- `host_name` (String) Override system host-name sent to DHCP server
- `mtu` (String) Use MTU value from DHCP server - ignore interface setting
- `no_default_route` (String) Do not install default route to system
- `reject` (String) IP addresses or subnets from which to reject DHCP leases

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
- `vendor_class_id` (String) Identify the vendor client type to the DHCP server


<a id="nestedatt--dhcpv6_options"></a>
### Nested Schema for `dhcpv6_options`

Optional:

- `duid` (String) DHCP unique identifier (DUID) to be sent by dhcpv6 client

Format: duid
DHCP unique identifier (DUID)
- `parameters_only` (String) Acquire only config parameters, no address
- `pd` (Attributes Map) DHCPv6 prefix delegation interface statement

Format: instance number
Prefix delegation instance (>= 0) (see [below for nested schema](#nestedatt--dhcpv6_options--pd))
- `rapid_commit` (String) Wait for immediate reply instead of advertisements
- `temporary` (String) IPv6 temporary address

<a id="nestedatt--dhcpv6_options--pd"></a>
### Nested Schema for `dhcpv6_options.pd`

Optional:

- `interface` (Attributes Map) Delegate IPv6 prefix from provider to this interface (see [below for nested schema](#nestedatt--dhcpv6_options--pd--interface))
- `length` (String) Request IPv6 prefix length from peer

Format: u32:32-64
Length of delegated prefix

<a id="nestedatt--dhcpv6_options--pd--interface"></a>
### Nested Schema for `dhcpv6_options.pd.interface`

Optional:

- `address` (String) Local interface address assigned to interface (default: EUI-64)

Format: >0
Used to form IPv6 interface address
- `sla_id` (String) Interface site-Level aggregator (SLA)

Format: u32:0-65535
Decimal integer which fits in the length of SLA IDs