---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_policy_large_community_list Resource - vyos"
subcategory: ""
description: |-
  Routing policy
  Add a BGP large community list entry
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  BGP large-community-list name  |
---

# vyos_policy_large_community_list (Resource)

Routing policy

Add a BGP large community list entry

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BGP large-community-list name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Add a BGP large community list entry

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BGP large-community-list name  |

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `rule` (Attributes Map) Rule for this BGP extended community list

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Large community-list rule number  | (see [below for nested schema](#nestedatt--rule))

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
- `regex` (String) Regular expression to match against a large community list

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN:NN  |  BGP large-community-list filter  |
|  IP:NN:NN  |  BGP large-community-list filter (IPv4 address format)  |
