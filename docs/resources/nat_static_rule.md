---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_nat_static_rule Resource - vyos"
subcategory: ""
description: |-
  Network Address Translation (NAT) parameters
  Static NAT (one-to-one)
  Rule number for NAT
---

# vyos_nat_static_rule (Resource)

Network Address Translation (NAT) parameters

Static NAT (one-to-one)

Rule number for NAT



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Rule number for NAT

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `destination` (Attributes) NAT destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Inbound interface of NAT traffic
- `translation` (Attributes) Translation address or prefix (see [below for nested schema](#nestedatt--translation))

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |


<a id="nestedatt--translation"></a>
### Nested Schema for `translation`

Optional:

- `address` (String) IP address, prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to match  |
|  ipv4net  |  IPv4 prefix to match  |