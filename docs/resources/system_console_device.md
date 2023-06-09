---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system_console_device Resource - vyos"
subcategory: ""
description: |-
  Serial console configuration
  Serial console device name
  |  Format  |  Description  |
  |----------|---------------|
  |  ttySN  |  TTY device name, regular serial port  |
  |  usbNbXpY  |  TTY device name, USB based  |
  |  hvcN  |  Xen console  |
---

# vyos_system_console_device (Resource)

Serial console configuration

Serial console device name

|  Format  |  Description  |
|----------|---------------|
|  ttySN  |  TTY device name, regular serial port  |
|  usbNbXpY  |  TTY device name, USB based  |
|  hvcN  |  Xen console  |



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Serial console device name

|  Format  |  Description  |
|----------|---------------|
|  ttySN  |  TTY device name, regular serial port  |
|  usbNbXpY  |  TTY device name, USB based  |
|  hvcN  |  Xen console  |

### Optional

- `speed` (String) Console baud rate

|  Format  |  Description  |
|----------|---------------|
|  1200  |  1200 bps  |
|  2400  |  2400 bps  |
|  4800  |  4800 bps  |
|  9600  |  9600 bps  |
|  19200  |  19200 bps  |
|  38400  |  38400 bps  |
|  57600  |  57600 bps  |
|  115200  |  115200 bps  |
