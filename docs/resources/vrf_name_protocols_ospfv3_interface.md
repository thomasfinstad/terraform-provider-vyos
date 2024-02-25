---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vrf_name_protocols_ospfv3_interface Resource - vyos"
subcategory: "vrf"
description: |-
  <div style="text-align: center">
  Virtual Routing and Forwarding

  <br>
  &darr;
  <br>
  Virtual Routing and Forwarding instance

  <br>
  &darr;
  <br>
  Routing protocol parameters

  <br>
  &darr;
  <br>
  Open Shortest Path First (OSPF) for IPv6

  <br>
  &darr;
  <br>
  <b>
  Enable routing on an IPv6 interface
  </b>
  </div>
---

# vyos_vrf_name_protocols_ospfv3_interface (Resource)

<div style="text-align: center">
Virtual Routing and Forwarding

<br>
&darr;
<br>
Virtual Routing and Forwarding instance

<br>
&darr;
<br>
Routing protocol parameters

<br>
&darr;
<br>
Open Shortest Path First (OSPF) for IPv6

<br>
&darr;
<br>
<b>
Enable routing on an IPv6 interface
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Enable routing on an IPv6 interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface used for routing information exchange  |
- `name_id` (String) Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

### Optional

- `area` (String) Enable OSPF on this interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area ID as decimal notation  |
    |  ipv4  &emsp; |  OSPF area ID in IP address notation  |
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `cost` (Number) Interface cost

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  OSPF interface cost  |
- `dead_interval` (Number) Interval after which a neighbor is declared dead

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Neighbor dead interval (seconds)  |
- `hello_interval` (Number) Interval between hello packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Hello interval (seconds)  |
- `ifmtu` (Number) Interface MTU

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Interface MTU  |
- `instance_id` (Number) Instance ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  Instance Id  |
- `mtu_ignore` (Boolean) Disable Maximum Transmission Unit (MTU) mismatch detection
- `network` (String) Network type

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  broadcast  &emsp; |  Broadcast network type  |
    |  point-to-point  &emsp; |  Point-to-point network type  |
- `passive` (Boolean) Configure passive mode for interface
- `priority` (Number) Router priority

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-255  &emsp; |  OSPF router priority cost  |
- `retransmit_interval` (Number) Interval between retransmitting lost link state advertisements

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Retransmit interval (seconds)  |
- `transmit_delay` (Number) Link state transmit delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Link state transmit delay (seconds)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  BFD profile name  |