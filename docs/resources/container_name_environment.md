---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_container_name_environment Resource - vyos"
subcategory: "container"
description: |-
  <div style="text-align: center">
  Container applications

  <br>
  &darr;
  <br>
  Container name

  <br>
  &darr;
  <br>
  <b>
  Add custom environment variables
  </b>
  </div>
---

# vyos_container_name_environment (Resource)

<div style="text-align: center">
Container applications

<br>
&darr;
<br>
Container name

<br>
&darr;
<br>
<b>
Add custom environment variables
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Add custom environment variables
- `name_id` (String) Container name

### Optional

- `value` (String) Set environment option value

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Set environment option value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).