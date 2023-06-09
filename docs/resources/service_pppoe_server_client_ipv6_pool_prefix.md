---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_pppoe_server_client_ipv6_pool_prefix Resource - vyos"
subcategory: ""
description: |-
  Point to Point over Ethernet (PPPoE) Server
  Pool of client IPv6 addresses
  Pool of addresses used to assign to clients
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv6net  |  IPv6 address and prefix length  |
---

# vyos_service_pppoe_server_client_ipv6_pool_prefix (Resource)

Point to Point over Ethernet (PPPoE) Server

Pool of client IPv6 addresses

Pool of addresses used to assign to clients

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Pool of addresses used to assign to clients

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |

### Optional

- `mask` (String) Prefix length used for individual client

|  Format  |  Description  |
|----------|---------------|
|  u32:48-128  |  Client prefix length  |
