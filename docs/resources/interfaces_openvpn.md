---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces_openvpn Resource - vyos"
subcategory: ""
description: |-
  OpenVPN Tunnel Interface
  |  Format  |  Description  |
  |----------|---------------|
  |  vtunN  |  OpenVPN interface name  |
---

# vyos_interfaces_openvpn (Resource)

OpenVPN Tunnel Interface

|  Format  |  Description  |
|----------|---------------|
|  vtunN  |  OpenVPN interface name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) OpenVPN Tunnel Interface

|  Format  |  Description  |
|----------|---------------|
|  vtunN  |  OpenVPN interface name  |

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `device_type` (String) OpenVPN interface device-type

|  Format  |  Description  |
|----------|---------------|
|  tun  |  TUN device, required for OSI layer 3  |
|  tap  |  TAP device, required for OSI layer 2  |
- `disable` (String) Administratively disable interface
- `encryption` (Attributes) Data Encryption settings (see [below for nested schema](#nestedatt--encryption))
- `hash` (String) Hashing Algorithm

|  Format  |  Description  |
|----------|---------------|
|  md5  |  MD5 algorithm  |
|  sha1  |  SHA-1 algorithm  |
|  sha256  |  SHA-256 algorithm  |
|  sha384  |  SHA-384 algorithm  |
|  sha512  |  SHA-512 algorithm  |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `keep_alive` (Attributes) Keepalive helper options (see [below for nested schema](#nestedatt--keep_alive))
- `local_address` (Attributes Map) Local IP address of tunnel (IPv4 or IPv6) (see [below for nested schema](#nestedatt--local_address))
- `local_host` (String) Local IP address to accept connections (all if not set)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Local IPv4 address  |
|  ipv6  |  Local IPv6 address  |
- `local_port` (String) Local port number to accept connections

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mode` (String) OpenVPN mode of operation

|  Format  |  Description  |
|----------|---------------|
|  site-to-site  |  Site-to-site mode  |
|  client  |  Client in client-server mode  |
|  server  |  Server in client-server mode  |
- `openvpn_option` (String) Additional OpenVPN options. You must use the syntax of openvpn.conf in this text-field. Using this without proper knowledge may result in a crashed OpenVPN server. Check system log to look for errors.
- `persistent_tunnel` (String) Do not close and reopen interface (TUN/TAP device) on client restarts
- `protocol` (String) OpenVPN communication protocol

|  Format  |  Description  |
|----------|---------------|
|  udp  |  UDP  |
|  tcp-passive  |  TCP and accepts connections passively  |
|  tcp-active  |  TCP and initiates connections actively  |
- `redirect` (String) Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
- `remote_address` (String) IP address of remote end of tunnel

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Remote end IPv4 address  |
|  ipv6  |  Remote end IPv6 address  |
- `remote_host` (String) Remote host to connect to (dynamic if not set)

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of remote host  |
|  ipv6  |  IPv6 address of remote host  |
|  txt  |  Hostname of remote host  |
- `remote_port` (String) Remote port number to connect to

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
- `replace_default_route` (Attributes) OpenVPN tunnel to be used as the default route (see [below for nested schema](#nestedatt--replace_default_route))
- `server` (Attributes) Server-mode options (see [below for nested schema](#nestedatt--server))
- `shared_secret_key` (String) Secret key shared with remote end of tunnel
- `tls` (Attributes) Transport Layer Security (TLS) options (see [below for nested schema](#nestedatt--tls))
- `use_lzo_compression` (String) Use fast LZO compression on this TUN/TAP interface
- `vrf` (String) VRF instance name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  VRF instance name  |

<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Password  |
- `username` (String) Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username  |


<a id="nestedatt--encryption"></a>
### Nested Schema for `encryption`

Optional:

- `cipher` (String) Standard Data Encryption Algorithm

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  bf128  |  Blowfish algorithm with 128-bit key  |
|  bf256  |  Blowfish algorithm with 256-bit key  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |
- `ncp_ciphers` (String) Cipher negotiation list for use in server or client mode

|  Format  |  Description  |
|----------|---------------|
|  none  |  Disable encryption  |
|  des  |  DES algorithm  |
|  3des  |  DES algorithm with triple encryption  |
|  aes128  |  AES algorithm with 128-bit key CBC  |
|  aes128gcm  |  AES algorithm with 128-bit key GCM  |
|  aes192  |  AES algorithm with 192-bit key CBC  |
|  aes192gcm  |  AES algorithm with 192-bit key GCM  |
|  aes256  |  AES algorithm with 256-bit key CBC  |
|  aes256gcm  |  AES algorithm with 256-bit key GCM  |


<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `adjust_mss` (String) Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |
- `arp_cache_timeout` (String) ARP cache entry timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:1-86400  |  ARP cache entry timout in seconds  |
- `disable_arp_filter` (String) Disable ARP filter on this interface
- `disable_forwarding` (String) Disable IP forwarding on this interface
- `enable_arp_accept` (String) Enable ARP accept on this interface
- `enable_arp_announce` (String) Enable ARP announce on this interface
- `enable_arp_ignore` (String) Enable ARP ignore on this interface
- `enable_directed_broadcast` (String) Enable directed broadcast forwarding on this interface
- `enable_proxy_arp` (String) Enable proxy-arp on this interface
- `proxy_arp_pvlan` (String) Enable private VLAN proxy ARP on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |


<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `adjust_mss` (String) Adjust TCP MSS value

|  Format  |  Description  |
|----------|---------------|
|  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
|  u32:536-65535  |  TCP Maximum segment size in bytes  |
- `disable_forwarding` (String) Disable IP forwarding on this interface
- `dup_addr_detect_transmits` (String) Number of NS messages to send while performing DAD (default: 1)

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable Duplicate Address Dectection (DAD)  |
|  u32:1-n  |  Number of NS messages to send while performing DAD  |

<a id="nestedatt--ipv6--address"></a>
### Nested Schema for `ipv6.address`

Optional:

- `autoconf` (String) Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)
- `eui64` (String) Prefix for IPv6 address with MAC-based EUI-64

|  Format  |  Description  |
|----------|---------------|
|  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
- `no_default_link_local` (String) Remove the default link-local address from the interface



<a id="nestedatt--keep_alive"></a>
### Nested Schema for `keep_alive`

Optional:

- `failure_count` (String) Maximum number of keepalive packet failures

|  Format  |  Description  |
|----------|---------------|
|  u32:0-1000  |  Maximum number of keepalive packet failures  |
- `interval` (String) Keepalive packet interval in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600  |  Keepalive packet interval (seconds)  |


<a id="nestedatt--local_address"></a>
### Nested Schema for `local_address`

Optional:

- `subnet_mask` (String) Subnet-mask for local IP address of tunnel (IPv4 only)


<a id="nestedatt--mirror"></a>
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |


<a id="nestedatt--replace_default_route"></a>
### Nested Schema for `replace_default_route`

Optional:

- `local` (String) Tunnel endpoints are on the same subnet


<a id="nestedatt--server"></a>
### Nested Schema for `server`

Optional:

- `client` (Attributes Map) Client-specific settings

|  Format  |  Description  |
|----------|---------------|
|  name  |  Client common-name in the certificate  | (see [below for nested schema](#nestedatt--server--client))
- `client_ip_pool` (Attributes) Pool of client IPv4 addresses (see [below for nested schema](#nestedatt--server--client_ip_pool))
- `client_ipv6_pool` (Attributes) Pool of client IPv6 addresses (see [below for nested schema](#nestedatt--server--client_ipv6_pool))
- `domain_name` (String) DNS suffix to be pushed to all clients

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Domain Name Server suffix  |
- `max_connections` (String) Number of maximum client connections

|  Format  |  Description  |
|----------|---------------|
|  u32:1-4096  |  Number of concurrent clients  |
- `mfa` (Attributes) multi-factor authentication (see [below for nested schema](#nestedatt--server--mfa))
- `name_server` (String) Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
|  ipv6  |  Domain Name Server (DNS) IPv6 address  |
- `push_route` (Attributes Map) Route to be pushed to all clients

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 network and prefix length  |
|  ipv6net  |  IPv6 network and prefix length  | (see [below for nested schema](#nestedatt--server--push_route))
- `reject_unconfigured_clients` (String) Reject connections from clients that are not explicitly configured
- `subnet` (String) Server-mode subnet (from which client IPs are allocated)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 network and prefix length  |
|  ipv6net  |  IPv6 network and prefix length  |
- `topology` (String) Topology for clients

|  Format  |  Description  |
|----------|---------------|
|  net30  |  net30 topology  |
|  point-to-point  |  Point-to-point topology  |
|  subnet  |  Subnet topology  |

<a id="nestedatt--server--client"></a>
### Nested Schema for `server.client`

Optional:

- `disable` (String) Disable instance
- `ip` (String) IP address of the client

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Client IPv4 address  |
|  ipv6  |  Client IPv6 address  |
- `push_route` (String) Route to be pushed to the client

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 network and prefix length  |
|  ipv6net  |  IPv6 network and prefix length  |
- `subnet` (String) Subnet belonging to the client (iroute)

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 network and prefix length belonging to the client  |
|  ipv6net  |  IPv6 network and prefix length belonging to the client  |


<a id="nestedatt--server--client_ip_pool"></a>
### Nested Schema for `server.client_ip_pool`

Optional:

- `disable` (String) Disable instance
- `start` (String) First IP address in the pool

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
- `stop` (String) Last IP address in the pool

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |
- `subnet_mask` (String) Subnet mask pushed to dynamic clients. If not set the server subnet mask will be used. Only used with topology subnet or device type tap. Not used with bridged interfaces.

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 subnet mask  |


<a id="nestedatt--server--client_ipv6_pool"></a>
### Nested Schema for `server.client_ipv6_pool`

Optional:

- `base` (String) Client IPv6 pool base address with optional prefix length

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Client IPv6 pool base address with optional prefix length (defaults: base = server subnet + 0x1000, prefix length = server prefix length)  |
- `disable` (String) Disable instance


<a id="nestedatt--server--mfa"></a>
### Nested Schema for `server.mfa`

Optional:

- `totp` (Attributes) Time-based one-time passwords (see [below for nested schema](#nestedatt--server--mfa--totp))

<a id="nestedatt--server--mfa--totp"></a>
### Nested Schema for `server.mfa.totp`

Optional:

- `challenge` (String) Expect password as result of a challenge response protocol

|  Format  |  Description  |
|----------|---------------|
|  disable  |  Disable challenge-response  |
|  enable  |  Enable chalenge-response  |
- `digits` (String) Number of digits to use for totp hash

|  Format  |  Description  |
|----------|---------------|
|  1-65535  |  Seconds  |
- `drift` (String) Time drift in seconds

|  Format  |  Description  |
|----------|---------------|
|  1-65535  |  Seconds  |
- `slop` (String) Maximum allowed clock slop in seconds

|  Format  |  Description  |
|----------|---------------|
|  1-65535  |  Seconds  |
- `step` (String) Step value for totp in seconds

|  Format  |  Description  |
|----------|---------------|
|  1-65535  |  Seconds  |



<a id="nestedatt--server--push_route"></a>
### Nested Schema for `server.push_route`

Optional:

- `metric` (String) Set metric for this route

|  Format  |  Description  |
|----------|---------------|
|  u32:0-4294967295  |  Metric for this route  |



<a id="nestedatt--tls"></a>
### Nested Schema for `tls`

Optional:

- `auth_key` (String) TLS shared secret key for tls-auth
- `ca_certificate` (String) Certificate Authority chain in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of certificate in PKI configuration  |
- `crypt_key` (String) Static key to use to authenticate control channel
- `dh_params` (String) Diffie Hellman parameters (server only)
- `role` (String) TLS negotiation role

|  Format  |  Description  |
|----------|---------------|
|  active  |  Initiate TLS negotiation actively  |
|  passive  |  Wait for incoming TLS connection  |
- `tls_version_min` (String) Specify the minimum required TLS version

|  Format  |  Description  |
|----------|---------------|
|  1.0  |  TLS v1.0  |
|  1.1  |  TLS v1.1  |
|  1.2  |  TLS v1.2  |
|  1.3  |  TLS v1.3  |