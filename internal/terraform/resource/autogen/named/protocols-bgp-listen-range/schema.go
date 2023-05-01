// Package namedprotocolsbgplistenrange code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgplistenrange

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpListenRange) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Border Gateway Protocol (BGP)

Listen for and accept BGP dynamic neighbors from range

BGP dynamic neighbors listen range

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 dynamic neighbors listen range  |
|  ipv6net  |  IPv6 dynamic neighbors listen range  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
