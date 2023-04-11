---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols nhrp tunnel Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_protocols nhrp tunnel (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cisco_authentication` (String) Pass phrase for cisco authentication

Format: txt
Pass phrase for cisco authentication
- `dynamic_map` (Attributes Map) Set an HUB tunnel address

Format: ipv4net
Set the IP address and prefix length (see [below for nested schema](#nestedatt--dynamic_map))
- `holding_time` (String) Holding time in seconds
- `map` (Attributes Map) Set an HUB tunnel address (see [below for nested schema](#nestedatt--map))
- `multicast` (String) Set multicast for NHRP
- `non_caching` (String) This can be used to reduce memory consumption on big NBMA subnets
- `redirect` (String) Enable sending of Cisco style NHRP Traffic Indication packets
- `shortcut` (String) Enable creation of shortcut routes. A received NHRP Traffic Indication will trigger the resolution and establishment of a shortcut route
- `shortcut_destination` (String) This instructs opennhrp to reply with authorative answers on NHRP Resolution Requests destined to addresses in this interface
- `shortcut_target` (Attributes Map) Defines an off-NBMA network prefix for which the GRE interface will act as a gateway (see [below for nested schema](#nestedatt--shortcut_target))

<a id="nestedatt--dynamic_map"></a>
### Nested Schema for `dynamic_map`

Optional:

- `nbma_domain_name` (String) Set HUB fqdn (nbma-address - fqdn)

Format: <fqdn>
Set the external HUB fqdn


<a id="nestedatt--map"></a>
### Nested Schema for `map`

Optional:

- `cisco` (String) If the statically mapped peer is running Cisco IOS, specify this
- `nbma_address` (String) Set HUB address (nbma-address - external hub address or fqdn)
- `register` (String) Specifies that Registration Request should be sent to this peer on startup


<a id="nestedatt--shortcut_target"></a>
### Nested Schema for `shortcut_target`

Optional:

- `holding_time` (String) Holding time in seconds