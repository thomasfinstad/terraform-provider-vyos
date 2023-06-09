---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_snmp_trap_target Resource - vyos"
subcategory: ""
description: |-
  Simple Network Management Protocol (SNMP)
  Address of trap target
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4  |  IPv4 address  |
  |  ipv6  |  IPv6 address  |
---

# vyos_service_snmp_trap_target (Resource)

Simple Network Management Protocol (SNMP)

Address of trap target

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
|  ipv6  |  IPv6 address  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Address of trap target

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
|  ipv6  |  IPv6 address  |

### Optional

- `community` (String) Community used when sending trap information
- `port` (String) Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
