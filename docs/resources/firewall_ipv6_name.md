---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_firewall_ipv6_name Resource - vyos"
subcategory: ""
description: |-
  Firewall
  IPv6 firewall rule-set name
---

# vyos_firewall_ipv6_name (Resource)

Firewall

IPv6 firewall rule-set name



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) IPv6 firewall rule-set name

### Optional

- `default_action` (String) Default-action for rule-set

|  Format  |  Description  |
|----------|---------------|
|  drop  |  Drop if no prior rules are hit  |
|  jump  |  Jump to another chain if no prior rules are hit  |
|  reject  |  Drop and notify source if no prior rules are hit  |
|  return  |  Return from the current chain and continue at the next rule of the last                   chain  |
|  accept  |  Accept if no prior rules are hit  |
- `default_jump_target` (String) Set jump target. Action jump must be defined in default-action to use this
                setting
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `enable_default_log` (String) Log packets hitting default-action
- `rule` (Attributes Map) Firewall rule number (IPv6)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number for this Firewall rule  | (see [below for nested schema](#nestedatt--rule))

<a id="nestedatt--rule"></a>
### Nested Schema for `rule`

Optional:

- `action` (String) Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  jump  |  Jump to another chain  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the                       last chain  |
|  drop  |  Drop matching entries  |
|  queue  |  Enqueue packet to userspace  |
- `connection_mark` (String) Connection mark

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection-mark to match  |
- `connection_status` (Attributes) Connection status (see [below for nested schema](#nestedatt--rule--connection_status))
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--rule--destination))
- `disable` (String) Option to disable firewall rule
- `dscp` (String) DSCP value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value to match  |
|  <start-end>  |  DSCP range to match  |
- `dscp_exclude` (String) DSCP value not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP value not to match  |
|  <start-end>  |  DSCP range not to match  |
- `fragment` (Attributes) IP fragment match (see [below for nested schema](#nestedatt--rule--fragment))
- `hop_limit` (Attributes) Hop limit (see [below for nested schema](#nestedatt--rule--hop_limit))
- `icmpv6` (Attributes) ICMPv6 type and code information (see [below for nested schema](#nestedatt--rule--icmpv6))
- `inbound_interface` (Attributes) Match inbound-interface (see [below for nested schema](#nestedatt--rule--inbound_interface))
- `ipsec` (Attributes) Inbound IPsec packets (see [below for nested schema](#nestedatt--rule--ipsec))
- `jump_target` (String) Set jump target. Action jump must be defined to use this setting
- `limit` (Attributes) Rate limit using a token bucket filter (see [below for nested schema](#nestedatt--rule--limit))
- `log` (String) Option to log packets matching rule

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable log  |
|  disable  |  Disable log  |
- `log_level` (String) Set log-level. Log must be enable.

|  Format  |  Description  |
|----------|---------------|
|  emerg  |  Emerg log level  |
|  alert  |  Alert log level  |
|  crit  |  Critical log level  |
|  err  |  Error log level  |
|  warn  |  Warning log level  |
|  notice  |  Notice log level  |
|  info  |  Info log level  |
|  debug  |  Debug log level  |
- `outbound_interface` (Attributes) Match outbound-interface (see [below for nested schema](#nestedatt--rule--outbound_interface))
- `packet_length` (String) Payload size in bytes, including header and data to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length to match  |
|  <start-end>  |  Packet length range to match  |
- `packet_length_exclude` (String) Payload size in bytes, including header and data not to match

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Packet length not to match  |
|  <start-end>  |  Packet length range not to match  |
- `packet_type` (String) Packet type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Match broadcast packet type  |
|  host  |  Match host packet type, addressed to local host  |
|  multicast  |  Match multicast packet type  |
|  other  |  Match packet addressed to another host  |
- `protocol` (String) Protocol to match (protocol name, number, or "all")

|  Format  |  Description  |
|----------|---------------|
|  all  |  All IP protocols  |
|  tcp_udp  |  Both TCP and UDP  |
|  u32:0-255  |  IP protocol number  |
|  <protocol>  |  IP protocol name  |
|  !<protocol>  |  IP protocol name  |
- `queue` (String) Queue target to use. Action queue must be defined to use this setting

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Queue target  |
- `queue_options` (String) Options used for queue target. Action queue must be defined to use this
                    setting

|  Format  |  Description  |
|----------|---------------|
|  bypass  |  Let packets go through if userspace application cannot back off  |
|  fanout  |  Distribute packets between several queues  |
- `recent` (Attributes) Parameters for matching recently seen sources (see [below for nested schema](#nestedatt--rule--recent))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--rule--source))
- `state` (Attributes) Session state (see [below for nested schema](#nestedatt--rule--state))
- `tcp` (Attributes) TCP flags to match (see [below for nested schema](#nestedatt--rule--tcp))
- `time` (Attributes) Time to match rule (see [below for nested schema](#nestedatt--rule--time))

<a id="nestedatt--rule--connection_status"></a>
### Nested Schema for `rule.connection_status`

Optional:

- `nat` (String) NAT connection status

|  Format  |  Description  |
|----------|---------------|
|  destination  |  Match connections that are subject to destination NAT  |
|  source  |  Match connections that are subject to source NAT  |


<a id="nestedatt--rule--destination"></a>
### Nested Schema for `rule.destination`

Optional:

- `address` (String) IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP address to match  |
|  ipv6net  |  Subnet to match  |
|  ipv6range  |  IP range to match  |
|  !ipv6  |  Match everything except the specified address  |
|  !ipv6net  |  Match everything except the specified prefix  |
|  !ipv6range  |  Match everything except the specified range  |
- `address_mask` (String) IP mask

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP mask to apply  |
- `fqdn` (String) Fully qualified domain name

|  Format  |  Description  |
|----------|---------------|
|  <fqdn>  |  Fully qualified domain name  |
- `geoip` (Attributes) GeoIP options - Data provided by DB-IP.com (see [below for nested schema](#nestedatt--rule--destination--geoip))
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--rule--destination--group))
- `mac_address` (String) MAC address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
|  !macaddr  |  Match everything except the specified MAC address  |
- `port` (String) Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n Multiple destination ports can be specified as a                           comma-separated list.\n For example: 'telnet,http,123,1001-1005'  |

<a id="nestedatt--rule--destination--geoip"></a>
### Nested Schema for `rule.destination.geoip`

Optional:

- `country_code` (String) GeoIP country code

|  Format  |  Description  |
|----------|---------------|
|  <country>  |  Country code (2 characters)  |
- `inverse_match` (String) Inverse match of country-codes


<a id="nestedatt--rule--destination--group"></a>
### Nested Schema for `rule.destination.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



<a id="nestedatt--rule--fragment"></a>
### Nested Schema for `rule.fragment`

Optional:

- `match_frag` (String) Second and further fragments of fragmented packets
- `match_non_frag` (String) Head fragments or unfragmented packets


<a id="nestedatt--rule--hop_limit"></a>
### Nested Schema for `rule.hop_limit`

Optional:

- `eq` (String) Match on equal value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Equal to value  |
- `gt` (String) Match on greater then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Greater then value  |
- `lt` (String) Match on less then value

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  Less then value  |


<a id="nestedatt--rule--icmpv6"></a>
### Nested Schema for `rule.icmpv6`

Optional:

- `code` (String) ICMPv6 code

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMPv6 code (0-255)  |
- `type` (String) ICMPv6 type

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMPv6 type (0-255)  |
- `type_name` (String) ICMPv6 type-name

|  Format  |  Description  |
|----------|---------------|
|  destination-unreachable  |  ICMPv6 type 1: destination-unreachable  |
|  packet-too-big  |  ICMPv6 type 2: packet-too-big  |
|  time-exceeded  |  ICMPv6 type 3: time-exceeded  |
|  echo-request  |  ICMPv6 type 128: echo-request  |
|  echo-reply  |  ICMPv6 type 129: echo-reply  |
|  mld-listener-query  |  ICMPv6 type 130: mld-listener-query  |
|  mld-listener-report  |  ICMPv6 type 131: mld-listener-report  |
|  mld-listener-reduction  |  ICMPv6 type 132: mld-listener-reduction  |
|  nd-router-solicit  |  ICMPv6 type 133: nd-router-solicit  |
|  nd-router-advert  |  ICMPv6 type 134: nd-router-advert  |
|  nd-neighbor-solicit  |  ICMPv6 type 135: nd-neighbor-solicit  |
|  nd-neighbor-advert  |  ICMPv6 type 136: nd-neighbor-advert  |
|  nd-redirect  |  ICMPv6 type 137: nd-redirect  |
|  parameter-problem  |  ICMPv6 type 4: parameter-problem  |
|  router-renumbering  |  ICMPv6 type 138: router-renumbering  |
|  ind-neighbor-solicit  |  ICMPv6 type 141: ind-neighbor-solicit  |
|  ind-neighbor-advert  |  ICMPv6 type 142: ind-neighbor-advert  |
|  mld2-listener-report  |  ICMPv6 type 143: mld2-listener-report  |


<a id="nestedatt--rule--inbound_interface"></a>
### Nested Schema for `rule.inbound_interface`

Optional:

- `interface_group` (String) Match interface-group
- `interface_name` (String) Match interface


<a id="nestedatt--rule--ipsec"></a>
### Nested Schema for `rule.ipsec`

Optional:

- `match_ipsec` (String) Inbound IPsec packets
- `match_none` (String) Inbound non-IPsec packets


<a id="nestedatt--rule--limit"></a>
### Nested Schema for `rule.limit`

Optional:

- `burst` (String) Maximum number of packets to allow in excess of rate

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Maximum number of packets to allow in excess of rate  |
- `rate` (String) Maximum average matching rate

|  Format  |  Description  |
|----------|---------------|
|  txt  |  integer/unit (Example: 5/minute)  |


<a id="nestedatt--rule--outbound_interface"></a>
### Nested Schema for `rule.outbound_interface`

Optional:

- `interface_group` (String) Match interface-group
- `interface_name` (String) Match interface


<a id="nestedatt--rule--recent"></a>
### Nested Schema for `rule.recent`

Optional:

- `count` (String) Source addresses seen more than N times

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Source addresses seen more than N times  |
- `time` (String) Source addresses seen in the last second/minute/hour

|  Format  |  Description  |
|----------|---------------|
|  second  |  Source addresses seen COUNT times in the last second  |
|  minute  |  Source addresses seen COUNT times in the last minute  |
|  hour  |  Source addresses seen COUNT times in the last hour  |


<a id="nestedatt--rule--source"></a>
### Nested Schema for `rule.source`

Optional:

- `address` (String) IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP address to match  |
|  ipv6net  |  Subnet to match  |
|  ipv6range  |  IP range to match  |
|  !ipv6  |  Match everything except the specified address  |
|  !ipv6net  |  Match everything except the specified prefix  |
|  !ipv6range  |  Match everything except the specified range  |
- `address_mask` (String) IP mask

|  Format  |  Description  |
|----------|---------------|
|  ipv6  |  IP mask to apply  |
- `fqdn` (String) Fully qualified domain name

|  Format  |  Description  |
|----------|---------------|
|  <fqdn>  |  Fully qualified domain name  |
- `geoip` (Attributes) GeoIP options - Data provided by DB-IP.com (see [below for nested schema](#nestedatt--rule--source--geoip))
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--rule--source--group))
- `mac_address` (String) MAC address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  MAC address to match  |
|  !macaddr  |  Match everything except the specified MAC address  |
- `port` (String) Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n Multiple destination ports can be specified as a                           comma-separated list.\n For example: 'telnet,http,123,1001-1005'  |

<a id="nestedatt--rule--source--geoip"></a>
### Nested Schema for `rule.source.geoip`

Optional:

- `country_code` (String) GeoIP country code

|  Format  |  Description  |
|----------|---------------|
|  <country>  |  Country code (2 characters)  |
- `inverse_match` (String) Inverse match of country-codes


<a id="nestedatt--rule--source--group"></a>
### Nested Schema for `rule.source.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



<a id="nestedatt--rule--state"></a>
### Nested Schema for `rule.state`

Optional:

- `established` (String) Established state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |
- `invalid` (String) Invalid state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |
- `new` (String) New state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |
- `related` (String) Related state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |


<a id="nestedatt--rule--tcp"></a>
### Nested Schema for `rule.tcp`

Optional:

- `flags` (Attributes) TCP flags to match (see [below for nested schema](#nestedatt--rule--tcp--flags))
- `mss` (String) Maximum segment size (MSS)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16384  |  Maximum segment size  |
|  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |

<a id="nestedatt--rule--tcp--flags"></a>
### Nested Schema for `rule.tcp.flags`

Optional:

- `ack` (String) Acknowledge flag
- `cwr` (String) Congestion Window Reduced flag
- `ecn` (String) Explicit Congestion Notification flag
- `fin` (String) Finish flag
- `not` (Attributes) Match flags not set (see [below for nested schema](#nestedatt--rule--tcp--flags--not))
- `psh` (String) Push flag
- `rst` (String) Reset flag
- `syn` (String) Synchronise flag
- `urg` (String) Urgent flag

<a id="nestedatt--rule--tcp--flags--not"></a>
### Nested Schema for `rule.tcp.flags.urg`

Optional:

- `ack` (String) Acknowledge flag
- `cwr` (String) Congestion Window Reduced flag
- `ecn` (String) Explicit Congestion Notification flag
- `fin` (String) Finish flag
- `psh` (String) Push flag
- `rst` (String) Reset flag
- `syn` (String) Synchronise flag
- `urg` (String) Urgent flag




<a id="nestedatt--rule--time"></a>
### Nested Schema for `rule.time`

Optional:

- `startdate` (String) Date to start matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter date using following notation - YYYY-MM-DD  |
- `starttime` (String) Time of day to start matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter time using using 24 hour notation - hh:mm:ss  |
- `stopdate` (String) Date to stop matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter date using following notation - YYYY-MM-DD  |
- `stoptime` (String) Time of day to stop matching rule

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Enter time using using 24 hour notation - hh:mm:ss  |
- `weekdays` (String) Comma separated weekdays to match rule on

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday,                           Saturday, Sunday)  |
|  u32:0-6  |  Day number (0 = Sunday ... 6 = Saturday)  |
