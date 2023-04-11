---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system login user Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_system login user (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `full_name` (String) Full name of the user (use quotes for names with spaces)
- `home_directory` (String) Home directory

<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `encrypted_password` (String) Encrypted password
- `otp` (Attributes) One-Time-Pad (two-factor) authentication parameters (see [below for nested schema](#nestedatt--authentication--otp))
- `plaintext_password` (String) Plaintext password used for encryption
- `public_keys` (Attributes Map) Remote access public keys

Format: txt
Key identifier used by ssh-keygen (usually of form user@host) (see [below for nested schema](#nestedatt--authentication--public_keys))

<a id="nestedatt--authentication--otp"></a>
### Nested Schema for `authentication.otp`

Optional:

- `key` (String) Key/secret the token algorithm (see RFC4226)

Format: txt
Base32 encoded key/token
- `rate_limit` (String) Limit number of logins (rate-limit) per rate-time

Format: u32:1-10
Number of attempts
- `rate_time` (String) Limit number of logins (rate-limit) per rate-time

Format: u32:15-600
Time interval
- `window_size` (String) Set window of concurrently valid codes

Format: u32:1-21
Window size


<a id="nestedatt--authentication--public_keys"></a>
### Nested Schema for `authentication.public_keys`

Optional:

- `key` (String) Public key value (Base64 encoded)
- `options` (String) Optional public key options
- `type` (String) SSH public key type

Format: ssh-dss
Digital Signature Algorithm (DSA) key support
Format: ssh-rsa
Key pair based on RSA algorithm
Format: ecdsa-sha2-nistp256
Elliptic Curve DSA with NIST P-256 curve
Format: ecdsa-sha2-nistp384
Elliptic Curve DSA with NIST P-384 curve
Format: ecdsa-sha2-nistp521
Elliptic Curve DSA with NIST P-521 curve
Format: ssh-ed25519
Edwards-curve DSA with elliptic curve 25519
Format: sk-ecdsa-sha2-nistp256@openssh.com
Elliptic Curve DSA security key
Format: sk-ssh-ed25519@openssh.com
Elliptic curve 25519 security key