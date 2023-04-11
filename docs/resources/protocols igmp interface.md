---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols igmp interface Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_protocols igmp interface (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `join` (Attributes Map) IGMP join multicast group

Format: ipv4
Multicast group address (see [below for nested schema](#nestedatt--join))
- `query_interval` (String) IGMP host query interval

Format: u32:1-1800
Query interval in seconds
- `query_max_response_time` (String) IGMP max query response time

Format: u32:10-250
Query response value in deci-seconds
- `version` (String) IGMP version

Format: 2
IGMP version 2
Format: 3
IGMP version 3

<a id="nestedatt--join"></a>
### Nested Schema for `join`

Optional:

- `source` (String) Source address

Format: ipv4
Source address