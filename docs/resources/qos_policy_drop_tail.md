---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_qos_policy_drop_tail Resource - vyos"
subcategory: ""
description: |-
  Quality of Service (QoS)
  Service Policy definitions
  Packet limited First In, First Out queue
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Policy name  |
---

# vyos_qos_policy_drop_tail (Resource)

Quality of Service (QoS)

Service Policy definitions

Packet limited First In, First Out queue

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Packet limited First In, First Out queue

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `queue_limit` (String) Maximum queue size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4294967295  |  Queue size in packets  |
