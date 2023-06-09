---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_qos_policy_fair_queue Resource - vyos"
subcategory: ""
description: |-
  Quality of Service (QoS)
  Service Policy definitions
  Stochastic Fairness Queueing
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Policy name  |
---

# vyos_qos_policy_fair_queue (Resource)

Quality of Service (QoS)

Service Policy definitions

Stochastic Fairness Queueing

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Stochastic Fairness Queueing

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Policy name  |

### Optional

- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `hash_interval` (String) Interval in seconds for queue algorithm perturbation

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  No perturbation  |
|  u32:1-127  |  Interval in seconds for queue algorithm perturbation (advised: 10)  |
- `queue_limit` (String) Upper limit of the SFQ

|  Format  |  Description  |
|----------|---------------|
|  u32:1-127  |  Queue size in packets  |
