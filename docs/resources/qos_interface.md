---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_qos_interface Resource - vyos"
subcategory: ""
description: |-
  Quality of Service (QoS)
  Interface to apply QoS policy
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Interface name  |
---

# vyos_qos_interface (Resource)

Quality of Service (QoS)

Interface to apply QoS policy

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Interface to apply QoS policy

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

### Optional

- `egress` (String) Interface egress traffic policy

|  Format  |  Description  |
|----------|---------------|
|  txt  |  QoS policy to use  |
- `ingress` (String) Interface ingress traffic policy

|  Format  |  Description  |
|----------|---------------|
|  txt  |  QoS policy to use  |
