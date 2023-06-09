// Package namedservicebroadcastrelayid code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicebroadcastrelayid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceBroadcastRelayID) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `UDP broadcast relay service

Unique ID for each UDP port to forward

|  Format  |  Description  |
|----------|---------------|
|  u32:1-99  |  Broadcast relay instance ID  |

`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
