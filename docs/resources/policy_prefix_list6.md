---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_policy_prefix_list6 Resource - vyos"
subcategory: ""
description: |-
  Routing policy
  IPv6 prefix-list filter
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Name of IPv6 prefix-list  |
---

# vyos_policy_prefix_list6 (Resource)

Routing policy

IPv6 prefix-list filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) IPv6 prefix-list filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of IPv6 prefix-list  |

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `rule` (Attributes Map) Rule for this prefix-list6

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Prefix-list rule number  | (see [below for nested schema](#nestedatt--rule))

<a id="nestedatt--rule"></a>
### Nested Schema for `rule`

Optional:

- `action` (String) Action to take on entries matching this rule

|  Format  |  Description  |
|----------|---------------|
|  permit  |  Permit matching entries  |
|  deny  |  Deny matching entries  |
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `ge` (String) Prefix length to match a netmask greater than or equal to it

|  Format  |  Description  |
|----------|---------------|
|  u32:0-128  |  Netmask greater than length  |
- `le` (String) Prefix length to match a netmask less than or equal to it

|  Format  |  Description  |
|----------|---------------|
|  u32:0-128  |  Netmask less than length  |
- `prefix` (String) Prefix to match

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 prefix  |