---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system_flow_accounting_netflow_server Resource - vyos"
subcategory: ""
description: |-
  Flow accounting settings
  NetFlow settings
  NetFlow destination server
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4  |  IPv4 server to export NetFlow  |
  |  ipv6  |  IPv6 server to export NetFlow  |
---

# vyos_system_flow_accounting_netflow_server (Resource)

Flow accounting settings

NetFlow settings

NetFlow destination server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 server to export NetFlow  |
|  ipv6  |  IPv6 server to export NetFlow  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) NetFlow destination server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 server to export NetFlow  |
|  ipv6  |  IPv6 server to export NetFlow  |

### Optional

- `port` (String) NetFlow port number

|  Format  |  Description  |
|----------|---------------|
|  u32:1025-65535  |  NetFlow port number  |
