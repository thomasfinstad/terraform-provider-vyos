---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_parameters_distance_prefix Resource - vyos"
subcategory: ""
description: |-
  Border Gateway Protocol (BGP)
  BGP parameters
  Administratives distances for BGP routes
  Administrative distance for a specific BGP prefix
  |  Format  |  Description  |
  |----------|---------------|
  |  ipv4net  |  Administrative distance for a specific BGP prefix  |
---

# vyos_protocols_bgp_parameters_distance_prefix (Resource)

Border Gateway Protocol (BGP)

BGP parameters

Administratives distances for BGP routes

Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |

### Optional

- `distance` (String) Administrative distance for prefix

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Administrative distance for external BGP routes  |