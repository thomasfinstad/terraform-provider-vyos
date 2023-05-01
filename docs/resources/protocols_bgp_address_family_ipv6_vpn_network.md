---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_address_family_ipv6_vpn_network Resource - vyos"
subcategory: ""
description: |-
  Border Gateway Protocol (BGP)
  BGP address-family parameters
  Unicast VPN IPv6 BGP settings
  Import BGP network/prefix into unicast VPN IPv6 RIB
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv6net  |  Unicast VPN IPv6 BGP network/prefix  |
---

# vyos_protocols_bgp_address_family_ipv6_vpn_network (Resource)

Border Gateway Protocol (BGP)

BGP address-family parameters

Unicast VPN IPv6 BGP settings

Import BGP network/prefix into unicast VPN IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Unicast VPN IPv6 BGP network/prefix  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Import BGP network/prefix into unicast VPN IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Unicast VPN IPv6 BGP network/prefix  |

### Optional

- `label` (String) MPLS label value assigned to route

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1048575  |  MPLS label value  |
- `rd` (String) Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |