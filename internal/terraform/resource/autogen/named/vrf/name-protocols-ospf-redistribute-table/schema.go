// Package namedvrfnameprotocolsospfredistributetable code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfredistributetable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsOspfRedistributeTable) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding

Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

Routing protocol parameters

Open Shortest Path First (OSPF)

Redistribute information from another routing protocol

Redistribute non-main Kernel Routing Table

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-200  |  Policy route table number  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
