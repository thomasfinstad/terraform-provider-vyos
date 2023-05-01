---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_policy_local_route_rule Resource - vyos"
subcategory: ""
description: |-
  IPv4 policy route of local traffic
  Policy local-route rule set number
  |  Format  |  Description  |
  |----------|---------------|
  |  u32:1-32765  |  Local-route rule number (1-32765)  |
---

# vyos_policy_local_route_rule (Resource)

IPv4 policy route of local traffic

Policy local-route rule set number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32765  |  Local-route rule number (1-32765)  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Policy local-route rule set number

|  Format  |  Description  |
|----------|---------------|
|  u32:1-32765  |  Local-route rule number (1-32765)  |

### Optional

- `destination` (String) Destination address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to match against  |
|  ipv4net  |  Prefix to match against  |
- `fwmark` (String) Match fwmark value

|  Format  |  Description  |
|----------|---------------|
|  u32:1-2147483647  |  Address to match against  |
- `inbound_interface` (String) Inbound Interface
- `set` (Attributes) Packet modifications (see [below for nested schema](#nestedatt--set))
- `source` (String) Source address or prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Address to match against  |
|  ipv4net  |  Prefix to match against  |

<a id="nestedatt--set"></a>
### Nested Schema for `set`

Optional:

- `table` (String) Routing table to forward packet with

|  Format  |  Description  |
|----------|---------------|
|  u32:1-200  |  Table number  |