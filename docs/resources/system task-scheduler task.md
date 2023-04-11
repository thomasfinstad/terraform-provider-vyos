---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system task-scheduler task Resource - vyos"
subcategory: ""
description: |-
  Example resource
---

# vyos_system task-scheduler task (Resource)

Example resource



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `crontab_spec` (String) UNIX crontab time specification string
- `executable` (Attributes) Executable path and arguments (see [below for nested schema](#nestedatt--executable))
- `interval` (String) Execution interval

Format: <minutes>
Execution interval in minutes
Format: <minutes>m
Execution interval in minutes
Format: <hours>h
Execution interval in hours
Format: <days>d
Execution interval in days

<a id="nestedatt--executable"></a>
### Nested Schema for `executable`

Optional:

- `arguments` (String) Arguments passed to the executable
- `path` (String) Path to executable