---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces_dummy Resource - vyos"
subcategory: ""
description: |-
  Dummy Interface
  |  Format  |  Description  |
  |----------|---------------|
  |  dumN  |  Dummy interface name  |
---

# vyos_interfaces_dummy (Resource)

Dummy Interface

|  Format  |  Description  |
|----------|---------------|
|  dumN  |  Dummy interface name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Dummy Interface

|  Format  |  Description  |
|----------|---------------|
|  dumN  |  Dummy interface name  |

### Optional

- `address` (String) IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `disable` (String) Administratively disable interface
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mtu` (String) Maximum Transmission Unit (MTU)

|  Format  |  Description  |
|----------|---------------|
|  u32:68-16000  |  Maximum Transmission Unit in byte  |
- `netns` (String) Network namespace name

|  Format  |  Description  |
|----------|---------------|
|  text  |  Network namespace name  |
- `redirect` (String) Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
- `vrf` (String) VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `disable_forwarding` (String) Disable IP forwarding on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |


<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `disable_forwarding` (String) Disable IP forwarding on this interface

<a id="nestedatt--ipv6--address"></a>
### Nested Schema for `ipv6.address`

Optional:

- `eui64` (String) Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
- `no_default_link_local` (String) Remove the default link-local address from the interface



<a id="nestedatt--mirror"></a>
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |