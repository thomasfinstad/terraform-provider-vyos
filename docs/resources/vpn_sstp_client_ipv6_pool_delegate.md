---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn_sstp_client_ipv6_pool_delegate Resource - vyos"
subcategory: ""
description: |-
  Secure Socket Tunneling Protocol (SSTP) server
  Pool of client IPv6 addresses
  Subnet used to delegate prefix through DHCPv6-PD (RFC3633)
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv6net  |  IPv6 address and prefix length  |
---

# vyos_vpn_sstp_client_ipv6_pool_delegate (Resource)

Secure Socket Tunneling Protocol (SSTP) server

Pool of client IPv6 addresses

Subnet used to delegate prefix through DHCPv6-PD (RFC3633)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Subnet used to delegate prefix through DHCPv6-PD (RFC3633)

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

### Optional

- `delegation_prefix` (String) Prefix length delegated to client

|  Format  |  Description  |
|----------|---------------|
|  u32:32-64  |  Delegated prefix length  |