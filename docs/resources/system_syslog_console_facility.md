---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system_syslog_console_facility Resource - vyos"
subcategory: ""
description: |-
  System logging
  logging to serial console
  Facility for logging
  |  Format  |  Description  |
  |----------|---------------|
  |  all  |  All facilities excluding "mark"  |
  |  auth  |  Authentication and authorization  |
  |  authpriv  |  Non-system authorization  |
  |  cron  |  Cron daemon  |
  |  daemon  |  System daemons  |
  |  kern  |  Kernel  |
  |  lpr  |  Line printer spooler  |
  |  mail  |  Mail subsystem  |
  |  mark  |  Timestamp  |
  |  news  |  USENET subsystem  |
  |  protocols  |  depricated will be set to local7  |
  |  security  |  depricated will be set to auth  |
  |  syslog  |  Authentication and authorization  |
  |  user  |  Application processes  |
  |  uucp  |  UUCP subsystem  |
  |  local0  |  Local facility 0  |
  |  local1  |  Local facility 1  |
  |  local2  |  Local facility 2  |
  |  local3  |  Local facility 3  |
  |  local4  |  Local facility 4  |
  |  local5  |  Local facility 5  |
  |  local6  |  Local facility 6  |
  |  local7  |  Local facility 7  |
---

# vyos_system_syslog_console_facility (Resource)

System logging

logging to serial console

Facility for logging

|  Format  |  Description  |
|----------|---------------|
|  all  |  All facilities excluding "mark"  |
|  auth  |  Authentication and authorization  |
|  authpriv  |  Non-system authorization  |
|  cron  |  Cron daemon  |
|  daemon  |  System daemons  |
|  kern  |  Kernel  |
|  lpr  |  Line printer spooler  |
|  mail  |  Mail subsystem  |
|  mark  |  Timestamp  |
|  news  |  USENET subsystem  |
|  protocols  |  depricated will be set to local7  |
|  security  |  depricated will be set to auth  |
|  syslog  |  Authentication and authorization  |
|  user  |  Application processes  |
|  uucp  |  UUCP subsystem  |
|  local0  |  Local facility 0  |
|  local1  |  Local facility 1  |
|  local2  |  Local facility 2  |
|  local3  |  Local facility 3  |
|  local4  |  Local facility 4  |
|  local5  |  Local facility 5  |
|  local6  |  Local facility 6  |
|  local7  |  Local facility 7  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Facility for logging

|  Format  |  Description  |
|----------|---------------|
|  all  |  All facilities excluding "mark"  |
|  auth  |  Authentication and authorization  |
|  authpriv  |  Non-system authorization  |
|  cron  |  Cron daemon  |
|  daemon  |  System daemons  |
|  kern  |  Kernel  |
|  lpr  |  Line printer spooler  |
|  mail  |  Mail subsystem  |
|  mark  |  Timestamp  |
|  news  |  USENET subsystem  |
|  protocols  |  depricated will be set to local7  |
|  security  |  depricated will be set to auth  |
|  syslog  |  Authentication and authorization  |
|  user  |  Application processes  |
|  uucp  |  UUCP subsystem  |
|  local0  |  Local facility 0  |
|  local1  |  Local facility 1  |
|  local2  |  Local facility 2  |
|  local3  |  Local facility 3  |
|  local4  |  Local facility 4  |
|  local5  |  Local facility 5  |
|  local6  |  Local facility 6  |
|  local7  |  Local facility 7  |

### Optional

- `level` (String) Logging level

|  Format  |  Description  |
|----------|---------------|
|  emerg  |  Emergency messages  |
|  alert  |  Urgent messages  |
|  crit  |  Critical messages  |
|  err  |  Error messages  |
|  warning  |  Warning messages  |
|  notice  |  Messages for further investigation  |
|  info  |  Informational messages  |
|  debug  |  Debug messages  |
|  all  |  Log everything  |
