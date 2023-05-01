---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_ospf_interface Resource - vyos"
subcategory: ""
description: |-
  Open Shortest Path First (OSPF)
  Interface configuration
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Interface name  |
---

# vyos_protocols_ospf_interface (Resource)

Open Shortest Path First (OSPF)

Interface configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Interface configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Interface name  |

### Optional

- `area` (String) Enable OSPF on this interface

|  Format  |  Description  |
|----------|---------------|
|  u32  |  OSPF area ID as decimal notation  |
|  ipv4  |  OSPF area ID in IP address notation  |
- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--authentication))
- `bandwidth` (String) Interface bandwidth (Mbit/s)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-100000  |  Bandwidth in Megabit/sec (for calculating OSPF cost)  |
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `cost` (String) Interface cost

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  OSPF interface cost  |
- `dead_interval` (String) Interval after which a neighbor is declared dead

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Neighbor dead interval (seconds)  |
- `hello_interval` (String) Interval between hello packets

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hello interval (seconds)  |
- `hello_multiplier` (String) Hello multiplier factor

|  Format  |  Description  |
|----------|---------------|
|  u32:1-10  |  Number of Hellos to send each second  |
- `mtu_ignore` (String) Disable Maximum Transmission Unit (MTU) mismatch detection
- `network` (String) Network type

|  Format  |  Description  |
|----------|---------------|
|  broadcast  |  Broadcast network type  |
|  non-broadcast  |  Non-broadcast network type  |
|  point-to-multipoint  |  Point-to-multipoint network type  |
|  point-to-point  |  Point-to-point network type  |
- `passive` (Attributes) Suppress routing updates on an interface (see [below for nested schema](#nestedatt--passive))
- `priority` (String) Router priority

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  OSPF router priority cost  |
- `retransmit_interval` (String) Interval between retransmitting lost link state advertisements

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Retransmit interval (seconds)  |
- `transmit_delay` (String) Link state transmit delay

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Link state transmit delay (seconds)  |

<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `md5` (Attributes) MD5 key id (see [below for nested schema](#nestedatt--authentication--md5))
- `plaintext_password` (String) Plain text password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Plain text password (8 characters or less)  |

<a id="nestedatt--authentication--md5"></a>
### Nested Schema for `authentication.md5`

Optional:

- `key_id` (Attributes Map) MD5 key id

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  MD5 key id  | (see [below for nested schema](#nestedatt--authentication--md5--key_id))

<a id="nestedatt--authentication--md5--key_id"></a>
### Nested Schema for `authentication.md5.key_id`

Optional:

- `md5_key` (String) MD5 authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  MD5 Key (16 characters or less)  |




<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

|  Format  |  Description  |
|----------|---------------|
|  txt  |  BFD profile name  |


<a id="nestedatt--passive"></a>
### Nested Schema for `passive`

Optional:

- `disable` (String) Disable instance