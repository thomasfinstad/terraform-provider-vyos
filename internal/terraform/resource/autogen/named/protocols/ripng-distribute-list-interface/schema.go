// Package namedprotocolsripngdistributelistinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsripngdistributelistinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsRIPngDistributeListInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing Information Protocol (RIPng) parameters

Filter networks in routing updates

Apply filtering to an interface

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Apply filtering to an interface  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
