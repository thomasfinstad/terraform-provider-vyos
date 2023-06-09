---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_rpki_cache Resource - vyos"
subcategory: ""
description: |-
  BGP prefix origin validation
  RPKI cache server address
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4  |  IP address of RPKI server  |
  |  ipv6  |  IPv6 address of RPKI server  |
  |  hostname  |  Fully qualified domain name of RPKI server  |
---

# vyos_protocols_rpki_cache (Resource)

BGP prefix origin validation

RPKI cache server address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of RPKI server  |
|  ipv6  |  IPv6 address of RPKI server  |
|  hostname  |  Fully qualified domain name of RPKI server  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) RPKI cache server address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of RPKI server  |
|  ipv6  |  IPv6 address of RPKI server  |
|  hostname  |  Fully qualified domain name of RPKI server  |

### Optional

- `port` (String) Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
- `preference` (String) Preference of the cache server

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Preference of the cache server  |
- `ssh` (Attributes) RPKI SSH connection settings (see [below for nested schema](#nestedatt--ssh))

<a id="nestedatt--ssh"></a>
### Nested Schema for `ssh`

Optional:

- `known_hosts_file` (String) RPKI SSH known hosts file
- `private_key_file` (String) RPKI SSH private key file
- `public_key_file` (String) RPKI SSH public key file path
- `username` (String) Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username  |
