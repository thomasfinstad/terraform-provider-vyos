---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_pim_interface Resource - vyos"
subcategory: ""
description: |-
  Protocol Independent Multicast (PIM)
  PIM interface
---

# vyos_protocols_pim_interface (Resource)

Protocol Independent Multicast (PIM)

PIM interface



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) PIM interface

### Optional

- `dr_priority` (String) Designated Router Election Priority

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Value of the new DR Priority  |
- `hello` (String) Hello Interval

|  Format  |  Description  |
|----------|---------------|
|  u32:1-180  |  Hello Interval in seconds  |
