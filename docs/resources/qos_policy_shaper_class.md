---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_qos_policy_shaper_class Resource - vyos"
subcategory: "qos"
description: |-
  <div style="text-align: center">
  Quality of Service (QoS)

  <br>
  &darr;
  <br>
  Service Policy definitions

  <br>
  &darr;
  <br>
  Traffic shaping based policy (Hierarchy Token Bucket)

  <br>
  &darr;
  <br>
  <b>
  Class ID
  </b>
  </div>
---

# vyos_qos_policy_shaper_class (Resource)

<div style="text-align: center">
Quality of Service (QoS)

<br>
&darr;
<br>
Service Policy definitions

<br>
&darr;
<br>
Traffic shaping based policy (Hierarchy Token Bucket)

<br>
&darr;
<br>
<b>
Class ID
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `class_id` (String) Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-4095  &emsp; |  Class Identifier  |
- `shaper_id` (String) Traffic shaping based policy (Hierarchy Token Bucket)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

### Optional

- `bandwidth` (String) Available bandwidth for this policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  auto  &emsp; |  Bandwidth matches interface speed  |
    |  <number>  &emsp; |  Bits per second  |
    |  <number>bit  &emsp; |  Bits per second  |
    |  <number>kbit  &emsp; |  Kilobits per second  |
    |  <number>mbit  &emsp; |  Megabits per second  |
    |  <number>gbit  &emsp; |  Gigabits per second  |
    |  <number>tbit  &emsp; |  Terabits per second  |
    |  <number>%%  &emsp; |  Percentage of interface link speed  |
- `burst` (String) Burst size for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Bytes  |
    |  <number><suffix>  &emsp; |  Bytes with scaling suffix (kb, mb, gb)  |
- `ceiling` (String) Bandwidth limit for this class

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Rate in kbit (kilobit per second)  |
    |  <number>%%  &emsp; |  Percentage of overall rate  |
    |  <number>bit  &emsp; |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
    |  <number>ibit  &emsp; |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  &emsp; |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps  &emsp; |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
- `codel_quantum` (Number) Deficit in the fair queuing algorithm

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-1048576  &emsp; |  Number of bytes used as 'deficit'  |
- `description` (String) Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |
- `flows` (Number) Number of flows into which the incoming packets are classified

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65536  &emsp; |  Number of flows  |
- `interval` (Number) Interval used to measure the delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Interval in milliseconds  |
- `priority` (Number) Priority for rule evaluation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-20  &emsp; |  Priority for match rule evaluation  |
- `queue_limit` (Number) Maximum queue size

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Queue size in packets  |
- `queue_type` (String) Queue type for default traffic

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  drop-tail  &emsp; |  First-In-First-Out (FIFO)  |
    |  fair-queue  &emsp; |  Stochastic Fair Queue (SFQ)  |
    |  fq-codel  &emsp; |  Fair Queue Codel  |
    |  priority  &emsp; |  Priority queuing  |
    |  random-detect  &emsp; |  Random Early Detection (RED)  |
- `set_dscp` (String) Change the Differentiated Services (DiffServ) field in the IP header

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  Priority order for bandwidth pool  |
    |  default  &emsp; |  match DSCP (000000)  |
    |  reliability  &emsp; |  match DSCP (000001)  |
    |  throughput  &emsp; |  match DSCP (000010)  |
    |  lowdelay  &emsp; |  match DSCP (000100)  |
    |  priority  &emsp; |  match DSCP (001000)  |
    |  immediate  &emsp; |  match DSCP (010000)  |
    |  flash  &emsp; |  match DSCP (011000)  |
    |  flash-override  &emsp; |  match DSCP (100000)  |
    |  critical  &emsp; |  match DSCP (101000)  |
    |  internet  &emsp; |  match DSCP (110000)  |
    |  network  &emsp; |  match DSCP (111000)  |
    |  AF11  &emsp; |  High-throughput data  |
    |  AF12  &emsp; |  High-throughput data  |
    |  AF13  &emsp; |  High-throughput data  |
    |  AF21  &emsp; |  Low-latency data  |
    |  AF22  &emsp; |  Low-latency data  |
    |  AF23  &emsp; |  Low-latency data  |
    |  AF31  &emsp; |  Multimedia streaming  |
    |  AF32  &emsp; |  Multimedia streaming  |
    |  AF33  &emsp; |  Multimedia streaming  |
    |  AF41  &emsp; |  Multimedia conferencing  |
    |  AF42  &emsp; |  Multimedia conferencing  |
    |  AF43  &emsp; |  Multimedia conferencing  |
    |  CS1  &emsp; |  Low-priority data  |
    |  CS2  &emsp; |  OAM  |
    |  CS3  &emsp; |  Broadcast video  |
    |  CS4  &emsp; |  Real-time interactive  |
    |  CS5  &emsp; |  Signaling  |
    |  CS6  &emsp; |  Network control  |
    |  CS7  &emsp; |    |
    |  EF  &emsp; |  Expedited Forwarding  |
- `target` (Number) Acceptable minimum standing/persistent queue delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Queue delay in milliseconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).