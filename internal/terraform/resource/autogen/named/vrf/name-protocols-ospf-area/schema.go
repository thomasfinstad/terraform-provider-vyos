// Package namedvrfnameprotocolsospfarea code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfarea

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsOspfArea) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding

Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

Routing protocol parameters

Open Shortest Path First (OSPF)

OSPF area settings

    |  Format  |  Description  |
    |----------|---------------|
    |  u32  |  OSPF area number in decimal notation  |
    |  ipv4  |  OSPF area number in dotted decimal notation  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
