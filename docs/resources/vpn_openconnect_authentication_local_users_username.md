---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn_openconnect_authentication_local_users_username Resource - vyos"
subcategory: ""
description: |-
  SSL VPN OpenConnect, AnyConnect compatible server
  Authentication for remote access SSL VPN Server
  Local user authentication
  Username used for authentication
  |  Format  |  Description  |
  |----------|---------------|
  |  txt  |  Username used for authentication  |
---

# vyos_vpn_openconnect_authentication_local_users_username (Resource)

SSL VPN OpenConnect, AnyConnect compatible server

Authentication for remote access SSL VPN Server

Local user authentication

Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username used for authentication  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Username used for authentication

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Username used for authentication  |

### Optional

- `disable` (String) Disable instance
- `otp` (Attributes) 2FA OTP authentication parameters (see [below for nested schema](#nestedatt--otp))
- `password` (String) Password used for authentication

<a id="nestedatt--otp"></a>
### Nested Schema for `otp`

Optional:

- `interval` (String) Time tokens interval in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32:5-86400  |  Time tokens interval in seconds.  |
- `key` (String) Token Key Secret key for the token algorithm (see RFC 4226)

|  Format  |  Description  |
|----------|---------------|
|  txt  |  OTP key in hex-encoded format  |
- `otp_length` (String) Number of digits in OTP code

|  Format  |  Description  |
|----------|---------------|
|  u32:6-8  |  Number of digits in OTP code  |
- `token_type` (String) Token type

|  Format  |  Description  |
|----------|---------------|
|  hotp-time  |  Time-based OTP algorithm  |
|  hotp-event  |  Event-based OTP algorithm  |
