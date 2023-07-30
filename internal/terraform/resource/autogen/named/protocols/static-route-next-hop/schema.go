// Package namedprotocolsstaticroutenexthop code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticroutenexthop

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsStaticRouteNextHop) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing protocols

Static Routing

Static IPv4 route

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  IPv4 static route  |

Next-hop IPv4 router address

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  Next-hop router address  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
