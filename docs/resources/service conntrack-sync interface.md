---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service conntrack-sync interface Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_service conntrack-sync interface (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `peer` (String) IP address of the peer to send the UDP conntrack info too. This disable multicast.

Format: ipv4
IP address to listen for incoming connections
- `port` (String) Port number used by connection

Format: u32:1-65535
Numeric IP port