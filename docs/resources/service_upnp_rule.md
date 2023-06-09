---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_upnp_rule Resource - vyos"
subcategory: ""
description: |-
  Universal Plug and Play (UPnP) service
  UPnP Rule
  |  Format  |  Description  |
  |----------|---------------|
  |  u32:0-65535  |  Rule number  |
---

# vyos_service_upnp_rule (Resource)

Universal Plug and Play (UPnP) service

UPnP Rule

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Rule number  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) UPnP Rule

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Rule number  |

### Optional

- `action` (String) Actions against the rule (REQUIRE)
- `disable` (String) Disable instance
- `external_port_range` (String) Port range (REQUIRE)

|  Format  |  Description  |
|----------|---------------|
|  <port>  |  single port  |
|  <portN>-<portM>  |  Port range (use '-' as delimiter)  |
- `internal_port_range` (String) Port range (REQUIRE)

|  Format  |  Description  |
|----------|---------------|
|  <port>  |  single port  |
|  <portN>-<portM>  |  Port range (use '-' as delimiter)  |
- `ip` (String) The IP to which this rule applies (REQUIRE)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  The IPv4 address to which this rule applies  |
|  ipv4net  |  The IPv4 to which this rule applies  |
