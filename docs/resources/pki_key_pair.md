---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_pki_key_pair Resource - vyos"
subcategory: ""
description: |-
  VyOS PKI configuration
  Public and private keys
---

# vyos_pki_key_pair (Resource)

VyOS PKI configuration

Public and private keys



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Public and private keys

### Optional

- `private` (Attributes) Private key (see [below for nested schema](#nestedatt--private))
- `public` (Attributes) Public key (see [below for nested schema](#nestedatt--public))

<a id="nestedatt--private"></a>
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (String) Private key is password protected


<a id="nestedatt--public"></a>
### Nested Schema for `public`

Optional:

- `key` (String) Public key in PEM format