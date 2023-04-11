---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn ipsec ike-group Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_vpn ipsec ike-group (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `close_action` (String) Action to take if a child SA is unexpectedly closed

Format: none
Do nothing
Format: hold
Attempt to re-negotiate when matching traffic is seen
Format: restart
Attempt to re-negotiate the connection immediately
- `dead_peer_detection` (Attributes) Dead Peer Detection (DPD) (see [below for nested schema](#nestedatt--dead_peer_detection))
- `disable_mobike` (String) Disable MOBIKE Support (IKEv2 only)
- `ikev2_reauth` (String) Re-authentication of the remote peer during an IKE re-key (IKEv2 only)
- `key_exchange` (String) IKE version

Format: ikev1
Use IKEv1 for key exchange
Format: ikev2
Use IKEv2 for key exchange
- `lifetime` (String) IKE lifetime

Format: u32:30-86400
IKE lifetime in seconds
- `mode` (String) IKEv1 phase 1 mode

Format: main
Use the main mode (recommended)
Format: aggressive
Use the aggressive mode (insecure, not recommended)
- `proposal` (Attributes Map) IKE proposal

Format: u32:1-65535
IKE group proposal (see [below for nested schema](#nestedatt--proposal))

<a id="nestedatt--dead_peer_detection"></a>
### Nested Schema for `dead_peer_detection`

Optional:

- `action` (String) Keep-alive failure action

Format: hold
Attempt to re-negotiate the connection when matching traffic is seen
Format: clear
Remove the connection immediately
Format: restart
Attempt to re-negotiate the connection immediately
- `interval` (String) Keep-alive interval

Format: u32:2-86400
Keep-alive interval in seconds
- `timeout` (String) Dead Peer Detection keep-alive timeout (IKEv1 only)

Format: u32:2-86400
Keep-alive timeout in seconds


<a id="nestedatt--proposal"></a>
### Nested Schema for `proposal`

Optional:

- `dh_group` (String) dh-grouphelp

Format: 1
Diffie-Hellman group 1 (modp768)
Format: 2
Diffie-Hellman group 2 (modp1024)
Format: 5
Diffie-Hellman group 5 (modp1536)
Format: 14
Diffie-Hellman group 14 (modp2048)
Format: 15
Diffie-Hellman group 15 (modp3072)
Format: 16
Diffie-Hellman group 16 (modp4096)
Format: 17
Diffie-Hellman group 17 (modp6144)
Format: 18
Diffie-Hellman group 18 (modp8192)
Format: 19
Diffie-Hellman group 19 (ecp256)
Format: 20
Diffie-Hellman group 20 (ecp384)
Format: 21
Diffie-Hellman group 21 (ecp521)
Format: 22
Diffie-Hellman group 22 (modp1024s160)
Format: 23
Diffie-Hellman group 23 (modp2048s224)
Format: 24
Diffie-Hellman group 24 (modp2048s256)
Format: 25
Diffie-Hellman group 25 (ecp192)
Format: 26
Diffie-Hellman group 26 (ecp224)
Format: 27
Diffie-Hellman group 27 (ecp224bp)
Format: 28
Diffie-Hellman group 28 (ecp256bp)
Format: 29
Diffie-Hellman group 29 (ecp384bp)
Format: 30
Diffie-Hellman group 30 (ecp512bp)
Format: 31
Diffie-Hellman group 31 (curve25519)
Format: 32
Diffie-Hellman group 32 (curve448)
- `encryption` (String) Encryption algorithm

Format: null
Null encryption
Format: aes128
128 bit AES-CBC
Format: aes192
192 bit AES-CBC
Format: aes256
256 bit AES-CBC
Format: aes128ctr
128 bit AES-COUNTER
Format: aes192ctr
192 bit AES-COUNTER
Format: aes256ctr
256 bit AES-COUNTER
Format: aes128ccm64
128 bit AES-CCM with 64 bit ICV
Format: aes192ccm64
192 bit AES-CCM with 64 bit ICV
Format: aes256ccm64
256 bit AES-CCM with 64 bit ICV
Format: aes128ccm96
128 bit AES-CCM with 96 bit ICV
Format: aes192ccm96
192 bit AES-CCM with 96 bit ICV
Format: aes256ccm96
256 bit AES-CCM with 96 bit ICV
Format: aes128ccm128
128 bit AES-CCM with 128 bit ICV
Format: aes192ccm128
192 bit AES-CCM with 128 bit IC
Format: aes256ccm128
256 bit AES-CCM with 128 bit ICV
Format: aes128gcm64
128 bit AES-GCM with 64 bit ICV
Format: aes192gcm64
192 bit AES-GCM with 64 bit ICV
Format: aes256gcm64
256 bit AES-GCM with 64 bit ICV
Format: aes128gcm96
128 bit AES-GCM with 96 bit ICV
Format: aes192gcm96
192 bit AES-GCM with 96 bit ICV
Format: aes256gcm96
256 bit AES-GCM with 96 bit ICV
Format: aes128gcm128
128 bit AES-GCM with 128 bit ICV
Format: aes192gcm128
192 bit AES-GCM with 128 bit ICV
Format: aes256gcm128
256 bit AES-GCM with 128 bit ICV
Format: aes128gmac
Null encryption with 128 bit AES-GMAC
Format: aes192gmac
Null encryption with 192 bit AES-GMAC
Format: aes256gmac
Null encryption with 256 bit AES-GMAC
Format: 3des
168 bit 3DES-EDE-CBC
Format: blowfish128
128 bit Blowfish-CBC
Format: blowfish192
192 bit Blowfish-CBC
Format: blowfish256
256 bit Blowfish-CBC
Format: camellia128
128 bit Camellia-CBC
Format: camellia192
192 bit Camellia-CBC
Format: camellia256
256 bit Camellia-CBC
Format: camellia128ctr
128 bit Camellia-COUNTER
Format: camellia192ctr
192 bit Camellia-COUNTER
Format: camellia256ctr
256 bit Camellia-COUNTER
Format: camellia128ccm64
128 bit Camellia-CCM with 64 bit ICV
Format: camellia192ccm64
192 bit Camellia-CCM with 64 bit ICV
Format: camellia256ccm64
256 bit Camellia-CCM with 64 bit ICV
Format: camellia128ccm96
128 bit Camellia-CCM with 96 bit ICV
Format: camellia192ccm96
192 bit Camellia-CCM with 96 bit ICV
Format: camellia256ccm96
256 bit Camellia-CCM with 96 bit ICV
Format: camellia128ccm128
128 bit Camellia-CCM with 128 bit ICV
Format: camellia192ccm128
192 bit Camellia-CCM with 128 bit ICV
Format: camellia256ccm128
256 bit Camellia-CCM with 128 bit ICV
Format: serpent128
128 bit Serpent-CBC
Format: serpent192
192 bit Serpent-CBC
Format: serpent256
256 bit Serpent-CBC
Format: twofish128
128 bit Twofish-CBC
Format: twofish192
192 bit Twofish-CBC
Format: twofish256
256 bit Twofish-CBC
Format: cast128
128 bit CAST-CBC
Format: chacha20poly1305
256 bit ChaCha20/Poly1305 with 128 bit ICV
- `hash` (String) Hash algorithm

Format: md5
MD5 HMAC
Format: md5_128
MD5_128 HMAC
Format: sha1
SHA1 HMAC
Format: sha1_160
SHA1_160 HMAC
Format: sha256
SHA2_256_128 HMAC
Format: sha256_96
SHA2_256_96 HMAC
Format: sha384
SHA2_384_192 HMAC
Format: sha512
SHA2_512_256 HMAC
Format: aesxcbc
AES XCBC
Format: aescmac
AES CMAC
Format: aes128gmac
128-bit AES-GMAC
Format: aes192gmac
192-bit AES-GMAC
Format: aes256gmac
256-bit AES-GMAC
- `prf` (String) Pseudo-Random Functions

Format: prfmd5
MD5 PRF
Format: prfsha1
SHA1 PRF
Format: prfaesxcbc
AES XCBC PRF
Format: prfaescmac
AES CMAC PRF
Format: prfsha256
SHA2_256 PRF
Format: prfsha384
SHA2_384 PRF
Format: prfsha512
SHA2_512 PRF