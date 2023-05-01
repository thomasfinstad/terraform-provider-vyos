---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_pppoe_server_interface Resource - vyos"
subcategory: ""
description: |-
  Point to Point over Ethernet (PPPoE) Server
  interface(s) to listen on
---

# vyos_service_pppoe_server_interface (Resource)

Point to Point over Ethernet (PPPoE) Server

interface(s) to listen on



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) interface(s) to listen on

### Optional

- `vlan` (String) VLAN monitor for automatic creation of VLAN interfaces

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4094  |  VLAN for automatic creation  |
|  start-end  |  VLAN range for automatic creation (e.g. 1-4094)  |