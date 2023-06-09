---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_ospf_area Resource - vyos"
subcategory: ""
description: |-
  Open Shortest Path First (OSPF)
  OSPF area settings
  |  Format  |  Description  |
  |----------|---------------|
  |  u32  |  OSPF area number in decimal notation  |
  |  ipv4  |  OSPF area number in dotted decimal notation  |
---

# vyos_protocols_ospf_area (Resource)

Open Shortest Path First (OSPF)

OSPF area settings

|  Format  |  Description  |
|----------|---------------|
|  u32  |  OSPF area number in decimal notation  |
|  ipv4  |  OSPF area number in dotted decimal notation  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) OSPF area settings

|  Format  |  Description  |
|----------|---------------|
|  u32  |  OSPF area number in decimal notation  |
|  ipv4  |  OSPF area number in dotted decimal notation  |

### Optional

- `area_type` (Attributes) Area type (see [below for nested schema](#nestedatt--area_type))
- `authentication` (String) OSPF area authentication type

|  Format  |  Description  |
|----------|---------------|
|  plaintext-password  |  Use plain-text authentication  |
|  md5  |  Use MD5 authentication  |
- `export_list` (String) Set the filter for networks announced to other areas

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access-list number  |
- `import_list` (String) Set the filter for networks from other areas announced

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Access-list number  |
- `network` (String) OSPF network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  OSPF network  |
- `range` (Attributes Map) Summarize routes matching a prefix (border routers only)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Area range prefix  | (see [below for nested schema](#nestedatt--range))
- `shortcut` (String) Area shortcut mode

|  Format  |  Description  |
|----------|---------------|
|  default  |  Set default  |
|  disable  |  Disable shortcutting mode  |
|  enable  |  Enable shortcutting mode  |
- `virtual_link` (Attributes Map) Virtual link

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  OSPF area in dotted decimal notation  | (see [below for nested schema](#nestedatt--virtual_link))

<a id="nestedatt--area_type"></a>
### Nested Schema for `area_type`

Optional:

- `normal` (String) Normal OSPF area
- `nssa` (Attributes) Not-So-Stubby OSPF area (see [below for nested schema](#nestedatt--area_type--nssa))
- `stub` (Attributes) Stub OSPF area (see [below for nested schema](#nestedatt--area_type--stub))

<a id="nestedatt--area_type--nssa"></a>
### Nested Schema for `area_type.nssa`

Optional:

- `default_cost` (String) Summary-default cost of an NSSA area

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Summary default cost  |
- `no_summary` (String) Do not inject inter-area routes into stub
- `translate` (String) Configure NSSA-ABR

|  Format  |  Description  |
|----------|---------------|
|  always  |  Always translate LSA types  |
|  candidate  |  Translate for election  |
|  never  |  Never translate LSA types  |


<a id="nestedatt--area_type--stub"></a>
### Nested Schema for `area_type.stub`

Optional:

- `default_cost` (String) Summary-default cost

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Summary default cost  |
- `no_summary` (String) Do not inject inter-area routes into the stub



<a id="nestedatt--range"></a>
### Nested Schema for `range`

Optional:

- `cost` (String) Metric for this range

|  Format  |  Description  |
|----------|---------------|
|  u32:0-16777215  |  Metric for this range  |
- `not_advertise` (String) Do not advertise this range
- `substitute` (String) Advertise area range as another prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Advertise area range as another prefix  |


<a id="nestedatt--virtual_link"></a>
### Nested Schema for `virtual_link`

Optional:

- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--virtual_link--authentication))
- `dead_interval` (String) Interval after which a neighbor is declared dead

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Neighbor dead interval (seconds)  |
- `hello_interval` (String) Interval between hello packets

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Hello interval (seconds)  |
- `retransmit_interval` (String) Interval between retransmitting lost link state advertisements

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Retransmit interval (seconds)  |
- `transmit_delay` (String) Link state transmit delay

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Link state transmit delay (seconds)  |

<a id="nestedatt--virtual_link--authentication"></a>
### Nested Schema for `virtual_link.authentication`

Optional:

- `md5` (Attributes) MD5 key id (see [below for nested schema](#nestedatt--virtual_link--authentication--md5))
- `plaintext_password` (String) Plain text password

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Plain text password (8 characters or less)  |

<a id="nestedatt--virtual_link--authentication--md5"></a>
### Nested Schema for `virtual_link.authentication.md5`

Optional:

- `key_id` (Attributes Map) MD5 key id

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  MD5 key id  | (see [below for nested schema](#nestedatt--virtual_link--authentication--md5--key_id))

<a id="nestedatt--virtual_link--authentication--md5--key_id"></a>
### Nested Schema for `virtual_link.authentication.md5.key_id`

Optional:

- `md5_key` (String) MD5 authentication type

|  Format  |  Description  |
|----------|---------------|
|  txt  |  MD5 Key (16 characters or less)  |
