---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_address_family_ipv4_multicast_network Resource - vyos"
subcategory: ""
description: |-
  Border Gateway Protocol (BGP)
  BGP address-family parameters
  Multicast IPv4 BGP settings
  Import BGP network/prefix into multicast IPv4 RIB
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4net  |  Multicast IPv4 BGP network/prefix  |
---

# vyos_protocols_bgp_address_family_ipv4_multicast_network (Resource)

Border Gateway Protocol (BGP)

BGP address-family parameters

Multicast IPv4 BGP settings

Import BGP network/prefix into multicast IPv4 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Multicast IPv4 BGP network/prefix  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Import BGP network/prefix into multicast IPv4 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Multicast IPv4 BGP network/prefix  |

### Optional

- `backdoor` (String) Use BGP network/prefix as a backdoor route
- `route_map` (String) Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |