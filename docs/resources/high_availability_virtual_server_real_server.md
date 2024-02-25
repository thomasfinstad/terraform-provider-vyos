---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_high_availability_virtual_server_real_server Resource - vyos"
subcategory: "high"
description: |-
  <div style="text-align: center">
  High availability settings

  <br>
  &darr;
  <br>
  Load-balancing virtual server alias

  <br>
  &darr;
  <br>
  <b>
  Real server address
  </b>
  </div>
---

# vyos_high_availability_virtual_server_real_server (Resource)

<div style="text-align: center">
High availability settings

<br>
&darr;
<br>
Load-balancing virtual server alias

<br>
&darr;
<br>
<b>
Real server address
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `real_server_id` (String) Real server address
- `virtual_server_id` (String) Load-balancing virtual server alias

### Optional

- `connection_timeout` (Number) Server connection timeout

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-86400  &emsp; |  Connection timeout to remote server  |
- `health_check` (Attributes) Health check script (see [below for nested schema](#nestedatt--health_check))
- `port` (Number) Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Numeric IP port  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

<a id="nestedatt--health_check"></a>
### Nested Schema for `health_check`

Optional:

- `script` (String) Health check script file