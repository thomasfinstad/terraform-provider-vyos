// Package namedservicentpserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicentpserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceNtpServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Network Time Protocol (NTP) configuration

Network Time Protocol (NTP) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of NTP server  |
|  ipv6  |  IPv6 address of NTP server  |
|  hostname  |  Fully qualified domain name of NTP server  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
