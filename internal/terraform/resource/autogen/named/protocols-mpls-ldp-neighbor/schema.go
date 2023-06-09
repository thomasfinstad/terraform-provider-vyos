// Package namedprotocolsmplsldpneighbor code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsmplsldpneighbor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsMplsLdpNeighbor) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Multiprotocol Label Switching (MPLS)

Label Distribution Protocol (LDP)

LDP neighbor parameters

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Neighbor IPv4 address  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
