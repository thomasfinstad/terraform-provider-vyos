---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces loopback Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_interfaces loopback (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (String) IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
- `description` (String) Description

Format: txt
Description
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `redirect` (String) Redirect incoming packet to destination

Format: txt
Destination interface name

<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `source_validation` (String) Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation


<a id="nestedatt--mirror"></a>
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

Format: txt
Destination interface name
- `ingress` (String) Mirror ingress traffic to destination interface

Format: txt
Destination interface name