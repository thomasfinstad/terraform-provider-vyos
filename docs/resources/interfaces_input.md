---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces_input Resource - vyos"
subcategory: ""
description: |-
  Input Functional Block (IFB) interface name
  |  Format  |  Description  |
  |----------|---------------|
  |  ifbN  |  Input interface name  |
---

# vyos_interfaces_input (Resource)

Input Functional Block (IFB) interface name

|  Format  |  Description  |
|----------|---------------|
|  ifbN  |  Input interface name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Input Functional Block (IFB) interface name

|  Format  |  Description  |
|----------|---------------|
|  ifbN  |  Input interface name  |

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `disable` (String) Administratively disable interface
- `redirect` (String) Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |