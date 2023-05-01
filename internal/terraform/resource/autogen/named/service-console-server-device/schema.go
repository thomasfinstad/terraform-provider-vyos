// Package namedserviceconsoleserverdevice code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceconsoleserverdevice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceConsoleServerDevice) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Serial Console Server

System serial interface name (ttyS or ttyUSB)

|  Format  |  Description  |
|----------|---------------|
|  ttySxxx  |  Regular serial interface  |
|  usbxbxpx  |  USB based serial interface  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
