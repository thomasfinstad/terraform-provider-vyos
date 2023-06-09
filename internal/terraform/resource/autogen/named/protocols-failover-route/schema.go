// Package namedprotocolsfailoverroute code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsfailoverroute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsFailoverRoute) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Failover Routing

Failover IPv4 route

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 failover route  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
