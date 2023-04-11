---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service dns dynamic interface Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_service dns dynamic interface (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ipv6_enable` (String) Allow explicit IPv6 addresses for Dynamic DNS for this interface
- `rfc2136` (Attributes Map) RFC2136 Update name (see [below for nested schema](#nestedatt--rfc2136))
- `service` (Attributes Map) Service being used for Dynamic DNS

Format: txt
Dynanmic DNS service with a custom name
Format: afraid
afraid.org Services
Format: changeip
changeip.com Services
Format: cloudflare
cloudflare.com Services
Format: dnspark
dnspark.com Services
Format: dslreports
dslreports.com Services
Format: dyndns
dyndns.com Services
Format: easydns
easydns.com Services
Format: namecheap
namecheap.com Services
Format: noip
noip.com Services
Format: sitelutions
sitelutions.com Services
Format: zoneedit
zoneedit.com Services (see [below for nested schema](#nestedatt--service))
- `use_web` (Attributes) Web check used for obtaining the external IP address (see [below for nested schema](#nestedatt--use_web))

<a id="nestedatt--rfc2136"></a>
### Nested Schema for `rfc2136`

Optional:

- `key` (String) File containing the secret key shared with remote DNS server

Format: filename
File in /config/auth directory
- `record` (String) Record to be updated
- `server` (String) Server to be updated
- `ttl` (String) Time To Live (default: 600)

Format: u32:1-86400
DNS forwarding cache size
- `zone` (String) Zone to be updated


<a id="nestedatt--service"></a>
### Nested Schema for `service`

Optional:

- `host_name` (String) Hostname registered with DDNS service
- `login` (String) Login for DDNS service
- `password` (String) Password for DDNS service
- `protocol` (String) ddclient protocol used for DDNS service

Format: changeip
ChangeIP protocol
Format: cloudflare
Cloudflare protocol
Format: dnsmadeeasy
DNS Made Easy protocol
Format: dnspark
DNS Park protocol
Format: dondominio
DonDominio protocol
Format: dslreports1
DslReports protocol
Format: dtdns
DtDNS protocol
Format: duckdns
DuckDNS protocol
Format: dyndns2
DynDNS protocol v2
Format: easydns
easyDNS protocol
Format: freedns
FreeDNS protocol
Format: freemyip
freemyip protocol
Format: googledomains
Google domains protocol
Format: hammernode1
Hammernode protocol
Format: namecheap
Namecheap protocol
Format: nfsn
NearlyFreeSpeech DNS protocol
Format: noip
No-IP protocol
Format: sitelutions
Sitelutions protocol
Format: woima
WOIMA protocol
Format: yandex
Yandex.DNS protocol
Format: zoneedit1
Zoneedit protocol
- `server` (String) Remote server to connect to

Format: ipv4
Server IPv4 address
Format: hostname
Server hostname/FQDN
- `zone` (String) DNS zone to update (only available with CloudFlare)


<a id="nestedatt--use_web"></a>
### Nested Schema for `use_web`

Optional:

- `skip` (String) Skip everything before this on the given URL
- `url` (String) URL to obtain the current external IP address