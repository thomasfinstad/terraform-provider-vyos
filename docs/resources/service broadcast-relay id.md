---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service broadcast-relay id Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_service broadcast-relay id (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (String) Set source IP of forwarded packets, otherwise original senders address is used

Format: ipv4
Optional source address for forwarded packets
- `description` (String) Description
- `disable` (String) Disable instance
- `interface` (String) Interface to use

Format: txt
Interface name
- `port` (String) Port number used by connection

Format: u32:1-65535
Numeric IP port