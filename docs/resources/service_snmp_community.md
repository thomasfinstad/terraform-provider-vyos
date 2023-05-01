---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_snmp_community Resource - vyos"
subcategory: ""
description: |-
  Simple Network Management Protocol (SNMP)
  Community name
---

# vyos_service_snmp_community (Resource)

Simple Network Management Protocol (SNMP)

Community name



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Community name

### Optional

- `authorization` (String) Authorization type

|  Format  |  Description  |
|----------|---------------|
|  ro  |  Read-Only  |
|  rw  |  Read-Write  |
- `client` (String) IP address of SNMP client allowed to contact system
- `network` (String) Subnet of SNMP client(s) allowed to contact system

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IP address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |