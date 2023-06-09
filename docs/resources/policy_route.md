---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_policy_route Resource - vyos"
subcategory: ""
description: |-
  Policy route rule set name for IPv4
---

# vyos_policy_route (Resource)

Policy route rule set name for IPv4



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Policy route rule set name for IPv4

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `enable_default_log` (String) Log packets hitting default-action
- `interface` (String) Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |
- `rule` (Attributes Map) Policy rule number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999999  |  Number of policy rule  | (see [below for nested schema](#nestedatt--rule))

<a id="nestedatt--rule"></a>
### Nested Schema for `rule`

Optional:

- `action` (String) Rule action

|  Format  |  Description  |
|----------|---------------|
|  accept  |  Accept matching entries  |
|  reject  |  Reject matching entries  |
|  return  |  Return from the current chain and continue at the next rule of the last chain  |
|  drop  |  Drop matching entries  |
- `connection_mark` (String) Connection mark

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection-mark to match  |
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
- `icmp` (Attributes) ICMP type and code information (see [below for nested schema](#nestedatt--rule--icmp))
- `ipsec` (Attributes) Inbound IPsec packets (see [below for nested schema](#nestedatt--rule--ipsec))
- `limit` (Attributes) Rate limit using a token bucket filter (see [below for nested schema](#nestedatt--rule--limit))
- `log` (String) Option to log packets matching rule

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable log  |
|  disable  |  Disable log  |
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
|  0-255  |  IP protocol number  |
|  !<protocol>  |  IP protocol number  |
- `recent` (Attributes) Parameters for matching recently seen sources (see [below for nested schema](#nestedatt--rule--recent))
- `set` (Attributes) Packet modifications (see [below for nested schema](#nestedatt--rule--set))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--rule--source))
- `state` (Attributes) Session state (see [below for nested schema](#nestedatt--rule--state))
- `tcp` (Attributes) TCP flags to match (see [below for nested schema](#nestedatt--rule--tcp))
- `time` (Attributes) Time to match rule (see [below for nested schema](#nestedatt--rule--time))
- `ttl` (Attributes) Time to live limit (see [below for nested schema](#nestedatt--rule--ttl))

<a id="nestedatt--rule--destination"></a>
### Nested Schema for `rule.destination`

Optional:

- `address` (String) IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--rule--destination--group))
- `port` (String) Port

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Named port (any name in /etc/services, e.g., http)  |
|  u32:1-65535  |  Numbered port  |
|  <start-end>  |  Numbered port range (e.g. 1001-1005)  |
|     |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

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


<a id="nestedatt--rule--icmp"></a>
### Nested Schema for `rule.icmp`

Optional:

- `code` (String) ICMP code (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP code (0-255)  |
- `type` (String) ICMP type (0-255)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  ICMP type (0-255)  |
- `type_name` (String) ICMP type-name

|  Format  |  Description  |
|----------|---------------|
|  echo-reply  |  ICMP type 0: echo-reply  |
|  destination-unreachable  |  ICMP type 3: destination-unreachable  |
|  source-quench  |  ICMP type 4: source-quench  |
|  redirect  |  ICMP type 5: redirect  |
|  echo-request  |  ICMP type 8: echo-request  |
|  router-advertisement  |  ICMP type 9: router-advertisement  |
|  router-solicitation  |  ICMP type 10: router-solicitation  |
|  time-exceeded  |  ICMP type 11: time-exceeded  |
|  parameter-problem  |  ICMP type 12: parameter-problem  |
|  timestamp-request  |  ICMP type 13: timestamp-request  |
|  timestamp-reply  |  ICMP type 14: timestamp-reply  |
|  info-request  |  ICMP type 15: info-request  |
|  info-reply  |  ICMP type 16: info-reply  |
|  address-mask-request  |  ICMP type 17: address-mask-request  |
|  address-mask-reply  |  ICMP type 18: address-mask-reply  |


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
|  u32:0-4294967295  |  Maximum average matching rate  |


<a id="nestedatt--rule--recent"></a>
### Nested Schema for `rule.recent`

Optional:

- `count` (String) Source addresses seen more than N times

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Source addresses seen more than N times  |
- `time` (String) Source addresses seen in the last N seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Source addresses seen in the last N seconds  |


<a id="nestedatt--rule--set"></a>
### Nested Schema for `rule.set`

Optional:

- `connection_mark` (String) Connection marking

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  Connection marking  |
- `dscp` (String) Packet Differentiated Services Codepoint (DSCP)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-63  |  DSCP number  |
- `mark` (String) Packet marking

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Packet marking  |
- `table` (String) Routing table to forward packet with

|  Format  |  Description  |
|----------|---------------|
|  u32:1-200  |  Table number  |
|  main  |  Main table  |
- `tcp_mss` (String) TCP Maximum Segment Size

|  Format  |  Description  |
|----------|---------------|
|  u32:500-1460  |  Explicitly set TCP MSS value  |


<a id="nestedatt--rule--source"></a>
### Nested Schema for `rule.source`

Optional:

- `address` (String) IP address, subnet, or range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |
|  ipv4range  |  IPv4 address range to match  |
|  !ipv4  |  Match everything except the specified address  |
|  !ipv4net  |  Match everything except the specified prefix  |
|  !ipv4range  |  Match everything except the specified range  |
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
|     |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

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

- `monthdays` (String) Monthdays to match rule on
- `startdate` (String) Date to start matching rule
- `starttime` (String) Time of day to start matching rule
- `stopdate` (String) Date to stop matching rule
- `stoptime` (String) Time of day to stop matching rule
- `utc` (String) Interpret times for startdate, stopdate, starttime and stoptime to be UTC
- `weekdays` (String) Weekdays to match rule on


<a id="nestedatt--rule--ttl"></a>
### Nested Schema for `rule.ttl`

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
