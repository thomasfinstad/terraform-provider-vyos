// Package namedsystemconsoledevice code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconsoledevice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemConsoleDevice) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Serial console configuration

Serial console device name

|  Format  |  Description  |
|----------|---------------|
|  ttySN  |  TTY device name, regular serial port  |
|  usbNbXpY  |  TTY device name, USB based  |
|  hvcN  |  Xen console  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
