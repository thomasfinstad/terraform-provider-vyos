---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_router_advert_interface Resource - vyos"
subcategory: ""
description: |-
  IPv6 Router Advertisements (RAs) service
  Interface to send RA on
---

# vyos_service_router_advert_interface (Resource)

IPv6 Router Advertisements (RAs) service

Interface to send RA on



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Interface to send RA on

### Optional

- `default_lifetime` (String) Lifetime associated with the default router in units of seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:4-9000  |  Router Lifetime in seconds  |
|  0  |  Not a default router  |
- `default_preference` (String) Preference associated with the default router,

|  Format  |  Description  |
|----------|---------------|
|  low  |  Default router has low preference  |
|  medium  |  Default router has medium preference  |
|  high  |  Default router has high preference  |
- `dnssl` (String) DNS search list
- `hop_limit` (String) Set Hop Count field of the IP header for outgoing packets

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Unspecified (by this router)  |
|  u32:1-255  |  Value should represent current diameter of the Internet  |
- `interval` (Attributes) Set interval between unsolicited multicast RAs (see [below for nested schema](#nestedatt--interval))
- `link_mtu` (String) Link MTU value placed in RAs, exluded in RAs if unset

|  Format  |  Description  |
|----------|---------------|
|  u32:1280-9000  |  Link MTU value in RAs  |
- `managed_flag` (String) Hosts use the administered (stateful) protocol for address autoconfiguration in addition to any addresses autoconfigured using SLAAC
- `name_server` (String) Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |
- `name_server_lifetime` (String) Maximum duration how long the RDNSS entries are used

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Name-servers should no longer be used  |
|  u32:1-7200  |  Maximum interval in seconds  |
- `no_send_advert` (String) Do not send router adverts
- `other_config_flag` (String) Hosts use the administered (stateful) protocol for autoconfiguration of other (non-address) information
- `prefix` (Attributes Map) IPv6 prefix to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 prefix to be advertized  | (see [below for nested schema](#nestedatt--prefix))
- `reachable_time` (String) Time, in milliseconds, that a node assumes a neighbor is reachable after having received a reachability confirmation

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Reachable Time unspecified by this router  |
|  u32:1-3600000  |  Reachable Time value in RAs (in milliseconds)  |
- `retrans_timer` (String) Time in milliseconds between retransmitted Neighbor Solicitation messages

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Time, in milliseconds, between retransmitted Neighbor Solicitation messages  |
|  u32:1-4294967295  |  Minimum interval in milliseconds  |
- `route` (Attributes Map) IPv6 route to be advertised in Router Advertisements (RAs)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 route to be advertized  | (see [below for nested schema](#nestedatt--route))
- `source_address` (String) Use IPv6 address as source address. Useful with VRRP.

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IPv6 address to be advertized (must be configured on interface)  |

<a id="nestedatt--interval"></a>
### Nested Schema for `interval`

Optional:

- `max` (String) Maximum interval between unsolicited multicast RAs

|  Format  |  Description  |
|----------|---------------|
|  u32:4-1800  |  Maximum interval in seconds  |
- `min` (String) Minimum interval between unsolicited multicast RAs

|  Format  |  Description  |
|----------|---------------|
|  u32:3-1350  |  Minimum interval in seconds  |


<a id="nestedatt--prefix"></a>
### Nested Schema for `prefix`

Optional:

- `decrement_lifetime` (String) Lifetime is decremented by the number of seconds since the last RA - use in conjunction with a DHCPv6-PD prefix
- `deprecate_prefix` (String) Upon shutdown, this option will deprecate the prefix by announcing it in the shutdown RA
- `no_autonomous_flag` (String) Prefix can not be used for stateless address auto-configuration
- `no_on_link_flag` (String) Prefix can not be used for on-link determination
- `preferred_lifetime` (String) Time in seconds that the prefix will remain preferred

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Time in seconds that the prefix will remain preferred  |
|  infinity  |  Prefix will remain preferred forever  |
- `valid_lifetime` (String) Time in seconds that the prefix will remain valid

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Time in seconds that the prefix will remain valid  |
|  infinity  |  Prefix will remain preferred forever  |


<a id="nestedatt--route"></a>
### Nested Schema for `route`

Optional:

- `no_remove_route` (String) Do not announce this route with a zero second lifetime upon shutdown
- `route_preference` (String) Preference associated with the route,

|  Format  |  Description  |
|----------|---------------|
|  low  |  Route has low preference  |
|  medium  |  Route has medium preference  |
|  high  |  Route has high preference  |
- `valid_lifetime` (String) Time in seconds that the route will remain valid

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Time in seconds that the route will remain valid  |
|  infinity  |  Route will remain preferred forever  |
