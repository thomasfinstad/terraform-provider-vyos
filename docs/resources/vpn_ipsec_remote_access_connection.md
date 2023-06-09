---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn_ipsec_remote_access_connection Resource - vyos"
subcategory: ""
description: |-
  Virtual Private Network (VPN)
  VPN IP security (IPsec) parameters
  IKEv2 remote access VPN
  IKEv2 VPN connection name
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Connection name  |
---

# vyos_vpn_ipsec_remote_access_connection (Resource)

Virtual Private Network (VPN)

VPN IP security (IPsec) parameters

IKEv2 remote access VPN

IKEv2 VPN connection name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Connection name  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) IKEv2 VPN connection name

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Connection name  |

### Optional

- `authentication` (Attributes) Authentication for remote access (see [below for nested schema](#nestedatt--authentication))
- `description` (String) Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |
- `disable` (String) Disable instance
- `esp_group` (String) Encapsulating Security Payloads (ESP) group name
- `ike_group` (String) Internet Key Exchange (IKE) group name
- `local` (Attributes) Local parameters for interesting traffic (see [below for nested schema](#nestedatt--local))
- `local_address` (String) IPv4 or IPv6 address of a local interface to use for VPN

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of a local interface for VPN  |
|  ipv6  |  IPv6 address of a local interface for VPN  |
|  any  |  Allow any IPv4 address present on the system to be used for VPN  |
- `pool` (String) IP address pool

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Predefined IP pool name  |
|  dhcp  |  Forward requests for virtual IP addresses to a DHCP server  |
|  radius  |  Forward requests for virtual IP addresses to a RADIUS server  |
- `timeout` (String) Timeout to close connection if no data is transmitted

|  Format  |  Description  |
|----------|---------------|
|  u32:0  |  Disable inactivity checks  |
|  u32:1-86400  |  Timeout in seconds  |
- `unique` (String) Connection uniqueness enforcement policy

|  Format  |  Description  |
|----------|---------------|
|  never  |  Never enforce connection uniqueness  |
|  keep  |  Reject new connection attempts if the same user already has an active connection  |
|  replace  |  Delete any existing connection if a new one for the same user gets established  |

<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `client_mode` (String) Client authentication mode

|  Format  |  Description  |
|----------|---------------|
|  eap-tls  |  Use EAP-TLS authentication  |
|  eap-mschapv2  |  Use EAP-MSCHAPv2 authentication  |
|  eap-radius  |  Use EAP-RADIUS authentication  |
- `local_id` (String) Local ID for peer authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Local ID used for peer authentication  |
- `local_users` (Attributes) Local user authentication (see [below for nested schema](#nestedatt--authentication--local_users))
- `pre_shared_secret` (String) Pre-shared secret key

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Pre-shared secret key  |
- `server_mode` (String) Server authentication mode

|  Format  |  Description  |
|----------|---------------|
|  pre-shared-secret  |  Use a pre-shared secret key  |
|  x509  |  Use x.509 certificate  |
- `x509` (Attributes) X.509 certificate (see [below for nested schema](#nestedatt--authentication--x509))

<a id="nestedatt--authentication--local_users"></a>
### Nested Schema for `authentication.local_users`

Optional:

- `username` (Attributes Map) Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username used for authentication  | (see [below for nested schema](#nestedatt--authentication--local_users--username))

<a id="nestedatt--authentication--local_users--username"></a>
### Nested Schema for `authentication.local_users.username`

Optional:

- `disable` (String) Disable instance
- `password` (String) Password used for authentication



<a id="nestedatt--authentication--x509"></a>
### Nested Schema for `authentication.x509`

Optional:

- `ca_certificate` (String) Certificate Authority in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Name of certificate in PKI configuration  |
- `passphrase` (String) Private key passphrase

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Passphrase to decrypt the private key  |



<a id="nestedatt--local"></a>
### Nested Schema for `local`

Optional:

- `port` (String) Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
- `prefix` (String) Local IPv4 or IPv6 prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Local IPv4 prefix  |
|  ipv6net  |  Local IPv6 prefix  |
