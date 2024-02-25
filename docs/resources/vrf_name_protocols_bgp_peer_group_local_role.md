---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vrf_name_protocols_bgp_peer_group_local_role Resource - vyos"
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
  Border Gateway Protocol (BGP)

  <br>
  &darr;
  <br>
  Name of peer-group

  <br>
  &darr;
  <br>
  <b>
  Local role for BGP neighbor (RFC9234)
  </b>
  </div>
---

# vyos_vrf_name_protocols_bgp_peer_group_local_role (Resource)

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
Border Gateway Protocol (BGP)

<br>
&darr;
<br>
Name of peer-group

<br>
&darr;
<br>
<b>
Local role for BGP neighbor (RFC9234)
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `local_role_id` (String) Local role for BGP neighbor (RFC9234)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  customer  &emsp; |  Using Transit  |
    |  peer  &emsp; |  Public/Private Peering  |
    |  provider  &emsp; |  Providing Transit  |
    |  rs-client  &emsp; |  RS Client  |
    |  rs-server  &emsp; |  Route Server  |
- `name_id` (String) Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |
- `peer_group_id` (String) Name of peer-group

### Optional

- `strict` (Boolean) Neighbor must send this exact capability, otherwise a role missmatch notification will be sent

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).