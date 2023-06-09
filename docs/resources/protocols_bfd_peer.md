---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bfd_peer Resource - vyos"
subcategory: ""
description: |-
  Bidirectional Forwarding Detection (BFD)
  Configures BFD peer to listen and talk to
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4  |  BFD peer IPv4 address  |
  |  ipv6  |  BFD peer IPv6 address  |
---

# vyos_protocols_bfd_peer (Resource)

Bidirectional Forwarding Detection (BFD)

Configures BFD peer to listen and talk to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  BFD peer IPv4 address  |
|  ipv6  |  BFD peer IPv6 address  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Configures BFD peer to listen and talk to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  BFD peer IPv4 address  |
|  ipv6  |  BFD peer IPv6 address  |

### Optional

- `echo_mode` (String) Enables the echo transmission mode
- `interval` (Attributes) Configure timer intervals (see [below for nested schema](#nestedatt--interval))
- `multihop` (String) Allow this BFD peer to not be directly connected
- `passive` (String) Do not attempt to start sessions
- `profile` (String) Use settings from BFD profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BFD profile name  |
- `shutdown` (String) Disable this peer
- `source` (Attributes) Bind listener to specified interface/address, mandatory for IPv6 (see [below for nested schema](#nestedatt--source))
- `vrf` (String) VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

<a id="nestedatt--interval"></a>
### Nested Schema for `interval`

Optional:

- `echo_interval` (String) Echo receive transmission interval

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  The minimal echo receive transmission interval that this system is capable of handling  |
- `multiplier` (String) Multiplier to determine packet loss

|  Format  |  Description  |
|----------|---------------|
|  u32:2-255  |  Remote transmission interval will be multiplied by this value  |
- `receive` (String) Minimum interval of receiving control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |
- `transmit` (String) Minimum interval of transmitting control packets

|  Format  |  Description  |
|----------|---------------|
|  u32:10-60000  |  Interval in milliseconds  |


<a id="nestedatt--source"></a>
### Nested Schema for `source`

Optional:

- `address` (String) Local address to bind our peer listener to

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Local IPv4 address used to connect to the peer  |
|  ipv6  |  Local IPv6 address used to connect to the peer  |
- `interface` (String) Interface to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |
